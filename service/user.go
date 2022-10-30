package service

import (
	"mall/models"
	"mall/tool"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// Login
// @Tags public_api
// @Router /login [post]
// @Summary 登录
// @Description
// @Param username formData string true "账号 varchar(31)"
// @Param passwd formData string true "密码 varchar(127)"
// @Success 200 {object} define.Res_login_success "grade可用于展示，也可用于隐藏无权限功能<br><br>失败则返回 {"code": -1,"msg": "$reason"}"
func Login(c *gin.Context) {
	
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请求格式错误",
		})
		c.Abort()
        return
    }

	if user.Name == "" || user.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "账号或密码为空",
		})
		return
	}
	data := new(models.User)
	err := models.DB.Where("name = ? AND passwd = ?", user.Name, user.Password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "错误的账号或密码",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "login system error",
		})
		return
	}
	token, err := tool.GenToken(data.Name, data.Grade)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "login system error: gentoken",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": map[string]interface{}{
			"token": token,
			"grade": data.Grade,
		},
	})
}

// @Tags root_api
// @Router /adduser [post]
// @Summary [3]增加用户
// @Description add user
// @Param token header string true "token"
// @Param name formData string true "账号 varchar(31)"
// @Param passwd formData string true "密码 varchar(127)"
// @Param grade formData int true "账号权限等级"
// @Success 200 {object} define.Res_success "若失败，"code": -1,"msg": 失败原因"
func AddUser(c *gin.Context) {
	needGrade := 3
	grade, err := getGrade(c.GetHeader("token"), c)
	if err != nil {
		return
	}
	if grade < needGrade {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "无权限",
		})
		c.Abort()
		return
	}

	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请求格式错误",
		})
		c.Abort()
        return
    }

	if user.Name=="" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "账号不能为空",
		})
		c.Abort()
		return
	}
	if user.Password=="" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "密码不能为空",
		})
		c.Abort()
		return
	}
	if user.Grade==0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "权限等级不能为空",
		})
		c.Abort()
		return
	}

	if user.Grade > 3 || user.Grade < 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "未知权限等级",
		})
		c.Abort()
		return
	}

	err = models.DB.Create(user).Error
	if err != nil {
		mysqlErrhandle(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
	})
}

func getGrade(tokenStr string, c *gin.Context) (int, error) {
	userClaim, err := tool.ParseToken(tokenStr)
	if err != nil {
		if err == jwt.ErrTokenExpired {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "登录凭证已过期，请重新登录",
			})
			c.Abort()
			return 0, err
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "登录凭证已被篡改，请重新登录",
		})
		c.Abort()
		return 0,err
	}

	return userClaim.Grade, nil
}
