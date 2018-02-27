package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age int
}

type IPAddr [4]byte

func (ip IPAddr) String() string{
	var strIp string
	for _,v := range ip {
		strIp = strIp + strconv.Itoa(int(v)) + "."
	}

	strIp = strIp[0:len(strIp)-1]
	return fmt.Sprintf(strIp)
}

func (p Person) String() string{
	return fmt.Sprintf("%v (%v years)\n",p.Name,p.Age)
}
func main() {
	a := Person{"Arthur Dent",42}
	z := Person{"Zaphod Beeblebrox",9001}
	fmt.Println(a,z)

	hosts := map[string]IPAddr{
		"loopback": {127,0,0,1},
		"googleDNS":{8,8,8,8},
	}

	for name, ip :=range hosts {
		fmt.Printf("%v: %v\n",name,ip)
	}
}
