package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"encoding/json"
)
var ZabbixTarget,ZabbixMessage,ZabbixType,PrometheusAlertUrl string
var help bool
func init()  {
	//zabbix传入参数发送目标，告警消息，目标类型
	flag.StringVar(&ZabbixTarget, "t", "https://oapi.dingtalk.com/robot/send?access_token=xxxxx", "指定告警消息的接收目标的`手机号/钉钉url/微信url`")
	flag.StringVar(&ZabbixMessage, "m", "zabbix告警测试", "需要发送的`告警消息内容`")
	flag.StringVar(&ZabbixType, "type", "dd", "告警消息的目标类型,支持`dx(短信)、dh(电话)、dd(钉钉)、wx(微信)`")
	flag.StringVar(&PrometheusAlertUrl, "d", "http://127.0.0.1:8080/zabbix", "`PrometheusAlert的地址`")
	flag.BoolVar(&help,"h", false, "显示帮助")
	flag.Usage = usage
}
func usage() {
	fmt.Fprintf(os.Stderr, `Version 1.0 If you need help contact 244217140@qq.com or visit https://github.com/feiyu563/PrometheusAlert
Usage: zabbixclient [-h] [-t SendTarget] [-m SendMessage] [-type SendType] [-d PrometheusAlertUrl]
Example(发送告警到钉钉)：zabbixclent -t https://oapi.dingtalk.com/robot/send?access_token=xxxxx -m zabbix告警测试 -type dh -d http://127.0.0.1:8080/zabbix

Options:
`)
	flag.PrintDefaults()
}
func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}

	JsonPrometheusAlertMessage:=make(map[string]interface{})
	JsonPrometheusAlertMessage["zabbixtarget"]=ZabbixTarget
	JsonPrometheusAlertMessage["zabbixmessage"]=ZabbixMessage
	JsonPrometheusAlertMessage["zabbixtype"]=ZabbixType
	PostMessage, err := json.Marshal(JsonPrometheusAlertMessage)
	if err != nil {
		fmt.Println(err.Error() )
		return
	}

	Post(PrometheusAlertUrl,PostMessage)
}

func Post(url string,message []byte) {
	reader := bytes.NewReader(message)
	resp, err := http.Post(url,"application/json",reader)
	if err != nil {
		fmt.Printf("消息发送失败：%s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("消息发送失败：%s", err)
	}
	fmt.Printf("消息发送完成,服务器返回内容：%s", string(body))
}