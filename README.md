
## 💡 简介

[Galang](https://github.com/gakkiyomi/galang) 是一款 Go 语言常用工具库。

## ✨ 功能

### 文件操作 `file.File`

* 获取文件大小
* 判断路径是否存在
* 判断文件内容是否为JSON格式
* 判断文件内容是否为XML格式
* 判断文件流是否为JSON格式
* 判断文件流是否为XML格式

### 网络相关 `net.Net`

* 子网掩码转CIDR
* CIDR转子网掩码
* Long类型IP转IP字符串
* IP字符串转Long类型IP
* 判断一个IP地址是否在一个网段内
* 获取本机dmidecode唯一标识UUID(linux)
* 获取一个网段下所有可用地址(排除广播地址和网络地址)

### 字符串 `utils.String`

* 字符串是否已xx开头
* 字符串是否已xx结尾
* 判断包含字串忽略大小写

### 转换 `utils.Transform`

* 封装字符串与其他类型相互转换
