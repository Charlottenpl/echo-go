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
