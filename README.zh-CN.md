
## 💡 简介

[Galang](https://github.com/gakkiyomi/galang) 是一款包含了网络地址相关，字符串相关，算法和数据结构等 Go 语言常用工具库。

> **[EN README](README.md)**

## 使用

   1. 直接在go.mod中引用
   2. `GOPROXY=direct go get -u github.com/gakkiyomi/galang`

## ✨ 功能

### 网络相关 `net.Network`

* 获取本地IP地址
* 获取公网IP地址
* JAVA Apache SubnetUtils go实现
* 通过子网掩码转换掩码位元数
* 通过掩码位元数转换子网掩码
* Long类型IP转IP字符串
* IP字符串转Long类型IP
* 判断一个IP地址是否在一个网段内
* 获取本地机器唯一标识UUID(linux)
* 获取一个网段下所有可用地址(排除广播地址和网络地址)
* IP最长前缀匹配算法(LPM)
* 获取网段的网络位和广播位
* 获取一个网段中的第一个可用地址和最后一个可用地址
* 获取一个网段中的可用地址数

### SNMP `net.SNMP`

* 获取系统信息(主机名，系统描述，snmp启动以来的运行时间，联系人，物理位置，厂商)
* 获取接口信息(IP地址，子网掩码，接口当前状态，MAC地址)

### NMAP `net.NMAP`

* 基于NMAP的网络扫描器(支持扫描网段或者范围ip来获取主机ip和端口开闭情况)

### 字符串 `string.String`

* 字符串是否已xx开头
* 字符串是否已xx结尾
* 判断包含字串忽略大小写
* 判断字符串是否为空白字符串
* 判断字符串是否包含任意空白字符串
* java StringBuilder go版本实现

### 数组 `array.Array`

* 数组中指定位置插入元素
* 去除数组中的重复元素
* 字符串转成字符数组
* 反转字符串数组
* 获取int数组里面最大的值
* 获取int数组里面最小的值
* 获取两个int数组的并集
* 获取两个string数组的并集
* 获取两个int数组的交集
* 获取两个string数组的交集
* 获取两个int数组的差集
* 获取两个string数组的差集
* 二分查找数组里面的值(请确保数组已排序)

### 数据结构 `structure`

* 栈
* 队列
* 环状缓冲区
* 二叉树 **不保证平衡**
* 堆(heap,大顶堆,小顶堆)
* TODO 红黑树
* TODO 图

### 排序算法 `sort`

* 冒泡排序
* 选择排序
* 插入排序
* 快速排序
* 堆排序
* 归并排序
* 希尔排序
* 桶排序

### 读取配置文件 `config`

* 读取JSON格式的配置文件
* 读取XML格式的配置文件
* TODO 读取YMAL格式的配置文件

### 文件操作 `file.File`

* 获取文件大小
* 判断路径是否存在
* 判断文件内容是否为JSON格式
* 判断文件内容是否为XML格式
* 判断文件流是否为JSON格式
* 判断文件流是否为XML格式

### 转换 `utils.Transform`

* 封装字符串与其他类型相互转换

### UUID `utils.UUID`

* 生成UUID
* 检查一个字符串是否为UUID

### 🔑 JetBrains 开源证书支持

`galang` 项目一直以来都是在 JetBrains 公司旗下的 `GoLand` 集成开发环境中进行开发，基于 free JetBrains Open Source license(s) 正版免费授权，在此表达我的谢意。

<a href="https://www.jetbrains.com/?from=galang" target="_blank"><img src="https://b3logfile.com/file/2021/05/jetbrains-variant-2-42d96aa4.png" width="250" align="middle"/></a>

## 特别感谢

[gosnmp # 不错的SNMP客户端](https://github.com/alouca/gosnmp)

[nmap # 牛逼的NMAP客户端](https://github.com/Ullaakut/nmap)

[dmidecode # dmidecode解析类库](https://github.com/dselans/dmidecode)

[etree # xml解析类库](https://github.com/beevik/etree)
