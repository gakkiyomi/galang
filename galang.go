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
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Ullaakut/nmap"
	"github.com/gakkiyomi/galang/net"
	"github.com/songtianyi/rrframework/logs"
)

func main() {

	sc, err := net.NMAP.NewScanner("192.168.1.222", "fangcong")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	if err != nil {
		logs.Error("unable to create nmap scanner: %v", err)
	}

	b, _ := json.Marshal(sc)

	logs.Info("===========================================================")
	logs.Info(string(b))
	logs.Info("===========================================================")

	do := make(chan *net.Scanner, 1)

	go func() {
		logs.Info("===========================================================")
		logs.Info("Runnning")
		logs.Info("===========================================================")
		sc.Status = net.Running
		sc, err = sc.Scanner(
			nmap.WithPorts("80", "8080", "5432"),
			nmap.WithContext(ctx),
		)
		if err != nil {
			logs.Error("scan error : %v", err)
		}
		do <- sc
	}()

	res := <-do

	result := res.Result
	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}
		fmt.Printf("addresses %q ", host.Addresses)
		fmt.Printf("Host %q , type %q ,verdor %q:\n", host.Addresses[0].Addr, host.Addresses[0].AddrType, host.Addresses[0].Vendor)
		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}

	fmt.Printf("Nmap done: %d hosts up scanned in %3f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)

	b2, _ := json.Marshal(sc)
	logs.Info("===========================================================")
	logs.Info(string(b2))
	logs.Info("===========================================================")

}
