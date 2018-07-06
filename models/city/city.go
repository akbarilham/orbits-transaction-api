package city

import (
	"gopkg.in/mgo.v2/bson"
	"orbits-master-api/models/auditTrail"
)

type City struct {
	Id                    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Province_code         string        `json:"province_code" bson:"province_code`
	Code                  string        `json:"code" bson:"code,omitempty`
	ISO_code              string        `json:"iso_code" bson:"iso_code,omitempty`
	Name                  string        `json:"name" bson:"name,omitempty`
	Description           string        `json:"description" bson:"description,omitempty`
	Status                string        `json:"status" bson:"status`
	auditTrail.AuditTrail `bson:"audit_trail" binding:"dive"`
}

type CityId struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
}

type CityCode struct {
	Code string `json:"code" bson:"code"`
}
type CityProvinceCode struct {
	Province_code string `json:"province_code" bson:"province_code"`
}
