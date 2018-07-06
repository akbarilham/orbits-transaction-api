package subscription

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"orbits-master-api/models/subscription"
)

func FindByEmail(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB("orbits-mdm").C("subscription")

	email := subscription.SubscriptionEmail{}

	err := c.Bind(&email)

	if err == nil {

		result := subscription.Subscription{}

		err = con.Find(bson.M{"email_address": email.Email_address}).One(&result)

		if err == nil {

			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "body": result})

		} else {

			c.JSON(200, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		}

	} else {

		c.JSON(200, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}

}