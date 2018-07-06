package subscription

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"orbits-master-api/models/subscription"
	"time"
)

func Insert(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB("orbits-mdm").C("subscription")

	subscription := subscription.Subscription{}

	err := c.Bind(&subscription)

	if err == nil {

		subscription.Id = bson.NewObjectId() //initialize object ID
		subscription.AuditTrail.Created_on = time.Now()
		subscription.AuditTrail.Created_by = c.ClientIP()

		err = con.Insert(&subscription)

		if err == nil {

			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Subscriber has been successfully inserted", "id": subscription.Id})

		} else {

			c.JSON(200, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		}

	} else {

		c.JSON(200, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}

}
