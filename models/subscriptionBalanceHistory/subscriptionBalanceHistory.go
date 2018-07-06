package subscriptionBalanceHistory

import (
	"gopkg.in/mgo.v2/bson"
	"orbits-master-api/models/auditTrail"
)

type SubscriptionBalanceHistory struct {
	Id                    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	SOF_Number            string        `bson:"SOF_number" json:"SOF_number"`
	Payment_means_number  string        `bson:"payment_means_number" json:"payment_means_number"`
	Balance               float64       `bson:"balance" json:"balance"`
	Ext_update_on         string        `bson:"ext_update_on" json:"ext_update_on"`
	auditTrail.AuditTrail `bson:"audit_trail" binding:"dive"`
}

type SubscriptionBalanceHistoryId struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
}

type SubscriptionBalanceHistoryPaymentMeansNumber struct {
	Payment_means_number string `json:"payment_means_number" bson:"payment_means_number"`
}
