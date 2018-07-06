package plaza

import (
	"gopkg.in/mgo.v2/bson"
	"orbits-master-api/models/auditTrail"
)

type Plaza struct {
	Id                    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Code                  string        `bson:"code" json:"code, omitempty"`
	Name                  string        `bson:"name" json:"name"`
	Description           string        `bson:"description" json:"description"`
	Status                int           `json:"status" bson:"status`
	auditTrail.AuditTrail `bson:"audit_trail" binding:"dive"`
}

type PlazaId struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
}

type PlazaCode struct {
	Code string `json:"code, required" bson:"code"`
}
