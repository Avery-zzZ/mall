package service

import (
	"log"
	"mall/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Tags public_api
// @Router /ping [get]
// @Summary ping
// @Description do ping
// @Success 200 {string} plain "Helloworld"
func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// PingExample godoc
// @Tags member_api
// @Router /getmalllist [post]
// @Summary 获取所有商场信息
// @Description get mall list
// @Param token header string true "token"
// @Param page formData int false "page(defualt 1)"
// @Param size formData int false "page size"
// @Param keyword formData string false "keyword"
// @Success 200 {string} json "mall list"
func GetMallList(c *gin.Context) {

	page, size, err := checkPageNSize(c)
	if err != nil {
		c.Abort()
		return
	}

	page = (page - 1) * size
	var count int64
	keyword := c.PostForm("keyword")

	tx := models.GetMallList(keyword)
	list := make([]*models.MallBasic, 0)
	err = tx.Count(&count).Omit("created_at", "updated_at", "deleted_at").Offset(page).Limit(size).Find(&list).Error
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
// @Summary 获取部门信息
// @Description get apt list
// @Param token header string true "token"
// @Param father_id formData int false "father_id"
// @Param page formData int false "page(defualt 1)"
// @Param size formData int false "page size"
// @Param keyword formData string false "keyword"
// @Success 200 {string} json "apt list"
func GetAptList(c *gin.Context) {

	fatherIdStr := c.PostForm("father_id")

	page, size, err := checkPageNSize(c)
	if err != nil {
		c.Abort()
		return
	}

	page = (page - 1) * size
	var count int64
	keyword := c.PostForm("keyword")

	tx := models.GetAptList(keyword)
	if fatherIdStr != "" {
		tx = tx.Where("father_id = ?", fatherIdStr)
	}
	list := make([]*models.ApartmentBacic, 0)
	err = tx.Count(&count).Offset(page).Limit(size).Find(&list).Error
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
// @Summary this is a summary
// @Description 获取员工信息
// @Param token header string true "token"
// @Param father_id formData int false "father_id"
// @Param page formData int false "page(defualt 1)"
// @Param size formData int false "page size"
// @Param keyword formData string false "keyword"
// @Success 200 {string} json "staff list"
func GetStaffList(c *gin.Context) {

	fatherIdStr := c.PostForm("father_id")

	page, size, err := checkPageNSize(c)
	if err != nil {
		c.Abort()
		return
	}

	page = (page - 1) * size
	var count int64
	keyword := c.PostForm("keyword")

	tx := models.GetStaffList(keyword)
	if fatherIdStr != "" {
		tx = tx.Where("father_id = ?", fatherIdStr)
	}
	list := make([]*models.StaffBasic, 0)
	err = tx.Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("GetStaffList error:", err)
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
