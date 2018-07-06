package cpeList

import (
	"gopkg.in/mgo.v2/bson"
	"orbits-master-api/models/auditTrail"
)

type CpeList struct {
	Id                                   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	SOF_number                           string        `bson:"SOF_number" json:"SOF_number"`
	Payment_means_number                 string        `bson:"payment_means_number" json:"payment_means_number" binding:"required"`
	Vehicle_registration_number          string        `bson:"vehicle_registration_number" json:"vehicle_registration_number"`
	Vehicle_registration_number_file_url string        `bson:"vehicle_registration_number_file_url" json:"vehicle_registration_number_file_url"`
	Vehicle_plate_number                 string        `bson:"vehicle_plate_number" json:"vehicle_plate_number"`
	Vehicle_machine_number               string        `bson:"vehicle_machine_number" json:"vehicle_machine_number"`
	Vehicle_frame_number                 string        `bson:"vehicle_frame_number" json:"vehicle_frame_number"`
	Vehicle_brand                        string        `bson:"vehicle_brand" json:"vehicle_brand"`
	Vehicle_model                        string        `bson:"vehicle_model" json:"vehicle_model"`
	Status_flag                          string        `bson:"status_flag" json:"status_flag"`
	Verification_flag                    string        `bson:"verification_flag" json:"verification_flag"`
	auditTrail.AuditTrail                `bson:"audit_trail" binding:"dive"`
}

type CpeListId struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
}

type CpeListPaymentMeansNumber struct {
	Payment_means_number string `json:"payment_means_number" bson:"payment_means_number"`
}
