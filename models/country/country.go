package country

import (
	"gopkg.in/mgo.v2/bson"
	"orbits-master-api/models/auditTrail"
)

type Country struct {
	Id                    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Code                  string        `json:"code" bson:"code,omitempty`
	ISO_Code              string        `json:"iso_code" bson:"iso_code,omitempty`
	Name                  string        `json:"name" bson:"name,omitempty`
	Description           string        `json:"description" bson:"description,omitempty`
	Status                string        `json:"status" bson:"status`
	auditTrail.AuditTrail `bson:"audit_trail" binding:"dive"`
}

type CountryId struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
}

type CountryCode struct {
	Code string `json:"code" bson:"code"`
}
