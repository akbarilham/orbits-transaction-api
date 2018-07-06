package subscriptionBalance

import (
	"gopkg.in/mgo.v2/bson"
	"orbits-master-api/models/auditTrail"
)

type SubscriptionBalance struct {
	Id                    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Subscription_id       bson.ObjectId `json:"subscription_id" bson:"subscription_id,omitempty"`
	SOF_number            string        `bson:"SOF_number" json:"SOF_number"`
	Payment_means_number  string        `bson:"payment_means_number" json:"payment_means_number"`
	Balance               float64       `bson:"balance" json:"balance"`
	Last_balance          float64       `bson:"last_balance" json:"last_balance"`
	Last_balance_date     string        `bson:"last_balance_date" json:"last_balance_date"`
	auditTrail.AuditTrail `bson:"audit_trail" binding:"dive"`
}

type SubscriptionBalanceId struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
}

type SubscriptionBalanceSOFNumber struct {
	SOF_number string `json:"SOF_Number" bson:"SOF_Number"`
}

type SubscriptionBalancePaymentMeansNumber struct {
	Payment_means_number string `json:"payment_means_number" bson:"payment_means_number"`
}
