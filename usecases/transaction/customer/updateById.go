package customer

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"orbits-master-api/models/customer"
	"time"
)

func UpdateById(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB("orbits-mdm").C("customer")

	customer := customer.Customer{}

	err := c.Bind(&customer)

	if err == nil {

		customer.AuditTrail.Modified_on = time.Now()
		customer.AuditTrail.Modified_by = c.ClientIP()

		updateParam := bson.M{
			"$set": bson.M{
				"full_name":               customer.Full_name,
				"username":                customer.Username,
				"email_address":           customer.Email_address,
				"date_of_birth":           customer.Date_of_birth,
				"gender":                  customer.Gender,
				"identity_card_id":        customer.Identity_card_id,
				"phone_number":            customer.Phone_number,
				"address_1":               customer.Address_1,
				"address_2":               customer.Address_2,
				"address_3":               customer.Address_3,
				"province_code":           customer.Province_code,
				"city_code":               customer.City_code,
				"zip_code":                customer.Zip_code,
				"status_flag":             customer.Status_flag,
				"verification_flag":       customer.Verification_flag,
				"audit_trail.modified_on": customer.AuditTrail.Modified_on,
				"audit_trail.modified_by": customer.AuditTrail.Modified_by,
			},
		}

		err = con.Update(bson.M{"_id": customer.Id}, updateParam)

		if err == nil {

			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Customer has been successfully updated", "id": customer.Id})

		} else {

			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		}

	} else {

		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}
}
