package customer

import (
	"gopkg.in/mgo.v2/bson"
	"orbits-master-api/models/auditTrail"
)

type Customer struct {
	Id                    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Full_name             string        `bson:"full_name" json:"full_name" binding:"required"`
	Username              string        `bson:"username" json:"username, omitempty"`
	Email_address         string        `bson:"email_address" json:"email_address" binding:"required"`
	Date_of_birth         string        `bson:"date_of_birth" json:"date_of_birth" binding:"required"`
	Gender                string        `bson:"gender" json:"gender"`
	Identity_card_id      string        `bson:"identity_card_id" json:"identity_card_id" binding:"required"`
	Phone_number          string        `bson:"phone_number" json:"phone_number"`
	Address_1             string        `bson:"address_1" json:"address_1"`
	Address_2             string        `bson:"address_2" json:"address_2"`
	Address_3             string        `bson:"address_3" json:"address_3"`
	Province_code         string        `bson:"province_code" json:"province_code"`
	City_code             string        `bson:"city_code" json:"city_code"`
	Zip_code              string        `bson:"zip_code" json:"zip_code"`
	Status_flag           string        `bson:"status_flag" json:"status_flag"`
	Verification_flag     string        `bson:"verification_flag" json:"verification_flag"`
	auditTrail.AuditTrail `bson:"audit_trail" binding:"dive"`
}

type CustomerVerifyById struct {
	Id                    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Verification_flag     string        `bson:"verification_flag" json:"verification_flag"`
	auditTrail.AuditTrail `bson:"audit_trail" binding:"dive"`
}

type CustomerId struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
}

type CustomerEmail struct {
	Email_address string `json:"email, required" bson:"email_address"`
}

type CustomerUsername struct {
	Username string `json:"username" bson:"username"`
}

type CustomerIdentityCardId struct {
	Identity_card_id string `json:"identity_card_id" bson:"identity_card_id"`
}
