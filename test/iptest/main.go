package main

import (
	"fmt"
	"go.uber.org/zap"
	"net"
)

func GetIpAddr() (ipaddr []string){
	ifa,err := net.InterfaceAddrs()
	if err != nil {
		zap.S().Info("获取本地IP地址失败",err.Error())
	}
	for _,address := range ifa{
		if ipnet,ok := address.(* net.IPNet);ok && !ipnet.IP.IsLoopback(){
			if ipnet.IP.To4() != nil{
				ipaddr = append(ipaddr,ipnet.IP.String())
			}
		}
	}
	return ipaddr
}



func main()  {
	fmt.Println(GetIpAddr())
}