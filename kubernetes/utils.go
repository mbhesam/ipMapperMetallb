package kubernetes

import (
	"fmt"
	"ipMapperApi/logger"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv() []string {
	logging := logger.GetLogger()
	err := godotenv.Load("/home/mb/ipMapperApi/.env")
	if err != nil {
		panic(err.Error())
	}
	kubeconfig := os.Getenv("KUBECONFIG_PATH")
	if kubeconfig == "" {
		message := "KUBECONFIG_PATH environment variable is not set."
		logging.Error(message)
		os.Exit(1)
	}
	namespace := os.Getenv("NAMESPACE")
	if namespace == "" {
		message := "NAMESPACE environment variable is not set."
		logging.Error(message)
		os.Exit(1)
	}
	return []string{kubeconfig, namespace}
}

func GiveResults() []map[string]string {
	kubeconfig := GetEnv()[0]
	namespace := GetEnv()[1]
	dynamicclient := connectK8sDynamic(K8sParams{kubeconfig: kubeconfig})
	resources_servicel2status := RetriveResources(K8sParams{namespace: namespace, dynamicclient: dynamicclient}, GvrParams{group: "metallb.io", version: "v1beta1", resource: "servicel2statuses"})
	serviceNodeNamespaceInfo := RetriveServiceNode(resources_servicel2status)
	clientset := connectK8sKubernetes(K8sParams{kubeconfig: kubeconfig})
	completeInfo := RetriveServiceInfo(K8sParams{clientset: clientset}, serviceNodeNamespaceInfo)
	return completeInfo
}

func GivePerIP(ip string) map[string]string {
	logging := logger.GetLogger()
	completeInfo := GiveResults()
	var result string
	for _, binding := range completeInfo {
		if binding["public_ip"] == ip {
			result = binding["private_ip"]
			break // Exit loop after finding the first match
		}
	}
	if result == "" {
		message := fmt.Sprintf("No Node Found for %s", ip)
		logging.Error(message)
	}
	return map[string]string{"node": result}
}
