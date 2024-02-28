package routes

import (
	controller "Gate/controllers"
	"Gate/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	routes.Use(middleware.Authenticate())
	routes.GET("users", controller.GetUsers())
	routes.GET("users/:user_id", controller.GetUser())
	routes.POST("admin/study_material", controller.AddStudyMaterial())
	routes.POST("admin/course", controller.AddCourse())
	routes.POST("admin/study_plan", controller.AddStudyPlan())
	routes.GET("admin/study_materials", controller.GetStudyMaterials())
	routes.GET("admin/courses", controller.GetCourses())
	routes.GET("admin/study_plans", controller.GetStudyPlans())
	routes.GET("study_materials/:study_material", controller.GetStudyMaterial())
	routes.GET("courses/:course", controller.GetCourse())
	routes.GET("study_plans/:study_plan", controller.GetStudyPlan())
}
