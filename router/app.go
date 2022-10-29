package router

import (
	"github.com/gin-contrib/cors"
	"mall/middlewares"
	"mall/service"

	_ "mall/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	// cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))

	r.GET("/ping", service.GetPing)
	r.POST("/login", service.Login)

	r.POST("/getmalllist", middlewares.AuthAdminCheck(), service.GetMallList)
	r.POST("/getaptlist", middlewares.AuthAdminCheck(), service.GetAptList)
	r.POST("/getstafflist", middlewares.AuthAdminCheck(), service.GetStaffList)

	r.POST("/addmall", middlewares.AuthAdminCheck(), service.AddMall)
	r.POST("/addapt", middlewares.AuthAdminCheck(), service.AddApt)
	r.POST("/addstaff", middlewares.AuthAdminCheck(), service.AddStaff)
	r.POST("/adduser", middlewares.AuthAdminCheck(), service.AddUser)
	r.POST("/deletemall", middlewares.AuthAdminCheck(), service.DeleteMall)
	r.POST("/deleteapt", middlewares.AuthAdminCheck(), service.DeleteMall)
	r.POST("/deletestaff", middlewares.AuthAdminCheck(), service.DeleteMall)
	r.POST("/editmall", middlewares.AuthAdminCheck(), service.EditMall)
	r.POST("/editapt", middlewares.AuthAdminCheck(), service.EditApt)
	r.POST("/editstaff", middlewares.AuthAdminCheck(), service.EditStaff)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
