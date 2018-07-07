package transactionList

import (
	"orbits-transaction-api/models/auditTrail"

	"gopkg.in/mgo.v2/bson"
)

type Tariff struct {
	Id                    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Type                  string        `bson:"type" json:"type"`
	Origin                string        `bson:"origin" json:"origin, omitempty"`
	Destination           string        `bson:"destination" json:"destination"`
	Classification        string        `bson:"classification" json:"classification"`
	Value                 string        `bson:"value" json:"value"`
	Status                string        `bson:"status" json:"status"`
	auditTrail.AuditTrail `bson:"audit_trail" binding:"dive"`
}

type TariffId struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
}
