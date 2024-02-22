package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                  primitive.ObjectID `bson:"_id"`
	First_Name          *string            `json:"first_name" validate:"required,min=3"`
	Last_Name           *string            `json:"last_name" validate:"required,min=2"`
	Password            *string            `json:"password" validate:"required,min=8,max=16"`
	Email               *string            `json:"email" validate:"required,email"`
	Phone               *string            `json:"phone" validate:"required"`
	Token               *string            `json:"token"`
	User_type           *string            `json:"user_type" validate:"required,eq=USER|eq=ADMIN"`
	Refresh_token       *string            `json:"refresh_token"`
	Created_at          time.Time          `json:"created_at"`
	Updated_at          time.Time          `json:"updated_at"`
	User_id             string             `json:"user_id"`
	Study_Plan_Selected Study_Plan         `json:"study_plan_selected"`
}
