package initalize

import (
	"filexxx/global"
	"flag"
	"fmt"
	"os"
)

func InitArgs(){
	fmt.Println("Filexxx")
	fmt.Println("-p 指定端口打开")
	fmt.Println("-v 查看版本号")
	flag.IntVar(&global.MyConfig.Port, "p", 8090, "打开端口号 默认为8090")
	versionFlag := flag.Bool("V", false, "显示版本号")
	flag.Parse()
	if *versionFlag {
		fmt.Println("Version 1.0")
		os.Exit(0)
	}
}
