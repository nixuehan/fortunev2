package lib

import (
	"fmt"
	"net/http"
	"strings"
)

type Alarm struct {}

const (
	WebHook = `https://oapi.dingtalk.com/robot/send?access_token=fe5089dd78ee5198f4ec64380651ee8127ccf2295503ca688c64986128459a3e`
	KeyWork = `财神到`
)

func (alarm *Alarm) Ring(msg string)  {
	//请求地址模板
	webHook := WebHook
	content := fmt.Sprintf(`{"msgtype": "text","text": {"content": "%s: `+ msg + `"}}`,KeyWork)
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	E(err)

	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//发送请求
	resp, err := client.Do(req)
	defer resp.Body.Close()
	E(err)
}
