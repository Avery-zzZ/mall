package service

import (

	"mall/models"
	"net/http"
	"strconv"

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

	mall_id, has := c.GetPostForm("mall_id")
	if !has {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "编码不能为空",
		})
		c.Abort()
		return
	}

	mall_name, _ := c.GetPostForm("mall_name")
	mall_address, _ := c.GetPostForm("mall_address")
	mall_tel, _ := c.GetPostForm("mall_tel")

	data := &models.Mall{
		Mall_id:      mall_id,
		Mall_name:    mall_name,
		Mall_address: mall_address,
		Mall_tel:     mall_tel,
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

	father_id_str, has := c.GetPostForm("father_id")
	if !has {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "所属商场id不能为空",
		})
		c.Abort()
		return
	}
	apt_id, has := c.GetPostForm("apt_id")
	if !has {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "编码不能为空",
		})
		c.Abort()
		return
	}

	father_id, err := strconv.Atoi(father_id_str)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "所属商场id格式错误",
		})
		c.Abort()
		return
	}

	apt_name, _ := c.GetPostForm("apt_name")
	apt_address, _ := c.GetPostForm("apt_address")
	apt_tel, _ := c.GetPostForm("apt_tel")

	data := &models.Apartment{
		Father_id:   father_id,
		Apt_id:      apt_id,
		Apt_name:    apt_name,
		Apt_address: apt_address,
		Apt_tel:     apt_tel,
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

	father_id_str, has := c.GetPostForm("father_id")
	if !has {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "所属部门id不能为空",
		})
		c.Abort()
		return
	}
	staff_id, has := c.GetPostForm("staff_id")
	if !has {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "编码不能为空",
		})
		c.Abort()
		return
	}

	father_id, err := strconv.Atoi(father_id_str)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "所属部门id格式错误",
		})
		c.Abort()
		return
	}

	staff_sal, has := c.GetPostForm("staff_sal")
	var staff_sal_value float64 = 0
	if has {
		staff_sal_value, err = strconv.ParseFloat(staff_sal, 64)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "薪水格式错误",
			})
			c.Abort()
			return
		}
	}

	staff_name, _ := c.GetPostForm("staff_name")
	staff_pos, _ := c.GetPostForm("staff_pos")
	staff_tel, _ := c.GetPostForm("staff_tel")

	data := &models.Staff{
		Father_id:  father_id,
		Staff_id:   staff_id,
		Staff_name: staff_name,
		Staff_pos:  staff_pos,
		Staff_tel:  staff_tel,
		Staff_sal:  staff_sal_value,
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


