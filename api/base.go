package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetView(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
