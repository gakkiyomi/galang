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

import "github.com/gakkiyomi/galang/net"

func main() {
	sbnts := []string{`192.168.1.0/24`, `192.168.0.0/16`, `192.0.0.0/8`, `192.168.2.0/24`}
	net.Network.LPM("192.168.1.2", sbnts)
}
