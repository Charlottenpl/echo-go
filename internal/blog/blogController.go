package blog

import (
	"echo-go/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetById(c *gin.Context) {

	blog, err := sql.GetById(1)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": blog})
	}

}

func Find(c *gin.Context) {
	// , title string, bt time.Time, et time.Time, content string, tag string, _type string, status int, limit, offset int
	var req struct {
		Title   string `form:"title" binding:"required"`
		Bt      string `form:"begin_time"  binding:"required"`
		Et      string `form:"end_time"   binding:"required"`
		Content string `form:"content" binding:"required"`
		Type    string `form:"type" binding:"required"`
		Status  int    `form:"status" binding:"required"`
		Limit   int    `form:"limit" binding:"required"`
		Offset  int    `form:"offset" binding:"required"`
	}

	// 将请求参数绑定到结构体中
	if err := c.BindJSON(&req); err != nil {
		c.JSON(502, gin.H{"error": err.Error()})
		return
	}

	blogs, err := sql.Find(req.Title, req.Bt, req.Et, req.Content, req.Type, req.Status, req.Limit, req.Offset)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": blogs})
	}
}

func Page(c *gin.Context, page, size int) {
	blogs, err := sql.ListPage(page, size)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": blogs})
	}
}
