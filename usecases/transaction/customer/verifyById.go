package customer

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"orbits-master-api/models/customer"
	"time"
)

func VerifyById(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB("orbits-mdm").C("customer")

	customer := customer.CustomerVerifyById{}

	err := c.Bind(&customer)

	if err == nil {

		customer.AuditTrail.Modified_on = time.Now()
		customer.AuditTrail.Modified_by = c.ClientIP()

		updateParam := bson.M{
			"$set": bson.M{
				"verification_flag":       "Y",
				"audit_trail.modified_on": customer.AuditTrail.Modified_on,
				"audit_trail.modified_by": customer.AuditTrail.Modified_by,
			},
		}

		err = con.Update(bson.M{"_id": customer.Id}, updateParam)

		if err == nil {

			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Customer has been successfully verified", "id": customer.Id})

		} else {

			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		}

	} else {

		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}
}
