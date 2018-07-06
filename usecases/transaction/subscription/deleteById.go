package subscription

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"orbits-master-api/models/subscription"
)

func DeleteById(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB("orbits-mdm").C("subscription")

	id := subscription.SubscriptionId{}

	err := c.Bind(&id)

	if err == nil {

		err = con.Remove(bson.M{"_id": id.Id})

		if err == nil {

			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Subscriber has been successfully deleted", "id": id.Id})

		} else {

			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		}

	} else {

		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}
}
