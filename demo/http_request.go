package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	url1 = "https://www.tinywan.com/api/open/returnUrl"
	url2 = "https://restapi.amap.com/v3/weather/weatherInfo?key=d2315f3b0b4e57bbf5428e755a73e692&city=110101"
)

func main() {
	res, err := doHttpGetRequest(url2)
	if err != nil {
		fmt.Println("net req error")
	} else {
		fmt.Println(res)
	}
}

// http get请求函数
// 输入参数：请求url
// 返回值：res，天气数据。err，错误信息
func doHttpGetRequest(url string) (res string, err error) {

	// http.Get在net/http中，所以要import "net/http"
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	} else {
		// 使用efer resp.Body.Close()。当doHttpGetRequest成功return之后，执行此行语句。多用于句柄关闭
		defer resp.Body.Close()

		// io流数据读取。需要引用io/ioutil
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return "", err
		} else {
			return string(body), err
		}
	}
}
