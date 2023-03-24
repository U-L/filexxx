package file

import (
	"go.uber.org/zap"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

func GetAbsFile() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		zap.S().Info("文件路径错误：", err.Error())
	}
	return dir
}

func GetFileData(dir string) []map[string]string {
	files, _ := filepath.Glob(dir)
	i := 0
	var ms []map[string]string
	for _, file := range files {
		fileInfo, _ := os.Stat(file)
		pathExt := path.Ext(fileInfo.Name())
		//裁切字符串
		viewName := fileInfo.Name()
		if len(viewName) > 30 {
			viewNameRune := []rune(viewName)
			viewName = string(viewNameRune[:30]) + pathExt
		}
		dMap := make(map[string]string)
		dMap["id"] = strconv.Itoa(i)
		dMap["name"] = fileInfo.Name()
		dMap["viewName"] = viewName
		dMap["ext"] = pathExt
		dMap["date"] = fileInfo.ModTime().Format("2006-01-02 15:04:05")
		dMap["path"] = "tmp/uploads/" + fileInfo.Name()
		dMap["size"] = strconv.FormatInt(fileInfo.Size()/1024, 10)
		ms = append(ms, dMap)
		i++
		//zap.S().Info("ms :",ms)
	}
	return ms
}
