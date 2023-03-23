package file

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func UploadFile(c *gin.Context) {
	form, _ := c.MultipartForm()
	f := form.File["uploads"]
	for _, files := range f {
		zap.S().Infof("上传文件：%s", files.Filename)
		//缩短文件名
		//if len(files.Filename) > 60 {
		//	files.Filename=files.Filename[:60]
		//}
		dst := path.Join("./tmp/upload/", files.Filename)
		err := c.SaveUploadedFile(files, dst)
		if err != nil {
			zap.S().Info("保存文件错误！：err%s", err.Error())
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func ListFile(c *gin.Context) {
	dir := GetAbsFile() + "/tmp/upload/*" //产品环境
	//dir := GetAbsFile()+"/files/*" //开发环境
	//zap.S().Info(dir)
	data := GetFileData(dir)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "文件列表",
		"data": data,
	})
}

func DownloadFile(c *gin.Context) {
	if pa := c.Param("name"); pa != "" {
		//target := filepath.Join(GetAbsFile()+"/files/", pa)
		target := filepath.Join(GetAbsFile()+"/tmp/upload/", pa)
		c.File(target)
	} else {
		c.Status(http.StatusNotFound)
	}
}

func RemoveFile(c *gin.Context) {
	if pa := c.Param("name"); pa != "" {
		//target := filepath.Join(GetAbsFile()+"/files/", pa)
		target := filepath.Join(GetAbsFile()+"/tmp/upload/", pa)
		err := os.Remove(target)
		if err != nil {
			zap.S().Info("删除文件失败")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "删除文件失败",
			})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "未找到文件",
		})
	}
}
