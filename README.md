
## 💡 简介

[Galang](https://github.com/gakkiyomi/galang) Some utils for the Go: network address,string,algorithms/structure,array/silce

> **[中文说明](README.zh-CN.md)**

## use

   1. use go.mod
   2. `GOPROXY=direct go get -u github.com/gakkiyomi/galang`

## ✨ Feature List

### Network `net.Network`

* Get Local IP
* Get Internet IP
* JAVA Apache SubnetUtils go version
* Convert the number of mask bits by subnet mask
* Convert the subnet mask by the number of mask bits
* Long type IP to IP string
* IP string to Long type IP
* Determine whether an IP address is range of a network segment
* Get local machine uuid (only linux)
* Get all available addresses under a network segment (excluding broadcast addresses and network addresses)
* IP longest prefix matching algorithm (LPM)
* Get the network bit and broadcast bit of the network segment
* Get the first available address and the last available address in a network segment
* Get the number of available addresses in a network segment

### SNMP `net.SNMP`

* Get system information (host name, system description, running time since snmp started, contact person, physical location, vendor)
* Get interface information (IP address, subnet mask, current interface status, MAC address)

### NMAP `net.NMAP`

* NMAP-based network scanner (support scanning network segment or range ip to obtain host ip and port opening and closing status)

### String `string.String`

* Does the string begin with xx
* Does the string end with xx
* Judgment contains string ignore case
* Determine whether the string is a blank string
* Determine whether the string has contains any blank string
* Java StringBuilder in go

### Array `array.Array`

* Insert an element at the specified position in the array
* Array deduplication
* String to String array
* Reverse array/silce
* Get max value from int array
* Get min value from int array
* Get unionsection from int array
* Get unionsection from string array
* Get intersection from int array
* Get intersection from string array
* Get complement from int array
* Get complement from string array
* BinarySerch(array must be sorted by asc)

### Structure `structure`

* Stack
* Queue
* RingBuffer
* BitMap
* BinaryTree **No guarantee of balance**
* Heap,(MaxHeap,MinHeap)
* TODO | Red–black tree
* TODO | Graph

### Sort `sort`

* Quick Sort
* Selection Sort
* Insertion Sort
* Quick Sort
* Heap Sort
* Merge Sort
* Shell Sort
* Bucket Sort

### Config `config`

* Read configuration file in JSON format
* Read configuration file in XML format
* TODO | Read configuration file in YAML format

### File `file.File`

* Get file size
* Check the path exists
* Check the file content is in JSON format
* Check the file content is in XML format
* Check the file stream content is in JSON format
* Check the file stream content is in XML format

### Transform `utils.Transform`

* Convert between string and other types

### UUID `utils.UUID`

* Generate a UUID
* check if UUID is legal

### 🔑 JetBrains OS licenses

`galang` had been being developed with `GoLand IDE` under the free JetBrains Open Source license(s) granted by JetBrains s.r.o., hence I would like to express my thanks here.

<a href="https://www.jetbrains.com/?from=galang" target="_blank"><img src="https://b3logfile.com/file/2021/05/jetbrains-variant-2-42d96aa4.png" width="250" align="middle"/></a>

## Thanks

[gosnmp # good SNMP client](https://github.com/alouca/gosnmp)

[nmap # great NMAP client](https://github.com/Ullaakut/nmap)

[dmidecode # dmidecode library](https://github.com/dselans/dmidecode)

[etree # xml parse library](https://github.com/beevik/etree)
