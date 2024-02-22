package controllers

import (
	"Gate/database"
	"Gate/helpers"
	"Gate/models"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var materialCollection *mongo.Collection = database.OpenOrCreateDB(database.Client, "study_material")
var courseCollection *mongo.Collection = database.OpenOrCreateDB(database.Client, "course")
var planCollection *mongo.Collection = database.OpenOrCreateDB(database.Client, "plan")

type paginate struct {
	limit int64
	page  int64
}

func newPaginate(limit int, page int) *paginate {
	return &paginate{
		limit: int64(limit),
		page:  int64(page),
	}
}

func (m *paginate) getPaginatedOpts() *options.FindOptions {
	l := m.limit
	skip := m.page*m.limit - m.limit
	fOpt := options.FindOptions{Limit: &l, Skip: &skip}

	return &fOpt
}

func AddStudyMaterial() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var material models.Study_Material

		if err := c.BindJSON(&material); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validation := validate.Struct(material)
		if validation != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validation.Error()})
			return
		}

		num, err := materialCollection.InsertOne(ctx, material)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Material was not created"})
		}
		defer cancel()
		c.JSON(http.StatusOK, num)
	}
}

func AddCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var course models.Course

		if err := c.BindJSON(&course); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validation := validate.Struct((course))
		if validation != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validation.Error()})
			return
		}

		num, err := courseCollection.InsertOne(ctx, course)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Course was not created"})
		}
		defer cancel()
		c.JSON(http.StatusOK, num)
	}
}

func AddStudyPlan() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var plan models.Study_Plan

		if err := c.BindJSON(&plan); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validation := validate.Struct((plan))
		if validation != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validation.Error()})
			return
		}

		num, err := planCollection.InsertOne(ctx, plan)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Study Plan was not created"})
		}
		defer cancel()
		c.JSON(http.StatusOK, num)
	}
}

func GetStudyMaterials() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)

		var study_material []models.Study_Material

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil || limit < 1 {
			limit = 10
		}

		findOptions := newPaginate(limit, page).getPaginatedOpts()

		filter := bson.M{}
		if search := c.Query("search"); search != "" {
			filter = bson.M{
				"$or": []bson.M{
					{
						"material_title": bson.M{
							"$regex": primitive.Regex{
								Pattern: search,
								Options: "i",
							},
						},
					},
					{
						"tags": bson.M{
							"$regex": primitive.Regex{
								Pattern: search,
								Options: "i",
							},
						},
					},
				},
			}
		}

		if sort := c.Query("sort"); sort != "" {
			if sort == "ASC" {
				findOptions.SetSort(bson.D{{"time_duration", 1}})
			} else if sort == "DESC" {
				findOptions.SetSort(bson.D{{"time_duration", -1}})
			}
		}

		cursor, err := materialCollection.Find(ctx, filter, findOptions)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while retreiving study materials"})
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var material models.Study_Material
			cursor.Decode(&material)
			study_material = append(study_material, material)
		}

		c.JSON(http.StatusOK, study_material[0])

	}
}

func GetCourses() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)

		var courses []models.Course

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil || limit < 1 {
			limit = 10
		}

		findOptions := newPaginate(limit, page).getPaginatedOpts()

		filter := bson.M{}
		if search := c.Query("search"); search != "" {
			filter = bson.M{
				"$or": []bson.M{
					{
						"course_name": bson.M{
							"$regex": primitive.Regex{
								Pattern: search,
								Options: "i",
							},
						},
					},
				},
			}
		}

		if sort := c.Query("sort"); sort != "" {
			if sort == "ASC" {
				findOptions.SetSort(bson.D{{"total_duration", 1}})
			} else if sort == "DESC" {
				findOptions.SetSort(bson.D{{"total_duration", -1}})
			}
		}

		cursor, err := courseCollection.Find(ctx, filter, findOptions)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while retreiving courses"})
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var course models.Course
			cursor.Decode(&course)
			courses = append(courses, course)
		}

		c.JSON(http.StatusOK, courses[0])

	}
}

func GetStudyPlans() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)

		var plans []models.Study_Plan

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil || limit < 1 {
			limit = 10
		}

		findOptions := newPaginate(limit, page).getPaginatedOpts()

		filter := bson.M{}
		if search := c.Query("search"); search != "" {
			filter = bson.M{
				"$or": []bson.M{
					{
						"plan_name": bson.M{
							"$regex": primitive.Regex{
								Pattern: search,
								Options: "i",
							},
						},
					},
				},
			}
		}

		if sort := c.Query("sort"); sort != "" {
			if sort == "ASC" {
				findOptions.SetSort(bson.D{{"number_of_days", 1}})
			} else if sort == "DESC" {
				findOptions.SetSort(bson.D{{"number_of_days", -1}})
			}
		}

		cursor, err := planCollection.Find(ctx, filter, findOptions)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while retreiving study plans"})
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var plan models.Study_Plan
			cursor.Decode(&plan)
			plans = append(plans, plan)
		}

		c.JSON(http.StatusOK, plans[0])

	}
}
