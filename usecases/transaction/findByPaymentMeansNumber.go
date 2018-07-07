package transaction

import (
	"net/http"
	"orbits-transaction-api/models/subscriptionBalance"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func FindByPaymentMeansNumber(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB("orbits-mdm").C("subscription-balance")

	payment_means_number := subscriptionBalance.SubscriptionBalancePaymentMeansNumber{}

	err := c.Bind(&payment_means_number)

	if err == nil {

		result := subscriptionBalance.SubscriptionBalance{}

		err = con.Find(bson.M{"payment_means_number": payment_means_number.Payment_means_number}).One(&result)

		if err == nil {

			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "body": result})

		} else {

			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Failed getting subscriber by payment means number"})
		}

	} else {

		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Failed getting subscriber by payment means number"})
	}

}
