package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	//    const jsonStream = `
	//    {"ip":"122.72.38.138","port":55336,"isp":"铁通|CRTC","initTime":1431762746174}
	//{"ip":"119.254.101.48","port":80,"isp":"华瑞信通|HUARUI","initTime":1431916106315}
	//{"ip":"101.4.136.67","port":80,"isp":"教育网|CERNET","initTime":1431931837004}
	//{"ip":"111.199.147.49","port":8118,"isp":"联通|UNICOM","initTime":1431950194471}
	//{"ip":"121.52.213.6","port":3128,"isp":"铜牛集团|CHINANET","initTime":1431951327294}
	//{"ip":"122.72.30.72","port":80,"isp":"铁通|CRTC","initTime":1431962587504}
	//{"ip":"210.14.158.122","port":80,"isp":"|OTHER","initTime":1431966095623}
	//{"ip":"122.72.30.199","port":80,"isp":"|CRTC","initTime":1431967907297}
	//{"ip":"101.4.136.66","port":80,"isp":"教育网|CERNET","initTime":1431968349160}
	//{"ip":"118.26.62.133","port":80,"isp":"|UNICOM","initTime":1431972713577}
	//{"ip":"101.4.136.65","port":81,"isp":"教育网|CERNET","initTime":1431972713577}
	//{"ip":"101.4.136.101","port":9999,"isp":"教育网|CERNET","initTime":1431975704457}
	//{"ip":"111.13.109.52","port":80,"isp":"移动|CMNET","initTime":1431977206064}
	//{"ip":"124.127.123.48","port":80,"isp":"电信|CHINANET","initTime":1431979156587}
	//{"ip":"221.176.14.72","port":80,"isp":"移动|CMNET","initTime":1431983830214}
	//{"ip":"210.31.15.35","port":80,"isp":"教育网|CERNET","initTime":1431987648645}
	//{"ip":"111.202.56.43","port":80,"isp":"联通|UNICOM","initTime":1431988973966}
	//{"ip":"101.4.136.103","port":9999,"isp":"教育网|CERNET","initTime":1431992183664}
	//{"ip":"124.202.174.150","port":8118,"isp":"鹏博士|DXTNET","initTime":1431997460056}
	//{"ip":"106.2.212.195","port":80,"isp":"|OTHER","initTime":1431997759875}
	//{"ip":"114.113.221.146","port":80,"isp":"|UNICOM","initTime":1431997909793}
	//{"ip":"124.16.131.101","port":18186,"isp":"中国科技网|CSTNET","initTime":1431999182671}
	//{"ip":"124.202.175.230","port":9797,"isp":"鹏博士|DXTNET","initTime":1432000092761}
	//{"ip":"114.246.151.104","port":8118,"isp":"联通|UNICOM","initTime":1432000609567}
	//{"ip":"58.132.25.19","port":3128,"isp":"|CHINANET","initTime":1432001277296}
	//{"ip":"106.37.177.251","port":3128,"isp":"电信|CHINANET","initTime":1432001411622}
	//{"ip":"124.200.38.46","port":8118,"isp":"鹏博士|DXTNET","initTime":1432001411622}
	//{"ip":"114.113.221.166","port":9999,"isp":"|UNICOM","initTime":1432002575500}
	//{"ip":"123.121.88.67","port":8118,"isp":"联通|UNICOM","initTime":1432003332333}
	//{"ip":"39.176.78.163","port":8123,"isp":"移动|CMNET","initTime":1432003332334}
	//{"ip":"114.255.183.189","port":8080,"isp":"联通|UNICOM","initTime":1432003332333}
	//{"ip":"114.255.183.173","port":8080,"isp":"联通|UNICOM","initTime":1432003332333}
	//    `
	const jsonStream = `
	[{"北京":29},{"重庆":24},{"昆明":26},{"哈尔滨":5},{"天津":17},{"成都":27},{"太原":8},{"深圳":30},{"青岛":6},{"佛山":14},{"苏州":13},           {"杭州":23},{"郑州":19},{"福州":9},{"无锡":8},{"长沙":9},{"沈阳":8},{"广州":30},{"温州":10},{"大连":4},{"宁波":9},{"贵阳":16},{"南宁":3},      {"上海":30},{"西安":9},{"长春":5},{"南京":38},{"常州":4},{"东莞":27},{"武汉":26},{"石家庄":3},{"厦门":8},{"合肥":8},{"济南":13}]
	`
	type Message struct {
		Ip       string
		Port     int
		Isp      string
		InitTime int
	}
	fmt.Println("hehe1")
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	fmt.Println("hehe2")
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		fmt.Printf("%s: %d\n", m.Ip, m.Port)
	}
}
