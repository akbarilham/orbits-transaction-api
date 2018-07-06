package customer

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"orbits-master-api/models/customer"
	"time"
)

func Insert(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB("orbits-mdm").C("customer")

	customer := customer.Customer{}

	err := c.Bind(&customer)

	if err == nil {

		customer.Id = bson.NewObjectId() //initialize object ID
		customer.AuditTrail.Created_on = time.Now()
		customer.AuditTrail.Created_by = c.ClientIP()

		/* Making sure no certain field Duplication ! */
		// index := mgo.Index{
		// 	Key:        []string{"identity_card_id", "email_address"},
		// 	Unique:     true,
		// 	DropDups:   true,
		// 	Background: true, // See notes.
		// 	Sparse:     true,
		// }

		// err = con.EnsureIndex(index)

		// if err == nil {

		// } else {
		// 	c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		// }

		err = con.Insert(&customer)

		if err == nil {

			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Customer has been successfully inserted", "id": customer.Id})

		} else {

			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		}

	} else {

		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}

}
