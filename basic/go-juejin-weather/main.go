package main

import (
	"fmt"
	"go-juejin-weather/juejin.cn/weather"
)

func main() {
	res := weather.CurrentWeather("101010100")
	fmt.Println(res)
}
