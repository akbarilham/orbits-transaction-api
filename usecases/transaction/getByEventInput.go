package transaction

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Subscriber struct {
	Payment_Means_Number string `bson:"payment_means_number" json:"payment_means_number"`
	SOF_Number           string `bson:"SOF_number" json:"SOF_number"`
	Full_Name            string `bson:"full_name" json:"full_name"`
}

type EventInput struct {
	Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Uuid_input string        `bson:"uuid_input" json:"uuid_input"`
	Trx_amount float64       `bson:"trx_amount" json:"trx_amount"`
	Pan        string        `bson:"iddevice" json:"iddevice"`
	//lookup join
	Event_rated []interface{} `bson:"event_rated" json:"event_rated"`
	Subscriber  Subscriber
	Status      string `bson:"status" json:"status'`
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

func GetByTransactionList(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	resp := []EventInput{}
	lmt, resBol := c.GetQuery("limit")
	if resBol == false {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "body": resp})
		return
	}
	if lmt == "" {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "body": resp})
		return
	}
	limit, _ := strconv.Atoi(lmt)
	pg, resBol := c.GetQuery("page")
	if resBol == false {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "body": resp})
		return
	}
	if pg == "" {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "body": resp})
		return
	}
	page, _ := strconv.Atoi(pg)

	session := c.MustGet("mongoSession").(*mgo.Session)
	eventInputCon := session.DB("orbits-transaction").C("event_input")

	subscriberCon := session.DB("orbits-mdm").C("subscription")

	var results []struct {
		Id bson.ObjectId `bson:"_id"`
	}

	var TotalPages int

	err := eventInputCon.Find(nil).Sort("_id").Select(bson.M{"_id": 1}).All(&results)

	resultsLen := len(results)

	if resultsLen == 0 {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "body": resp})
		return
	}

	if resultsLen%limit == 0 {
		TotalPages = resultsLen / limit
	} else {
		TotalPages = resultsLen/limit + 1
	}

	if page > TotalPages {
		page = TotalPages
	}

	skip := (page - 1) * limit
	max := skip + limit

	if max > resultsLen {
		max = resultsLen
	}

	if resultsLen != 1 {
		results = results[skip:max]
	}

	var ids []bson.ObjectId
	for _, v := range results {
		ids = append(ids, v.Id)
	}

	if len(ids) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "body": resp})
		return
	}

	fmt.Println("pages", page)
	if page >= 1 {
		page--
	}
	pipeline := []bson.M{
		bson.M{"$sort": bson.M{"timestamp_tr.time": 1}},
		bson.M{"$lookup": bson.M{"from": "event_rated", "localField": "uuid_input", "foreignField": "externalid", "as": "event_rated"}},

		bson.M{"$limit": page + limit},
		bson.M{"$skip": page},
		// bson.M{"$project": bson.M{"_id": 0, "iddevice": 1, "uuid_input": 1, "event_rated": 1}},
	}
	page++

	pipe := eventInputCon.Pipe(pipeline)
	err = pipe.All(&resp)

	for i := 0; i < len(resp); i++ {

		var subscription Subscriber
		err := subscriberCon.Find(bson.M{"payment_means_number": resp[i].Pan}).Sort("_id").One(&subscription)
		if err != nil {
			fmt.Println("error", err.Error(), subscription)
		}
		if len(resp[i].Event_rated) < 1 {
			resp[i].Status = "uncharged"
		} else {
			resp[i].Status = "charged"
		}
		resp[i].Subscriber = subscription
	}
	if err != nil {
		fmt.Println("Errored: %#v \n", err)
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "body": resp, "page": page, "totalPages": TotalPages})

}

func GetByTransactionListSearch(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	// resp2 := []EventInput{}

}
