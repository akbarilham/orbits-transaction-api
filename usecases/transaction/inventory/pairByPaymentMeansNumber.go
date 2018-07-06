package inventory

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"orbits-master-api/models/inventory"
	"time"
)

func PairByPaymentMeansNumber(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB("orbits-mdm").C("inventory")

	inventory := inventory.InventoryPairByPaymentMeansNumber{}

	err := c.Bind(&inventory)

	if err == nil {

		inventory.AuditTrail.Modified_on = time.Now()
		inventory.AuditTrail.Modified_by = c.ClientIP()

		updateParam := bson.M{
			"$set": bson.M{
				"status":                  "PAIRED",
				"audit_trail.modified_on": inventory.AuditTrail.Modified_on,
				"audit_trail.modified_by": inventory.AuditTrail.Modified_by,
			},
		}

		err = con.Update(bson.M{"payment_means_number": inventory.Payment_means_number}, updateParam)

		if err == nil {

			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Inventory has been successfully paired", "payment_means_number": inventory.Payment_means_number})

		} else {

			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		}

	} else {

		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}
}
