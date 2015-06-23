package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
	"net"
	"sync/atomic"
	"regexp"
	"errors"
)

var wg    = sync.WaitGroup{}
var right =int64(0)
var wrong =int64(0)
var ErrFindPublicIp = errors.New("can'tfind public ip")
var publicipmap = make(map[string]bool)

var timeout = time.Duration(3 * time.Second)
//var timeout = time.Duration(1 * time.Microsecond)
func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

func fetchUrl(proxyurl string)  {
	defer wg.Done()
//	defer fmt.Printf("now:%s, wrong:%d, right:%d \n", time.Now(), wrong, right)

	proxyUrl, err := url.Parse(proxyurl)
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl), Dial: dialTimeout,}, Timeout: timeout, }
	req, err := http.NewRequest("GET", "http://city.ip138.com/ip2city.asp", nil)
	if err != nil {
		atomic.AddInt64(&wrong, 1)
		fmt.Printf("now:%s, fetch error :%s \n", time.Now(), proxyurl)
		return
	}
	resp, err := client.Do(req)

	if err != nil {
		// handle error
		fmt.Printf("now:%s, fetch error :%s \n", time.Now(), proxyurl)
		atomic.AddInt64(&wrong, 1)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Printf("now:%s, fetch error :%s \n", time.Now(), proxyurl)
		atomic.AddInt64(&wrong, 1)
		return
	}
	//resLen := len(body)
	//if !strings.Contains(string(body), "http://s.ip-cdn.com/img/logo.gif"){
	//	fmt.Printf("now:%s, fetch error :%s body:%s\n", time.Now(), proxyurl, string(body))
	//	atomic.AddInt64(&wrong, 1)
	//	return
	//}

	publicip, err := findPublicIP(body)
	if err != nil{
		fmt.Printf("now:%s, fetch error :%s body:%s\n", time.Now(), proxyurl, string(body))
                atomic.AddInt64(&wrong, 1)
                return
	}
	publicipmap[publicip] = true

	atomic.AddInt64(&right, 1)
	fmt.Printf("now:%s, fetch right:%s \n", time.Now(), proxyurl)
	fmt.Printf("publicip:%s\n", publicip)
	fmt.Printf("after now:%s, wrong:%d, right:%d, wg:%s \n", time.Now(), wrong, right, wg)
}

// 获取某个代理的公网地址
func findPublicIP(body []byte) (string, error) {
    re := regexp.MustCompile(`\[(.*)\]`)
    //re := regexp.MustCompile(`\[(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9])\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[0-9])\]`)
    matchResult := re.FindStringSubmatch(string(body))
    if len(matchResult) == 0 || len(matchResult) > 15{
        return "", ErrFindPublicIp
     }
    return matchResult[1], nil


}

func fetchProxy(url_map map[string]string) {
	for source, value := range url_map{

		resp, err := http.Get(value)
		if err != nil {
			// handle error
		}

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		resBody := string(body)
		subReses := strings.Split(resBody, "|")
	//	wg.Add(len(subReses))

		num:= len(subReses)
	//	wg.Add(num)
		for i := 0; i < num; i++ {
			subRes := strings.Split(subReses[i], ",")[0]
			proxyurl := fmt.Sprintf("http://%s", subRes)
			wg.Add(1)
			go fetchUrl(proxyurl)
			if (i + 1)%100 == 0{
				time.Sleep(timeout)
			}

		}
		fmt.Println(wg)
		fmt.Printf("publicip source:%s, urls:%d\n", source, len(publicipmap))
	}
	wg.Wait()
}

func main() {
	url_map := make(map[string]string)
	url_map[""] = ""
	fmt.Printf("start now:%s, wrong:%d, right:%d \n", time.Now(), wrong, right)
	fetchProxy(url_map)
	fmt.Printf("last now:%s, wrong:%d, right:%d, publicip:%d \n", time.Now(), wrong, right, len(publicipmap))
}
