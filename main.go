package main

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	//"go-mini-arch/middleware/authentication"
	"io"
	"log"
	_ "net/http"
	"./driver/databases"
	
	"./usecases/transaction"
	//"orbits-transaction-api/usecases/transaction/customer"
	//"orbits-transaction-api/usecases/transaction/inventory"
	//"orbits-transaction-api/usecases/transaction/subscription"
	"os"
)

type Config struct {
	PORT string
}

func main() {

	/* Create connection engine */
	database.ConnectMgo()
	defer database.MgoSession.Close()

	/* Getting app configuration */
	var config Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Println(err)
	}

	/* Welcome to gin framework ! */

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	/* Create gin route */
	//r := gin.Default()
	r := gin.Default()

	/* Inject mongo session to Gin Context*/
	r.Use(mapMongo)
	/* Bind route to middleware (in this case authentication)*/
	// r.Use(authentication.Authentication)

	Transaction := r.Group("/transaction")
	{
		List := Transaction.Group("/list")
		{
			// List.POST("/getbyeventinput", transaction.GetByEventInput)
			// List.POST("/getbyeventrated", transaction.GetByEventRated)
			List.GET("/getbyall", transaction.GetByTransactionList)
			// List.POST("/gettransactionhistorybydate", transaction.GetTransactionHistoryByDate)
			// List.POST("/gettransactionrecapbydate", transaction.GetTransactionRecapByDate)
		}

	}
	/* Register some route (ping) and bind it to certain usecase(controller) */
	/* Listen and serve to a port */
	r.Run(config.PORT)
}

/* put mongo db session to gin context */
func mapMongo(c *gin.Context) {
	session := database.MgoSession.Clone()

	defer session.Close()

	c.Set("mongoSession", session)
	c.Next()
}
