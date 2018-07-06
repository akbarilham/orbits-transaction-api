package province

import (
	"gopkg.in/mgo.v2/bson"
	"orbits-master-api/models/auditTrail"
)

type Province struct {
	Id                    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Country_code          string        `json:"country_code" bson:"country_code`
	Code                  string        `json:"code" bson:"code,omitempty`
	ISO_code              string        `json:"iso_code" bson:"iso_code,omitempty`
	Name                  string        `json:"name" bson:"name,omitempty`
	Description           string        `json:"description" bson:"description,omitempty`
	Status                string        `json:"status" bson:"status`
	auditTrail.AuditTrail `bson:"audit_trail" binding:"dive"`
}

type ProvinceId struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
}

type ProvinceCode struct {
	Code string `json:"code" bson:"code"`
}

type ProvinceCountryCode struct {
	Country_code string `json:"country_code" bson:"country_code"`
}
