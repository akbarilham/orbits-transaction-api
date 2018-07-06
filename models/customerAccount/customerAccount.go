package customerAccount

import (
	"gopkg.in/mgo.v2/bson"
	"orbits-master-api/models/auditTrail"
)

type CustomerAccount struct {
	Id                    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Customer_id           bson.ObjectId `json:"customer_id" bson:"customer_id" binding:"required"`
	Full_name             string        `bson:"full_name" json:"full_name" binding:"required"`
	Email_address         string        `bson:"email_address" json:"email_address"`
	Date_of_birth         string        `bson:"date_of_birth" json:"date_of_birth"`
	Gender                string        `bson:"gender" json:"gender"`
	Identity_card_id      string        `bson:"identity_card_id" json:"identity_card_id"`
	Phone_number          string        `bson:"phone_number" json:"phone_number" binding:"required"`
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

type CustomerAccountId struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
}

type CustomerAccountEmail struct {
	Email_address string `json:"email, required" bson:"email_address"`
}

type CustomerAccountCustomerId struct {
	Customer_id bson.ObjectId `json:"customer_id" bson:"customer_id"`
}

type CustomerAccountIdentityCardId struct {
	Identity_card_id string `json:"identity_card_id" bson:"identity_card_id"`
}
