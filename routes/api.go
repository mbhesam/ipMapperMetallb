package routes

import (
	_ "fmt"
	"ipMapperApi/kubernetes"
	"ipMapperApi/mac"
	"net/http"

	"github.com/gin-gonic/gin"
)

// // ShowIP godoc
// // @Summary Show IP addresses
// // @Description Get a list of IP addresses
// // @Tags IP
// // @Produce json
// // @Success 200 {object} map[string]interface{}
// // @Router /show [get]
// func ShowIP(c *gin.Context) {
// 	ips := []string{"5.106.9.24", "5.106.9.25", "5.106.9.11"}
// 	results := mac.GiveResult(ips)
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": results})
// }

/////version V1 API

// ShowIP godoc
// @Summary Show IP addresses
// @Description Get a list of IP addresses
// @Tags IP
// @Produce json
// @Success 200 {object} map[string][]map[string]string{}
// @Router /v1/bindings [get]
func V1ShowAll(c *gin.Context) {
	results := kubernetes.GiveResults()
	c.JSON(http.StatusOK, gin.H{"bindings": results})
}

// ShowIPPerIP godoc
// @Summary Show node of specefic ip
// @Description Get node of specefic ip
// @Tags IP
// @Produce json
// @Param ip path string true "IP Address"  // Add this line
// @Success 200 {object} map[string]map[string]string{}
// @Router /v1/bind_ip/{ip} [get]
func V1ShowPerIP(c *gin.Context) {
	ip := c.Param("ip")
	if !mac.IsValidIPv4(ip) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: not a valid IPv4 address"})
		return
	}
	results := kubernetes.GivePerIP(ip)
	if results["node"] == "" {
		c.JSON(http.StatusOK, gin.H{ip: "No node found for your public ip"})
		return
	}
	c.JSON(http.StatusOK, results)
}

func V1redirect(c *gin.Context) {
	c.Redirect(http.StatusFound, "/docs/index.html")
}
