package service

import (

	"mall/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Tags admin_api
// @Router /addmall [post]
// @Summary 添加商场
// @Description add mall
// @Param token header string true "token"
// @Param mall_id formData string true "商场编码（不可重复）"
// @Param mall_name formData string false "商场名称"
// @Param mall_address formData string false "商场地址"
// @Param mall_tel formData string false "商场电话"
// @Success 200 {string} json "add mall"
func AddMall(c *gin.Context) {
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

	var mall models.MallBasic
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

	err = models.DB.Create(&models.Mall{MallBasic: mall}).Error
	if err != nil {
		mysqlErrhandle(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
	})
}

// PingExample godoc
// @Tags admin_api
// @Router /addapt [post]
// @Summary 添加部门
// @Description add apt
// @Param token header string true "token"
// @Param father_id formData int true "所属商场id"
// @Param apt_id formData string true "部门编码（不可重复）"
// @Param apt_name formData string false "部门名称"
// @Param apt_address formData string false "部门地址"
// @Param apt_tel formData string false "部门电话"
// @Success 200 {string} json "add apt"
func AddApt(c *gin.Context) {
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

	var apt models.ApartmentBacic
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

	err = models.DB.Create(&models.Apartment{ApartmentBacic: apt}).Error
	if err != nil {
		mysqlErrhandle(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
	})
}



// PingExample godoc
// @Tags admin_api
// @Router /addstaff [post]
// @Summary 添加员工
// @Description add staff
// @Param token header string true "token"
// @Param father_id formData int true "所属部门id"
// @Param staff_id formData string true "员工编码（不可重复）"
// @Param staff_name formData string false "员工姓名"
// @Param staff_pos formData string false "员工岗位"
// @Param staff_tel formData string false "员工电话"
// @Param staff_sal formData float64 false "员工薪水"
// @Success 200 {object} define.Res_success "若失败，"code": -1,"msg": 失败原因"
func AddStaff(c *gin.Context) {
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

	var staff models.StaffBasic
	if err = c.ShouldBind(&staff) ; err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请求格式错误",
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

	err = models.DB.Create(&models.Staff{StaffBasic: staff}).Error
	if err != nil {
		mysqlErrhandle(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
	})
}


