package service

import (
	"mall/models"
	"mall/tool"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// Login
// @Tags public_api
// @Router /login [post]
// @Summary 登录
// @Description
// @Param username formData string true "account"
// @Param passwd formData string true "password"
// @Success 200 {object} define.Res_login_success "若失败，"code": -1,"msg": 失败原因"
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

// PingExample godoc
// @Tags root_api
// @Router /adduser [post]
// @Summary 增加用户
// @Description add user
// @Param token header string true "token"
// @Param name formData string true "账号"
// @Param passwd formData string false "密码"
// @Param grade formData int false "账号权限等级"
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

	name, has := c.GetPostForm("name")
	if !has {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "账号不能为空",
		})
		c.Abort()
		return
	}
	passwd, has := c.GetPostForm("passwd")
	if !has {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "密码不能为空",
		})
		c.Abort()
		return
	}
	grade_str, has := c.GetPostForm("grade")
	if !has {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "权限等级不能为空",
		})
		c.Abort()
		return
	}
	grade_, err := strconv.Atoi(grade_str)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "权限等级格式错误",
		})
		c.Abort()
		return
	}
	if grade_ > 3 || grade_ < 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "未知权限等级",
		})
		c.Abort()
		return
	}

	data := &models.User{
		Name:     name,
		Password: passwd,
		Grade:    grade_,
	}

	err = models.DB.Create(data).Error
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
