package weather

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CurrentWeather(cityCode string) string {
	resp, err := http.Get("https://devapi.qweather.com/v7/weather/now?location=" +
		cityCode + "&key=17a5416c0fd742848e7b06c32e08d0db")
	if err != nil {
		fmt.Println("HTTP请求失败：", err)
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取网络响应失败：", err)
		panic(err)
	}
	return string(body)
}
