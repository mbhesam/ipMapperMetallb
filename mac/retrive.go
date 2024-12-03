package mac

import (
	"bytes"
	"fmt"
	"ipMapperApi/logger"
	"os/exec"
	"regexp"
)

const ShellToUse = "bash"

var finalMatch string
var ipPrivate string

func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func ExecuteArping(ip string, counts ...int) [2]string {
	count := 2 // Default count
	if len(counts) > 0 {
		count = counts[0] // Use the first provided count
	}
	command := fmt.Sprintf("arping -c %d %s", count, ip)
	logging := logger.GetLogger()
	output, errout, err := Shellout(command)
	if err != nil {
		message := fmt.Sprintf("error: %v with message %s\n", err, errout)
		logging.Error(message)
	}
	ipOutput := [2]string{ip, output}
	return ipOutput
}

func FindMacaddress(ipOutput [2]string) [2]string {
	logging := logger.GetLogger()
	// Use a regex to extract the MAC address
	macRegex := regexp.MustCompile(`\[(.*?)\]`)
	ip := ipOutput[0]
	output := ipOutput[1]
	matches := macRegex.FindAllStringSubmatch(output, -1)

	if matches != nil {
		for _, match := range matches {
			if len(match) > 0 {
				finalMatch = match[len(match)-1]
				message := fmt.Sprintf("Extracted MAC: %s", finalMatch)
				logging.Info(message)
			}
		}
	} else {
		message := "No MAC address found."
		logging.Info(message)
		finalMatch = ""
	}
	ipMac := [2]string{ip, finalMatch}
	return ipMac
}

func MapIP(ipMac [2]string, physicalInterfaces ...string) [2]string {
	logging := logger.GetLogger()
	physicalInterface := "ens160" // Default interface
	if len(physicalInterfaces) > 0 {
		physicalInterface = physicalInterfaces[0] // Use the first provided interface
	}
	ipPublic := ipMac[0]
	macNode := ipMac[1]
	command := fmt.Sprintf("arp-scan --interface=%s -l | grep -i '%s'", physicalInterface, macNode)
	output, errout, err := Shellout(command)
	if err != nil {
		message := fmt.Sprintf("error: %v with message %s\n", err, errout)
		logging.Error(message)
	}
	re := regexp.MustCompile(`([0-9]{1,3}\.){3}[0-9]{1,3}\s+([0-9a-fA-F]{2}[:-]){5}[0-9a-fA-F]{2}`)

	// Find string that matches the pattern
	match := re.FindStringSubmatch(output)
	if match != nil && len(match) > 0 {
		// The first part of the match is the entire line
		// The first group (IP) is at index 0, and we need to split it
		parts := regexp.MustCompile(`\s+`).Split(output, -1)
		fmt.Println(parts[0])
		ipPrivate = parts[0]
	}
	ipPublicPrivate := [2]string{ipPublic, ipPrivate}
	return ipPublicPrivate
}
