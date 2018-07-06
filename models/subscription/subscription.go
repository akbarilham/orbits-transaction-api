package subscription

import (
	"gopkg.in/mgo.v2/bson"
	"orbits-master-api/models/auditTrail"
)

type Subscription struct {
	Id                                   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Subscription_type_code               string        `json:"subscription_type_code" bson:"subscription_type_code,omitempty"`
	Customer_account_id                  bson.ObjectId `json:"customer_account_id" bson:"customer_account_id,omitempty"`
	SOF_number                           string        `bson:"SOF_number" json:"SOF_number"`
	Payment_means_number                 string        `bson:"payment_means_number" json:"payment_means_number" binding:"required"`
	Vehicle_registration_number          string        `bson:"vehicle_registration_number" json:"vehicle_registration_number"`
	Vehicle_registration_number_file_url string        `bson:"vehicle_registration_number_file_url" json:"vehicle_registration_number_file_url"`
	Vehicle_plate_number                 string        `bson:"vehicle_plate_number" json:"vehicle_plate_number"`
	Vehicle_machine_number               string        `bson:"vehicle_machine_number" json:"vehicle_machine_number"`
	Vehicle_frame_number                 string        `bson:"vehicle_frame_number" json:"vehicle_frame_number"`
	Vehicle_brand                        string        `bson:"vehicle_brand" json:"vehicle_brand"`
	Vehicle_model                        string        `bson:"vehicle_model" json:"vehicle_model"`
	Full_name                            string        `bson:"full_name" json:"full_name"`
	Email_address                        string        `bson:"email_address" json:"email_address" binding:"required"`
	Date_of_birth                        string        `bson:"date_of_birth" json:"date_of_birth"`
	Gender                               string        `bson:"gender" json:"gender"`
	Phone_number                         string        `bson:"phone_number" json:"phone_number"`
	Address_1                            string        `bson:"address_1" json:"address_1"`
	Address_2                            string        `bson:"address_2" json:"address_2"`
	Address_3                            string        `bson:"address_3" json:"address_3"`
	Province_code                        string        `bson:"province_code" json:"province_code"`
	City_code                            string        `bson:"city_code" json:"city_code"`
	Zip_code                             string        `bson:"zip_code" json:"zip_code"`
	Status_flag                          string        `bson:"status_flag" json:"status_flag"`
	Verification_flag                    string        `bson:"verification_flag" json:"verification_flag"`
	auditTrail.AuditTrail                `bson:"audit_trail" binding:"dive"`
}

type SubscriptionId struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
}

type SubscriptionEmail struct {
	Email_address string `json:"email_address" bson:"email_address"`
}

type SubscriptionSOFNumber struct {
	SOF_number string `json:"SOF_number" bson:"SOF_number"`
}

type SubscriptionCustomerAccountId struct {
	Customer_account_id bson.ObjectId `json:"customer_account_id" bson:"customer_account_id"`
}

type SubscriptionPaymentMeansNumber struct {
	Payment_means_number string `json:"payment_means_number" bson:"payment_means_number"`
}

type SubscriptionVehicleRegistrationNumber struct {
	Vehicle_registration_number string `json:"vehicle_registration_number, required" bson:"vehicle_registration_number"`
}

type SubscriptionVehiclePlateNumber struct {
	Vehicle_plate_number string `json:"vehicle_plate_number" bson:"vehicle_plate_number"`
}
