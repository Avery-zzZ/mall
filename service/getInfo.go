package service

import (
	"log"
	"mall/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getStruct struct {
	Page      int    `json:"page" form:"page"`
	Size      int    `json:"size" form:"size"`
	Keyword   string `json:"keyword" form:"keyword"`
	Father_id int    `json:"father_id" form:"father_id"`
}

// PingExample godoc
// @Tags public_api
// @Router /ping [get]
// @Summary ping
// @Description do ping
// @Success 200 {string} json
func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// PingExample godoc
// @Tags member_api
// @Router /getmalllist [post]
// @Summary [1]获取所有商场信息
// @Param token header string true "token"
// @Param page formData int false "页号(默认1)"
// @Param size formData int false "页大小(默认20)"
// @Param keyword formData string false "搜索关键字(针对编码和名称的%*%)"
// @Success 200 {object} define.Res_get_success{data=define.Res_get_success_data{list=[]models.MallBasic}} "失败则返回 {"code": -1,"msg": "$reason"}"
func GetMallList(c *gin.Context) {

	var getStruct getStruct
	if err := c.ShouldBind(&getStruct); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请求格式错误",
		})
		c.Abort()
		return
	}

	page, size := checkPageNSize(getStruct.Page, getStruct.Size)

	page = (page - 1) * size
	var count int64

	tx := models.GetMallList(getStruct.Keyword)
	list := make([]*models.MallBasic, 0)
	err := tx.Count(&count).Omit("created_at", "updated_at", "deleted_at").Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("GetMallList error:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}

// PingExample godoc
// @Tags member_api
// @Router /getaptlist [post]
// @Summary [1]获取部门信息
// @Description father_id为空则获取所有
// @Param token header string true "token"
// @Param father_id formData int false "father_id"
// @Param page formData int false "页号(默认1)"
// @Param size formData int false "页大小(默认20)"
// @Param keyword formData string false "搜索关键字(针对编码和名称的%*%)"
// @Success 200 {object} define.Res_get_success{data=define.Res_get_success_data{list=[]models.ApartmentBasic}} "失败则返回 {"code": -1,"msg": "$reason"}"
func GetAptList(c *gin.Context) {

	var getStruct getStruct
	if err := c.ShouldBind(&getStruct); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请求格式错误",
		})
		c.Abort()
		return
	}

	page, size := checkPageNSize(getStruct.Page, getStruct.Size)

	page = (page - 1) * size
	var count int64

	tx := models.GetAptList(getStruct.Keyword)
	if getStruct.Father_id != 0 {
		tx = tx.Where("father_id = ?", getStruct.Father_id)
	}
	list := make([]*models.ApartmentBasic, 0)
	err := tx.Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("GetAptList error:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": map[string]interface{}{
			"count": count,
			"list":  list,
		},
	})
}

// PingExample godoc
// @Tags member_api
// @Router /getstafflist [post]
// @Summary [1]获取员工信息
// @Description father_id为空则获取所有
// @Param token header string true "token"
// @Param father_id formData int false "father_id"
// @Param page formData int false "页号(默认1)"
// @Param size formData int false "页大小(默认20)"
// @Param keyword formData string false "搜索关键字(针对编码和名称的%*%)"
// @Success 200 {object} define.Res_get_success{data=define.Res_get_success_data{list=[]models.StaffBasic}} "失败则返回 {"code": -1,"msg": "$reason"}"
func GetStaffList(c *gin.Context) {

	var getStruct getStruct
	if err := c.ShouldBind(&getStruct); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请求格式错误",
		})
		c.Abort()
		return
	}

	page, size := checkPageNSize(getStruct.Page, getStruct.Size)

	page = (page - 1) * size
	var count int64

	tx := models.GetStaffList(getStruct.Keyword)
	if getStruct.Father_id != 0 {
		tx = tx.Where("father_id = ?", getStruct.Father_id)
	}
	list := make([]*models.ApartmentBasic, 0)
	err := tx.Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("GetAptList error:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": map[string]interface{}{
			"count": count,
			"list":  list,
		},
	})
}
