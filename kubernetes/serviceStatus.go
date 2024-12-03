package kubernetes

import (
	"context"
	"fmt"
	"ipMapperApi/mac"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"

	// "k8s.io/apimachinery/pkg/util/errors"
	// "k8s.io/client-go/kubernetes"
	"ipMapperApi/logger"
	"os"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sParams struct {
	kubeconfig    string
	namespace     string
	dynamicclient dynamic.Interface
	clientset     *kubernetes.Clientset
}

type GvrParams struct {
	group    string
	version  string
	resource string
}

func connectK8sDynamic(k8sparams K8sParams) dynamic.Interface {
	logging := logger.GetLogger()
	config, err := clientcmd.BuildConfigFromFlags("", k8sparams.kubeconfig)
	if err != nil {
		message := fmt.Sprintf("Error building kubeconfig: %s\n", err)
		logging.Error(message)
	}

	// Create Kubernetes client
	dynamicclient, err := dynamic.NewForConfig(config)
	if err != nil {
		message := fmt.Sprintf("Error creating Kubernetes client: %s\n", err)
		logging.Error(message)
	}
	return dynamicclient
}

func RetriveResources(k8sparams K8sParams, gvrparams GvrParams) *unstructured.UnstructuredList {
	logging := logger.GetLogger()
	namespace := k8sparams.namespace
	dynamicclient := k8sparams.dynamicclient
	// Define the GVR (GroupVersionResource) for servicel2statuses
	gvr := schema.GroupVersionResource{
		Group:    gvrparams.group,
		Version:  gvrparams.version,
		Resource: gvrparams.resource,
	}
	resources, err := dynamicclient.Resource(gvr).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		message := fmt.Sprintf("Error getting %s in namespace %s: %s\n", gvrparams.resource, namespace, err)
		logging.Error(message)
	}
	return resources
}

func RetriveServiceNode(resources *unstructured.UnstructuredList) []map[string]string {
	servicesInfo := []map[string]string{}
	// fmt.Println(resources)
	for _, resource := range resources.Items {
		labels := resource.GetLabels()
		node := labels["metallb.io/node"]
		service := labels["metallb.io/service-name"]
		namespace := labels["metallb.io/service-namespace"]
		servicesInfo = append(servicesInfo, map[string]string{"name": service, "node": node, "namespace": namespace})
	}
	return servicesInfo
}

func connectK8sKubernetes(k8sparams K8sParams) *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", k8sparams.kubeconfig)
	if err != nil {
		fmt.Printf("Error building kubeconfig: %s\n", err)
		os.Exit(1)
	}

	// Create Kubernetes client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating Kubernetes client: %s\n", err)
		os.Exit(1)
	}
	return clientset
}

// func RetriveServices(k8sparams K8sParams) []v1.Service {
// 	logging := logger.GetLogger()
// 	namespace := k8sparams.namespace
// 	clientset := k8sparams.clientset
// 	services, err := clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
// 	if err != nil {
// 		message := fmt.Sprintf("Error getting services in namespace %s: %s\n", namespace, err)
// 		logging.Error(message)
// 	}

// 	var loadBalancerServices []v1.Service

//     // Filter services for LoadBalancer type
//     for _, service := range services.Items {
//         if service.Spec.Type == "LoadBalancer" {
//             loadBalancerServices = append(loadBalancerServices, service)
//         }
//     }
// 	return loadBalancerServices

// }

func RetriveServiceInfo(k8sparams K8sParams, serviceNodeNamespaceInfo []map[string]string) []map[string]string {
	logging := logger.GetLogger()
	clientset := k8sparams.clientset
	servicesInfo := []map[string]string{}
	for _, svc := range serviceNodeNamespaceInfo {
		namespace := svc["namespace"]
		serviceName := svc["name"]
		node := mac.GetIPByHostname(svc["node"])
		service, err := clientset.CoreV1().Services(namespace).Get(context.TODO(), serviceName, metav1.GetOptions{})
		if err != nil {
			message := fmt.Sprintf("Error getting service %s in namespace %s: %s\n", service, namespace, err)
			logging.Error(message)
		}
		anotation := service.GetAnnotations()
		lbIP := anotation["metallb.universe.tf/loadBalancerIPs"]
		servicesInfo = append(servicesInfo, map[string]string{"public_ip": lbIP, "private_ip": node})
	}
	return servicesInfo
}
