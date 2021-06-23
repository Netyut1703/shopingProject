package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Seller struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Fullname    string             `json:"fullname" bson:"fullname"`
	Address     string             `json:"address" bson:"address"`
	PhoneNumber int                `json:"phonenumber" bson:"phonenumber"`
	Email       string             `json:"email" bson:"email"`
	Createdat   time.Time          `json:"createdat" bson:"createdat"`
	Updatedat   time.Time          `json:"updatedat" bson:"updatedat"`
	Gender      string             `json:"gender" bson:"gender"`
	Password    string             `json:"password" bson:"password"`
	//ResetPassword string `json:"resetpassword" bson:"resetpassword"`
	Location    Location  `json:"location" bson:"location"`
	DateOfBirth time.Date `json:"dateofbirth" bson:"dateofbirth"`
}
