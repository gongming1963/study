package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 流量组提供的代理的来源，ip地址，以及负载设置
const (
	TaobaoADSLStat  = "" // 城市信息
	TaobaoADSLUrl   = ""     // 具体的代理信息
	TaobaoADSLLimit = 20
	TaobaoProxyType = "http"
)

// 返回代理的城市列表
func getCityList() ([]string, error) {
	body, err := gethtml(TaobaoADSLStat)
	if err != nil {
		fmt.Printf("failed to load city list of taobao proxy: %s", err.Error())
		return nil, err
	}
	var cities []map[string]int

	citylist := make([]string, 0)

	json.Unmarshal([]byte(body), &cities)
	for _, value := range cities {
		for city, num := range value {
			citylist = append(citylist, city)
			fmt.Printf("taobao city:%s,  num:%d", city, num)
		}
	}

	return citylist, nil
}

// 返回页面
func gethtml(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

// 从一个url得到一个proxy
func getProxy(url string) ([]string, error) {
	body, err := gethtml(url)
	if err != nil {
		fmt.Printf("failed to load url:url, err: %s", url, err.Error())
		return nil, err
	}

	type Proxy struct {
		Ip       string
		Port     int
		Isp      string
		InitTime int
	}

	var y []Proxy

	addrs := make([]string, 0)
	json.Unmarshal(body, &y)
	for _, value := range y {
		addrs = append(addrs, fmt.Sprintf("%s:%d", value.Ip, value.Port))
	}
	return addrs, nil
}

// 把城市代理加载出来，返回所有的地址
func loadTaobaoADSL() ([]string, error) {
	url, err := url.Parse(TaobaoADSLUrl)
	if err != nil {
		fmt.Println(err)
	}
	query := url.Query()
	query.Set("all", "true")

	proxyaddrs := make([]string, 0)
	citylist, err := getCityList()
	if err != nil {
		return nil, err
	}

	for _, city := range citylist {
		query.Set("city", city)
		url.RawQuery = query.Encode()

		tmpProxyAddrs, err := getProxy(url.String())
		if err != nil {
			return nil, err
		}

		fmt.Printf("get %d proxies from city:%s", len(tmpProxyAddrs), city)

		for _, addr := range tmpProxyAddrs {
			proxyaddrs = append(proxyaddrs, addr)
		}
	}
	fmt.Printf("get %d proxies totally", len(proxyaddrs))
	return proxyaddrs, nil
}
