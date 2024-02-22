package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Study_Plan struct {
	ID             primitive.ObjectID `bson:"_id"`
	Plan_Name      *string            `json:"plan_name" validate:"required,min=3"`
	Number_Of_Days int64              `json:"number_of_days" validate:"required"`
	Daily_Minutes  int64              `json:"daily_minutes"`
	Courses        []Course           `json:"course"`
}

type Course struct {
	ID               primitive.ObjectID `bson:"_id"`
	Course_Name      *string            `json:"course_name" validate:"required,min=3"`
	Total_Duration   time.Duration      `json:"total_duration"`
	Course_Materials []Study_Material   `json:"course_materials"`
}

type Study_Material struct {
	ID             primitive.ObjectID `bson:"_id"`
	Material_Title *string            `json:"material_title" validate:"required,min=3"`
	Material_Url   *string            `json:"material_url" validate:"required"`
	IsVideo        bool               `json:"isVideo"`
	Time_Duration  time.Duration      `json:"time_duration"`
	Tags           []string           `json:"tags"`
}
