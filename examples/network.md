# network example

## SubnetInfo

~~~go
package main

import (
	"fmt"

	"github.com/gakkiyomi/galang/net"
)

func main() {
	info, _ := net.NewSubnetInfo("192.168.1.0/24")
	fmt.Println(info.IsRangeOf(`192.168.1.22`))
	fmt.Println(info.AddressString())
	fmt.Println(info.NetmaskString())
	fmt.Println(info.NetworkString())
	fmt.Println(info.HighAddress())
	fmt.Println(info.LowAddress())
	fmt.Println(info.Size())
	fmt.Println(info.GetCidrSignature())
	fmt.Println(info.BradcastString())
	fmt.Println(info.ToString())
}
~~~

## IP

~~~go

import (
	"fmt"

	"github.com/gakkiyomi/galang/net"
)

func main() {
addri, _ := net.Network.IP2long(`192.168.1.4`)

addr := net.Network.Long2ip(12222)

longest, _ := net.Network.LPM("192.168.1.4", []string{"192.168.1.0/24", "192.168.0.0/16"})

b, _ := net.Network.IsRangeOf(`192.168.1.2`, `192.168.2.0/24`)
}
~~~