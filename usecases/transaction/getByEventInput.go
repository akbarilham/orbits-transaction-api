package transaction

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type EventInput struct {
	Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Iddevice   string        `bson:"iddevice" json:"iddevice"`
	Uuid_input string        `bson:"uuid_input" json:"uuid_input"`
	//lookup join
	Event_rated       interface{} `bson:"event_rated" json:"event_rated"`
	Event_rated_index interface{} `bson:"event_rated_index" json:"event_rated_index"`
	Balance_info      BalanceInfo `bson:"balanceinfo" json:"balanceinfo"`
}

type ETime struct {
	ETime time.Time `json:"time" bson:"time"`
}

func GetByEventInput(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	collection := session.DB("orbits-transaction").C("event_input")

	result := []EventInput{}

	err := collection.Find(bson.M{}).All(&result)

	if err == nil {

		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "body": result})

	} else {

		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Failed getting transaction by event input"})
	}

}

type EventRated struct {
	Id            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	External_id   string        `bson:"externalid" json:"externalid"`
	Start_date    time.Time     `bson:"startdate" json:"startdate"`
	Transfer_date time.Time     `bson:"transferdate" json:"transferdate"`
	Balance_info  BalanceInfo   `bson:"balanceinfo" json:"balanceinfo"`
}

type BalanceInfo struct {
	Account_info       string    `json:"accountid" bson:"accountid"`
	Balance_amount     int       `json:"balanceamount" bson:"balanceamount"`
	Balance_queried_at time.Time `json:"balancequeriedat" bson:"balancequeriedat"`
}

func GetByEventRated(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	collection := session.DB("orbits-transaction").C("event_rated")

	result := []EventRated{}

	err := collection.Find(bson.M{}).All(&result)

	if err == nil {

		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "body": result})

	} else {

		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Failed getting transaction by event input"})
	}

}

// type RequestUuid struct {
// 	Uuid_input  string `bson:"uuid_input" json:"uuid_input"`
// 	External_id string `bson:"externalid" json:"externalid"`
// }

func GetByTransactionList(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	eventInputCon := session.DB("orbits-transaction").C("event_input")

	pipeline := []bson.M{
		bson.M{"$sort": bson.M{"timestamp_tr.time": 1}},
		bson.M{"$lookup": bson.M{"from": "event_rated", "localField": "uuid_input", "foreignField": "externalid", "as": "event_rated"}},
		// bson.M{"$project": bson.M{"_id": 0, "iddevice": 1, "uuid_input": 1, "event_rated": 1}},
	}

	pipe := eventInputCon.Pipe(pipeline)
	resp := []EventInput{}
	err := pipe.All(&resp)

	if err != nil {
		fmt.Println("Errored: %#v \n", err)
	}

	// dataRecap := []EventInput{}
	eventRatedCon := session.DB("orbits-transaction").C("event_rated")
	for i, data := range resp {
		fmt.Println("1 : ", data)
		os.Exit(1)
		//get subscription
		chargedPipeline := []bson.M{
			bson.M{"$match": bson.M{"externalid": ""}},
		}

		charged := EventRated{}

		pipe2 := eventRatedCon.Pipe(chargedPipeline)
		err2 := pipe2.One(&charged)

		if err2 != nil {
			fmt.Println("Errored: %#v \n", err2)
		}
		fmt.Println("2 : ", charged)
		fmt.Println("3 : ", i)
		//push plaza to plaza_name
		resp[i].Event_rated = charged.Balance_info
	}
	// var transactions []EventRated
	// err = eventRatedCon.Find(bson.M{}).All(&transactions)
	// // handle err
	// for index, transaction := range transactions {
	// 	fmt.Printf("%d: %+v\n", index, transaction.External_id)
	// }
	fmt.Println("5 : ")

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "body": resp})

}
