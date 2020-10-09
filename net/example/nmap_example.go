package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Ullaakut/nmap"
	"github.com/gakkiyomi/galang/net"
	"github.com/songtianyi/rrframework/logs"
)

func main() {
	sc, err := net.NMAP.NewScanner("192.168.1.146", "fangcong")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	if err != nil {
		logs.Error("unable to create nmap scanner: %v", err)
	}

	do := make(chan *net.Scanner, 1)

	go func() {

		logs.Info("Runnning")

		sc.Status = net.Running
		sc, err = sc.Scanner(
			nmap.WithContext(ctx),
		)
		if err != nil {
			logs.Error("scan error : %v", err)
		}
		do <- sc
	}()

	res := <-do

	b2, _ := json.Marshal(res.Hosts())
	logs.Info(string(b2))

}
