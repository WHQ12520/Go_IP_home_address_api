package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/thinkeridea/go-extend/exnet"
	"github.com/xiaoqidun/qqwry"
)

func main() {

	var listen_addr string
	if len(os.Args) > 1 && os.Args[1] != "" {
		listen_addr = os.Args[1]
	} else {
		listen_addr = "0.0.0.0:12520"
	}

	var ipdate string
	if len(os.Args) > 2 && os.Args[2] != "" {
		ipdate = os.Args[2]
	} else {
		ipdate = "./ipdata/qqwry.dat"
	}

	// 从文件加载IP数据库
	if err := qqwry.LoadFile(ipdate); err != nil {
		fmt.Println("Error: ", ipdate, "地址数据库加载失败")
		panic(err)
	}

	http.HandleFunc("/", web_http_client)
	fmt.Println("INFO:监听地址", listen_addr)
	http.ListenAndServe(listen_addr, nil)
}

func web_http_client(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	if ip == "" {
		ip = exnet.ClientPublicIP(r)
	}
	ip_data, _ := qqwry.QueryIP(ip)
	fmt.Println(
		nowtime(),
		"IP:", ip_data.IP,
		"地区：", ip_data.Country,
		"省份：", ip_data.Province,
		"城市：", ip_data.City,
		"区县：", ip_data.District,
		"运营商：", ip_data.ISP,
	)
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		http.Error(w, "500 内部服务器错误", http.StatusInternalServerError)
		fmt.Println("Error:模版解析失败 ", err)
		return
	}
	data := map[string]string{
		"ip":       ip_data.IP,
		"Country":  ip_data.Country,
		"Province": ip_data.Province,
		"City":     ip_data.City,
		"District": ip_data.District,
		"ISP":      ip_data.ISP,
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "500 内部服务器错误", http.StatusInternalServerError)
		fmt.Println("Error:写入请求失败 ", err)
	}
}

func nowtime() string {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)
	return timeStr
}
