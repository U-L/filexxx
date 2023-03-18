package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"path"
	"path/filepath"
)


func UploadFile(c *gin.Context) {
	form, _ := c.MultipartForm()
	f := form.File["uploads"]
	for _, file := range f {
		zap.S().Infof("上传文件：%s",file.Filename)
		dst := path.Join("./tmp/upload/", file.Filename)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func ListFile(c *gin.Context) {
	//dir := GetAbsFile()+"/tmp/upload/*" //产品环境
	dir := GetAbsFile()+"/files/*" //开发环境
	zap.S().Info(dir)
	data := GetFileData(dir)
	c.JSON(http.StatusOK,gin.H{
		"msg":"文件列表",
		"data":data,
	})
}

func DownloadFile(c *gin.Context) {
	if pa := c.Param("name"); pa != "" {
		target := filepath.Join(GetAbsFile()+"/files/", pa)
		fmt.Println(target)
		//c.Header("Content-Description", "File Transfer")
		//c.Header("Content-Transfer-Encoding", "binary")
		//c.Header("Content-Disposition", "attachment; filename="+pa)
		//c.Header("Content-Type", "application/octet-stream")
		c.File(target)
	} else {
		c.Status(http.StatusNotFound)
	}}