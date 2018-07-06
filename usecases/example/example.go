package example

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
	"orbits-master-api/driver/databases"
	"orbits-master-api/models/customer"
	"orbits-master-api/models/example" // Include necessary model
	log "orbits-master-api/util/logger"
	"time"
)

/* Example usecase */
func Example(c *gin.Context) {
	fmt.Println(" Hello, this is example process... ")

	/*  Write log using logger */
	log.WithFields(logrus.Fields{
		"animal": "wallmart",
	}).Info("A walrus appears")
}

func InsertExample(c *gin.Context) {

	customer := customer.Customer{}

	err := c.Bind(&customer)

	if err == nil {

		customer.AuditTrail.Created_on = time.Now()
		customer.AuditTrail.Created_by = c.ClientIP()

		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "body": customer})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}
}

/* Basic Example usage for database operation */
func My_DB_Select(c *gin.Context) {

	/* Create connection engine */
	database.ConnectPGEngine()
	defer database.PGEngine.Close()

	/* Initialize "slice" of user Model */
	users := []example.User{}

	/* Select all rows Statement */
	//con.Find(&users)

	/* Select with conditions Statement*/
	//con.Where("id = ?", 1).Find(&users)

	/* Select with join Statement */
	database.PGEngine.Select("address").Find(&users).Joins(" left join public.user_details on public.user.id=public.user_details.user_id ").Rows()

	/* Print *JSON response via gin response*/
	c.JSON(200, users)

}
