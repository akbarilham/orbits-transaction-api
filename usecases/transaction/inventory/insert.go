package inventory

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"orbits-master-api/models/inventory"
	"time"
)

func Insert(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB("orbits-mdm").C("inventory")

	inventory := inventory.Inventory{}

	err := c.Bind(&inventory)

	if err == nil {

		inventory.Id = bson.NewObjectId() //initialize object ID
		inventory.AuditTrail.Created_on = time.Now()
		inventory.AuditTrail.Created_by = c.ClientIP()

		err = con.Insert(&inventory)

		if err == nil {

			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Inventory has been successfully inserted", "id": inventory.Id})

		} else {

			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		}

	} else {

		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}

}
