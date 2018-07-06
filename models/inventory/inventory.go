package inventory

import (
	"gopkg.in/mgo.v2/bson"
	"orbits-master-api/models/auditTrail"
)

type Inventory struct {
	Id                        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Inventory_type_code       string        `bson:"inventory_type_code" json:"inventory_type_code"`
	Company_code              string        `bson:"company_code" json:"company_code, omitempty"`
	Payment_means_number      string        `bson:"payment_means_number" json:"payment_means_number"`
	Manufacture_serial_number string        `bson:"manufacture_serial_number" json:"manufacture_serial_number"`
	Distribution_level        string        `bson:"distribution_level" json:"distribution_level"`
	Description               string        `bson:"description" json:"description"`
	Storage_id                string        `bson:"storage_id" json:"storage_id"`
	Status                    string        `bson:"status" json:"status"`
	auditTrail.AuditTrail     `bson:"audit_trail" binding:"dive"`
}

type InventoryPairByPaymentMeansNumber struct {
	Payment_means_number  string `bson:"payment_means_number" json:"payment_means_number"`
	Status                string `bson:"status" json:"status"`
	auditTrail.AuditTrail `bson:"audit_trail" binding:"dive"`
}

type InventoryId struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
}

type InventoryPaymentMeansNumber struct {
	Payment_means_number string `json:"payment_means_number" bson:"payment_means_number"`
}

type InventoryCompanyCode struct {
	Company_code string `json:"company_code" bson:"company_code"`
}
