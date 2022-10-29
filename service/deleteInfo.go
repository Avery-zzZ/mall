package service

import (
	"mall/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags admin_api
// @Router /deletemall [post]
// @Summary [2]删除商场（连锁删除）
// @Description 需确认操作：此操作会将所有所属部门和员工删除
// @Param token header string true "token"
// @Param ID formData int true "商场ID"
// @Success 200 {object} define.Res_success "失败则返回 {"code": -1,"msg": "$reason"}"
func DeleteMall(c *gin.Context) {
	needGrade := 2
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

	var mall models.Mall
	if err = c.ShouldBind(&mall) ; err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请求格式错误",
		})
		c.Abort()
		return
	}

	if mall.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "ID不能为空",
		})
		c.Abort()
		return
	}

	err = models.DB.Model(new(models.Mall)).Where("id = ?", mall.ID).Unscoped().Delete(&models.Mall{}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "删除成功",
	})
}

// @Tags admin_api
// @Router /deleteapt [post]
// @Summary [2]删除部门（连锁删除）
// @Description 需确认操作：此操作会将所有所属员工删除
// @Param token header string true "token"
// @Param ID formData int true "部门ID"
// @Success 200 {object} define.Res_success "失败则返回 {"code": -1,"msg": "$reason"}"
func DeleteApartment(c *gin.Context) {
	needGrade := 2
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

	ID, has := c.GetPostForm("ID")
	if !has {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "ID不能为空",
		})
		c.Abort()
		return
	}

	err = models.DB.Model(new(models.Apartment)).Where("id = ?", ID).Unscoped().Delete(&models.Apartment{}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "删除成功",
	})
}

// @Tags admin_api
// @Router /deletestaff [post]
// @Summary [2]删除员工
// @Description 需确认操作
// @Param token header string true "token"
// @Param ID formData int true "员工ID"
// @Success 200 {object} define.Res_success "失败则返回 {"code": -1,"msg": "$reason"}"
func DeleteStaff(c *gin.Context) {
	needGrade := 2
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

	ID, has := c.GetPostForm("ID")
	if !has {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "ID不能为空",
		})
		c.Abort()
		return
	}

	err = models.DB.Model(new(models.Staff)).Where("id = ?", ID).Unscoped().Delete(&models.Staff{}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "删除成功",
	})
}
