package inventory

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"orbits-master-api/models/inventory"
)

func FindById(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB("orbits-mdm").C("inventory")

	id := inventory.InventoryId{}

	err := c.Bind(&id)

	if err == nil {

		result := inventory.Inventory{}

		err = con.Find(bson.M{"_id": id.Id}).One(&result)

		if err == nil {

			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "body": result})

		} else {

			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		}

	} else {

		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}

}