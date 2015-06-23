package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// IT提供的代理的来源，ip地址，以及负载设置
const (
	CityADSLURL   = ""
	DefaultHost   = ""
	CityADSLLimit = 20
	CityProxyType = "socks5"
)

// 把城市代理加载出来，返回所有的地址
func loadCityADSL() ([]string, error) {
	resp, err := http.Get(CityADSLURL)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	addrs := make([]string, 0)

	for _, line := range strings.Split(string(body), "\n") {
		line = strings.TrimSpace(line)
		tokens := strings.Split(line, ",")
		if len(tokens) == 2 {
			addrs = append(addrs, fmt.Sprintf("%s:%s", DefaultHost, tokens[1]))
		}
	}

	return addrs, nil
}
