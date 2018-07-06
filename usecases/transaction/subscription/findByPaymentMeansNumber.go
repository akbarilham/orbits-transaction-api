package subscription

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"orbits-master-api/models/subscription"
)

func FindByPaymentMeansNumber(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB("orbits-mdm").C("subscription")

	param := subscription.SubscriptionPaymentMeansNumber{}

	err := c.Bind(&param)

	if err == nil {

		result := subscription.Subscription{}

		err = con.Find(bson.M{"payment_means_number": param.Payment_means_number}).One(&result)

		if err == nil {

			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "body": result})

		} else {

			c.JSON(200, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		}

	} else {

		c.JSON(200, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}

}
