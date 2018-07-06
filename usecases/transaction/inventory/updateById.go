package inventory

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"orbits-master-api/models/inventory"
	"time"
)

func UpdateById(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB("orbits-mdm").C("inventory")

	inventory := inventory.Inventory{}

	err := c.Bind(&inventory)

	if err == nil {

		inventory.AuditTrail.Modified_on = time.Now()
		inventory.AuditTrail.Modified_by = c.ClientIP()

		updateParam := bson.M{
			"$set": bson.M{
				"inventory_type_code":       inventory.Inventory_type_code,
				"company_code":              inventory.Company_code,
				"payment_means_number":      inventory.Payment_means_number,
				"manufacture_serial_number": inventory.Manufacture_serial_number,
				"distribution_level":        inventory.Distribution_level,
				"storage_id":                inventory.Storage_id,
				"description":               inventory.Description,
				"status":                    inventory.Status,
				"audit_trail.modified_on":   inventory.AuditTrail.Modified_on,
				"audit_trail.modified_by":   inventory.AuditTrail.Modified_by,
			},
		}

		err = con.Update(bson.M{"_id": inventory.Id}, updateParam)

		if err == nil {

			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Inventory has been successfully updated", "id": inventory.Id})

		} else {

			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		}

	} else {

		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}
}
