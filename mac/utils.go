package mac

import (
	"fmt"
)

func GiveResult(ips []string) [][2]string {
	var results [][2]string
	for _, ip := range ips {
		output := ExecuteArping(ip)
		ipMac := FindMacaddress(output)
		ipPublicPrivate := MapIP(ipMac)
		results = append(results, ipPublicPrivate)
	}
	fmt.Println(results)
	return results
}
