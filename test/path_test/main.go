package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func main()  {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err !=nil{
		zap.S().Info("错误：",err.Error())
	}
	//str := strings.Replace(dir, "\\", "/", -1)
	//str = strings.TrimRight(str, "/")
	dir = dir+"/files/*"
	fmt.Println("abs dir:",dir)
	files, _ := filepath.Glob(dir)
	fmt.Println("glob files:",files)
	for _,file := range files{
		fileInfo, _ := os.Stat(file)
		//fmt.Println("stat fileInfo:",fileInfo)
		mfile := filepath.Base(file)
		fmt.Println("base mfile:",mfile)
		mext := strings.ToUpper(path.Ext(mfile))
		fmt.Println("ext mext:",mext)

		//var i = 0

		m :=make(map[string]string)
		var ms []map[string]string

		m["name"] = fileInfo.Name()
		m["ext"] = mext
		m["size"] = strconv.FormatInt(fileInfo.Size(), 10)
		m["date"] = fileInfo.ModTime().Format("2006-01-02 15:04:05")
		m["path"] = "files/"+mfile
		ms = append(ms,m)
		fmt.Println("ms :",ms)
	}


}
