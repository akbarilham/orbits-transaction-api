package subscription

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"orbits-master-api/models/subscription"
	"time"
)

func UpdateById(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB("orbits-mdm").C("subscription")

	subscription := subscription.Subscription{}

	err := c.Bind(&subscription)

	if err == nil {

		subscription.AuditTrail.Modified_on = time.Now()
		subscription.AuditTrail.Modified_by = c.ClientIP()

		updateParam := bson.M{
			"$set": bson.M{
				"subscription_type_code":               subscription.Subscription_type_code,
				"full_name":                            subscription.Full_name,
				"email_address":                        subscription.Email_address,
				"SOF_number":                           subscription.SOF_number,
				"payment_means_number":                 subscription.Payment_means_number,
				"vehicle_registration_number":          subscription.Vehicle_registration_number,
				"vehicle_registration_number_file_url": subscription.Vehicle_registration_number_file_url,
				"vehicle_plate_number":                 subscription.Vehicle_plate_number,
				"vehicle_machine_number":               subscription.Vehicle_machine_number,
				"vehicle_frame_number":                 subscription.Vehicle_frame_number,
				"vehicle_brand":                        subscription.Vehicle_brand,
				"vehicle_model":                        subscription.Vehicle_model,
				"date_of_birth":                        subscription.Date_of_birth,
				"gender":                               subscription.Gender,
				"phone_number":                         subscription.Phone_number,
				"address_1":                            subscription.Address_1,
				"address_2":                            subscription.Address_2,
				"address_3":                            subscription.Address_3,
				"province_code":                        subscription.Province_code,
				"city_code":                            subscription.City_code,
				"zip_code":                             subscription.Zip_code,
				"status_flag":                          subscription.Status_flag,
				"verification_flag":                    subscription.Verification_flag,
				"audit_trail.modified_on":              subscription.AuditTrail.Modified_on,
				"audit_trail.modified_by":              subscription.AuditTrail.Modified_by,
			},
		}

		err = con.Update(bson.M{"_id": subscription.Id}, updateParam)

		if err == nil {

			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Customer has been successfully updated", "id": subscription.Id})

		} else {

			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		}

	} else {

		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}
}
