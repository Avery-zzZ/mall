package service

import (
	"errors"
	"mall/define"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func checkPageNSize (c *gin.Context) (page, size int, err error){
	page, err = strconv.Atoi(c.DefaultPostForm("page",define.DefaultPage)) 
	if err!=nil {
		c.JSON(http.StatusOK, gin.H{
			"code":-1,
			"msg":"请填入整数页号",
		})
		return 
	}
	size, err = strconv.Atoi(c.DefaultPostForm("size",define.DefaultSize))
	if err!=nil {
		c.JSON(http.StatusOK, gin.H{
			"code":-1,
			"msg":"请填入整数页容量",
		})
		return
	}
	return
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