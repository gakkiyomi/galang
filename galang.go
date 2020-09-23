// Galang - Golang common utilities
// Copyright (c) 2019-present, gakkiiyomi@gamil.com
//
// gakkiyomi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package main

import (
	"fmt"

	"github.com/alouca/gosnmp"
	"github.com/gakkiyomi/galang/net"
)

func main() {
	wrapper, err := net.SNMP.NewSnmpWrapper("192.168.1.222", "public", gosnmp.Version2c, 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wrapper.SystemInfo())
	/*p, err := wrapper.Client.Get(".1.3.6.1.4.1.8072.3.2.10")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range p.Variables {
		fmt.Println(v)
	}*/

}
