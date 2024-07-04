package blog

import (
	"echo-go/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Get(c *gin.Context) {
	var data struct {
		Data struct {
			ID   int    `json:"id"`
			Type string `json:"type"`
		}
	}

	// 绑定JSON请求体到结构体
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if data.Data.ID != 0 {

	}
	blog, err := sql.GetById(data.Data.ID)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": blog})
	}

}

func GetByType(c *gin.Context) {
	var data struct {
		Data struct {
			Type int `json:"type"`
		}
	}

	// 绑定JSON请求体到结构体
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if data.Data.Type != 0 {

	}
	blog, err := sql.GetByType(data.Data.Type)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": blog})
	}

}

func GetById(c *gin.Context) {
	idStr, found := c.Params.Get("id")
	id, err := strconv.Atoi(idStr)
	if !found || err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "id not found"})
	}
	blog, err := sql.GetById(id)

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
		Bt      int64  `form:"begin_time"  binding:"required"`
		Et      int64  `form:"end_time"   binding:"required"`
		Content string `form:"content" binding:"required"`
		Tags    string `form:"tags" binding:"required"`
		Type    int    `form:"type" binding:"required"`
		Status  int    `form:"status" binding:"required"`
		Limit   int    `form:"limit" binding:"required"`
		Offset  int    `form:"offset" binding:"required"`
	}

	// 将请求参数绑定到结构体中
	if err := c.BindJSON(&req); err != nil {
		c.JSON(502, gin.H{"error": err.Error()})
		return
	}

	// title string, bt int64, et int64, content string, tag string, _type string, status int, limit, offset int
	blogs, err := sql.Find(req.Title, req.Bt, req.Et, req.Content, req.Tags, req.Type, req.Status, req.Limit, req.Offset)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": blogs})
	}
}

func Page(c *gin.Context) {
	var req struct {
		Page int `form:"page" binding:"required"`
		Size int `form:"size"  binding:"required"`
	}

	// 将请求参数绑定到结构体中
	if err := c.BindJSON(&req); err != nil {
		c.JSON(502, gin.H{"error": err.Error()})
		return
	}

	blogs, err := sql.ListPage(req.Page, req.Size)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": blogs})
	}
}

func Add(c *gin.Context) {
	var req struct {
		Title   string `json:"title" binding:"required"`
		Pic     string `json:"pic"  binding:"required"`
		Content string `json:"content"  binding:"required"`
		Type    string `json:"type"  binding:"required"`
	}

	// 将请求参数绑定到结构体中
	if err := c.BindJSON(&req); err != nil {
		c.JSON(502, gin.H{"error": err.Error()})
		return
	}

	id, err := sql.Add(req.Title, req.Pic, req.Content, req.Type)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": id})
	}
}
