package service

import (
	"mall/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags admin_api
// @Router /editmall [post]
// @Summary [2]修改商场信息
// @Description 进入修改界面后，将原来的值填入对应输入框
// @Param token header string true "token"
// @Param ID formData int true "商场ID"
// @Param mall_id formData string true "商场编码 varchar(15)"
// @Param mall_name formData string false "商场名称 varchar(63)"
// @Param mall_address formData string false "商场地址 varchar(255)"
// @Param mall_tel formData string false "商场电话 varchar(255)"
// @Success 200 {object} define.Res_success "失败则返回 {"code": -1,"msg": "$reason"}"
func EditMall(c *gin.Context) {
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


	if mall.Mall_id=="" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "编码不能为空",
		})
		c.Abort()
		return
	}
	if mall.ID==0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "ID不能为空",
		})
		c.Abort()
		return
	}

	err = models.DB.Model(new(models.Mall)).Where("id = ?", mall.ID).Updates(mall).Error
	if err != nil {
		mysqlErrhandle(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
	})
}

// @Tags admin_api
// @Router /editapt [post]
// @Summary [2]修改部门信息
// @Description 进入修改界面后，将原来的值填入对应输入框
// @Param token header string true "token"
// @Param ID formData int true "部门ID"
// @Param apt_id formData string true "部门编码 varchar(15)"
// @Param apt_name formData string false "部门名称 varchar(63)"
// @Param apt_address formData string false "部门地址 varchar(255)"
// @Param apt_tel formData string false "部门电话 varchar(255)"
// @Success 200 {object} define.Res_success "失败则返回 {"code": -1,"msg": "$reason"}"
func EditApt(c *gin.Context) {
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

	var apt models.Apartment
	if err = c.ShouldBind(&apt) ; err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请求格式错误",
		})
		c.Abort()
		return
	}

	if apt.Apt_id=="" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "编码不能为空",
		})
		c.Abort()
		return
	}
	if apt.ID==0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "ID不能为空",
		})
		c.Abort()
		return
	}

	err = models.DB.Model(new(models.Apartment)).Where("id = ?", apt.ID).Updates(apt).Error
	if err != nil {
		mysqlErrhandle(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
	})
}

// @Tags admin_api
// @Router /editstaff [post]
// @Summary [2]修改员工信息
// @Description 进入修改界面后，将原来的值填入对应输入框
// @Param token header string true "token"
// @Param ID formData int true "员工ID"
// @Param staff_id formData string true "员工编码 varchar(15)"
// @Param staff_name formData string false "员工姓名 varchar(63)"
// @Param staff_pos formData string false "员工岗位 varchar(255)"
// @Param staff_tel formData string false "员工电话 varchar(255)"
// @Param staff_sal formData float64 false "员工薪水"
// @Success 200 {object} define.Res_success "失败则返回 {"code": -1,"msg": "$reason"}"
func EditStaff(c *gin.Context) {
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

	var staff models.Staff
	if err = c.ShouldBind(&staff) ; err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请求格式错误",
		})
		c.Abort()
		return
	}

	if staff.ID==0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "ID不能为空",
		})
		c.Abort()
		return
	}
	if staff.Staff_id=="" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "编码不能为空",
		})
		c.Abort()
		return
	}
	if staff.Staff_sal<0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "薪水不能为负数",
		})
		c.Abort()
		return
	}

	err = models.DB.Model(new(models.Staff)).Where("id = ?", staff.ID).Updates(staff).Error
	if err != nil {
		mysqlErrhandle(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
	})
}
