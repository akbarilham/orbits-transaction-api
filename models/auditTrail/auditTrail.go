package auditTrail

import (
	"time"
)

type AuditTrail struct {
	Created_on   time.Time `bson:"created_on" json:"created_on,omitempty"`
	Created_by   string    `bson:"created_by" json:"created_by,omitempty"`
	Modified_on  time.Time `bson:"modified_on" json:"modified_on"`
	Modified_by  string    `bson:"modified_by" json:"modified_by"`
	Company_Id   string    `bson:"company_id" json:"company_id,omitempty"`
	Company_code string    `bson:"company_code" json:"company_code,omitempty"`
	Company_Name string    `bson:"company_name" json:"company_name,omitempty" `
}
