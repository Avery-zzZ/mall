package service

import (
	"errors"
	"mall/define"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func checkPageNSize (page, size int) (int,int){
	if page==0 {
		page = define.DefaultPage
	}
	if size==0 {
		size = define.DefaultSize
	}
	return page,size
}

func mysqlErrhandle(err error, c *gin.Context) {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		case 1062:
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "该编号已存在",
			})
		case 1406:
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "存在过长文本",
			})
		case 1452:
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "ID不存在",
			})
		case 1264:
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "存在过大数值",
			})
		default:
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "失败: 未定义的sql错误",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "失败",
		})
	}

	c.Abort()
}