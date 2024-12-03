package mac

import (
	"fmt"
	"ipMapperApi/logger"
	"net"
)

func GetIPByHostname(hostname string) string {
	logging := logger.GetLogger()
	ips, err := net.LookupIP(hostname)
	if err != nil {
		message := fmt.Sprintf("failed to get IP addresses for hostname %s: %v", hostname, err)
		logging.Error(message)
	}
	return ips[0].String()
}

// isValidIPv4 checks if the provided string is a valid IPv4 address.
func IsValidIPv4(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil && parsedIP.To4() != nil
}
