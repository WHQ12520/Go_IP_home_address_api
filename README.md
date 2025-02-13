# IP 地址归属地查询 API

这是一个用 Go 编写的查询 IPv4 地址归属地的 API。

## IP 数据库

- IP 数据库使用了 QQ 纯真数据库
- 使用前请下载 IP 数据库，放置到默认路径 `./ipdata/qqwry.dat`
- 或者在参数中传入 IP 数据库路径
- IP 数据库下载地址：[https://github.com/nmgliangwei/qqwry](https://github.com/nmgliangwei/qqwry) （感谢数据库作者的辛勤维护~）

## 运行配置

- 默认监听地址：`0.0.0.0:12520`
- 默认数据库路径  `./ipdata/qqwry.dat`
- 可以在第一个参数中配置监听地址
- 可以在第二个参数中传入 IP 数据库路径
- 例子：`./Go_IP_home_address_api 0.0.0.0:1000 ./qqwry.dat`

## 调用方法

1. 请求不带参数，返回 API 调用者的 IP 地址和 IP 对应归属地
2. 请求携带参数，返回参数中的 IP 地址和 IP 对应归属地

## 示例

- 不带参数的请求：
    ```sh
    curl http://localhost:12520/
    ```
- 带参数的请求：
    ```sh
    curl http://localhost:12520/?ip=1.1.1.1
    ```

## 依赖

- [thinkeridea/go-extend](https://github.com/thinkeridea/go-extend)
- [xiaoqidun/qqwry](https://github.com/xiaoqidun/qqwry)

- 在以源代码运行前请确保安装这些依赖：
    ```sh
    go get github.com/thinkeridea/go-extend/exnet
    go get github.com/xiaoqidun/qqwry
    ```