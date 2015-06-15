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
)

var wg    = sync.WaitGroup{}
var right =int64(0)
var wrong =int64(0)

var timeout = time.Duration(4 * time.Second)
//var timeout = time.Duration(1 * time.Microsecond)
func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

func fetchUrl(proxyurl string)  {
	defer wg.Done()
//	defer fmt.Printf("now:%s, wrong:%d, right:%d \n", time.Now(), wrong, right)

	proxyUrl, err := url.Parse(proxyurl)
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl), Dial: dialTimeout,}}
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		atomic.AddInt64(&wrong, 1)
		return
	}
	resp, err := client.Do(req)

	if err != nil {
		// handle error
		atomic.AddInt64(&wrong, 1)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		atomic.AddInt64(&wrong, 1)
		return
	}
	resLen := len(body)
	if resLen < 90000 {
		atomic.AddInt64(&wrong, 1)
	}
	atomic.AddInt64(&right, 1)
}


func fetchProxy() {
	resp, err := http.Get(KUAIDAILI_URL)
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
		if i == num{
			break
		}

	}
	fmt.Println(wg)
	wg.Wait()
}

func main() {
	fmt.Printf("start now:%s, wrong:%d, right:%d \n", time.Now(), wrong, right)
	fetchProxy()
	fmt.Printf("last now:%s, wrong:%d, right:%d \n", time.Now(), wrong, right)
}
