package plazaLane

import (
	"gopkg.in/mgo.v2/bson"
	"orbits-master-api/models/auditTrail"
)

type PlazaLane struct {
	Id                    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Code                  string        `bson:"code" json:"code, omitempty"`
	Plaza_code            string        `bson:"plaza_code" json:"plaza_code, omitempty"`
	Sequence_number       int           `bson:"sequence_number" json:"sequence_number"`
	Description           string        `bson:"description" json:"description"`
	Status                int           `json:"status" bson:"status`
	auditTrail.AuditTrail `bson:"audit_trail" binding:"dive"`
}

type PlazaLaneId struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
}

type PlazaLaneCode struct {
	Code string `json:"code" bson:"code"`
}

type PlazaLanePlazaCode struct {
	Plaza_code string `json:"plaza_code" bson:"plaza_code"`
}
