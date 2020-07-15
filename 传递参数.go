package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main(){
	var host string
	// &host 就是接收命令行中 -h后面的参数值
	flag.StringVar(&host, "h", "localhost", "主机名,默认localhost")
	//解析命令行参数写入注册的flag里
	flag.Parse()
	//输出结果
	var url strings.Builder
	url.WriteString("http://")
	url.WriteString(host)

	//resp,err := http.Get("http://10.1.1.147:9000/health")
	resp,err := http.Get(url.String())
	if err != nil {
		fmt.Println("get faile,err: ", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read from resp.Body failed err: ", err)
		return
	}
	var json_str string = string(body)
	var data map[string]interface{}
	json.Unmarshal([]byte(json_str), &data)
	//fmt.Println(data)
	fmt.Println(data["status"])

}
