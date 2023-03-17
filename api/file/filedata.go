package file

import (
	"go.uber.org/zap"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

func GetAbsFile() string{
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err !=nil{
		zap.S().Info("文件路径错误：",err.Error())
	}
	return dir
}
func GetFileData(dir string) []map[string]string{
	files, _ := filepath.Glob(dir)
	var ms []map[string]string
	for _,file := range files{
		fileInfo, _ := os.Stat(file)
		//pathBase := filepath.Base(file)
		pathExt := path.Ext(fileInfo.Name())

		dMap :=make(map[string]string)

		dMap["name"] = fileInfo.Name()
		dMap["ext"] = pathExt
		dMap["size"] = strconv.FormatInt(fileInfo.Size(), 10)
		dMap["date"] = fileInfo.ModTime().Format("2006-01-02 15:04:05")
		dMap["path"] = "files/"+fileInfo.Name()
		ms = append(ms,dMap)
		//zap.S().Info("ms :",ms)
	}
	return ms
}