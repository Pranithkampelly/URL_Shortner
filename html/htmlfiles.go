package html

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Getting(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome-template.html", gin.H{})
}
func Custom(c *gin.Context) {
	c.HTML(http.StatusOK, "custom.html", gin.H{})
}
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "root.html", gin.H{})
}
func Bulkupload(c *gin.Context){
	c.HTML(http.StatusOK, "bulk_upload.html", gin.H{})

}