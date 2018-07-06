package transaction

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetByEventInputk(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB("orbits-transaction").C("event_input")

	pipeline := []bson.M{
		bson.M{"$match": bson.M{"_id": bson.ObjectIdHex("5ae1f04ae138232efa0a98d6")}},
		bson.M{"$sort": bson.M{"timestamp_tr.time": 1}},
		bson.M{"$lookup": bson.M{"from": "event_rated", "localField": "uuid_input", "foreignField": "externalid", "as": "event_rated"}},
		bson.M{"$unwind": bson.M{"path": "$event_rated", "preserveNullAndEmptyArrays": true, "includeArrayIndex": "event_rated_index"}},
	}

	pipe := con.Pipe(pipeline)
	resp := []bson.M{}
	// var resp = []string{"apple", "pineapple", "pie"}
	err := pipe.All(&resp)

	if err != nil {
		fmt.Println("Errored: %#v \n", err)
	}
	fmt.Println(resp)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "body": resp})

}

type RequestPayload struct {
	Start_date string `bson:"start_date" json:"start_date"`
	End_date   string `bson:"end_date" json:"end_date"`
	Plaza_code string `bson:"plaza_code" json:"plaza_code"`
}

type EventInputk struct {
	Id               bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Obu_id           string        `bson:"obu_id" json:"obu_id"`
	Trx_amount       int           `bson:"trx_amount" json:"trx_amount"`
	Plaza_code       string        `bson:"plaza_code" json:"plaza_code"`
	Processing_date  time.Time     `bson:"processing_date" json:"processing_date"`
	Lane             string        `bson:"lane" json:"lane"`
	Dst              string        `bson:"dst" json:"dst"`
	Signaling_code   string        `bson:"signaling_code" json:"signaling_code"`
	Manufactureid    string        `bson:"manufactureid" json:"manufactureid"`
	Obu_status       string        `bson:"obu_status" json:"obu_status"`
	Pricecurrency    string        `bson:"pricecurrency" json:"pricecurrency"`
	Vehicleclass     string        `bson:"vehicleclass" json:"vehicleclass"`
	Signalcodebitmap string        `bson:"signalcodebitmap" json:"signalcodebitmap"`
	Lanemode         string        `bson:"lanemode" json:"lanemode"`
	Lightsignalcode  string        `bson:"lightsignalcode" json:"lightsignalcode"`
	Iddevice         string        `bson:"iddevice" json:"iddevice"`
	Picturefilename  string        `bson:"picturefilename" json:"picturefilename"`
	Is_hit           string        `bson:"is_hit" json:"is_hit"`
	Shift            string        `bson:"shift" json:"shift"`
	Balance          int           `bson:"balance" json:"balance"`
	Uuid_input       string        `bson:"uuid_input" json:"uuid_input"`
	Status           int           `bson:"status" json:"status"`
	Event_begin_time ETime         `bson:"event_begin_time" json:"event_begin_time"`
	Timestamp_tr     ETime         `bson:"timestamp_tr" json:"timestamp_tr"`
	Shift_date       ETime         `bson:"shift_date" json:"shift_date"`
	//lookup join
	Event_rated            interface{} `bson:"event_rated" json:"event_rated"`
	Event_rated_index      interface{} `bson:"event_rated_index" json:"event_rated_index"`
	Subscription           interface{} `bson:"subscription" json:"subscription"`
	Subscription_index     interface{} `bson:"subscription_index" json:"subscription_index"`
	Customer               interface{} `bson:"customer" json:"customer"`
	Customer_index         interface{} `bson:"customer_index" json:"customer_index"`
	Customer_account       interface{} `bson:"customer_account" json:"customer_account"`
	Customer_account_index interface{} `bson:"customer_account_index" json:"customer_account_index"`
}

type ETimek struct {
	ETime time.Time `json:"time" bson:"time"`
}

func GetTransactionHistoryByDate(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB("orbits-transaction").C("event_input")

	result := []EventInputk{}

	var reqp RequestPayload
	/*Bind request payload*/
	errreq := c.Bind(&reqp)
	if errreq != nil {
		fmt.Println(errreq.Error())
	}

	/* parse date string to time object */
	//Start date
	fromDate, err := time.Parse(time.RFC3339, reqp.Start_date)

	if err != nil {
		fmt.Println(err)
	}

	//End date
	toDate, err := time.Parse(time.RFC3339, reqp.End_date)

	if err != nil {
		fmt.Println(err)
	}

	//CHECK PLAZA CODE
	match := bson.M{}
	if reqp.Plaza_code == "all" {
		match = bson.M{"$match": bson.M{"timestamp_tr.time": bson.M{"$gte": fromDate, "$lte": toDate}}}

	} else {
		match = bson.M{"$match": bson.M{"timestamp_tr.time": bson.M{"$gte": fromDate, "$lte": toDate}, "plaza_code": reqp.Plaza_code}}
	}

	pipeline := []bson.M{
		match, // validate between date
		bson.M{"$sort": bson.M{"timestamp_tr.time": 1}},
		bson.M{"$lookup": bson.M{"from": "event_rated", "localField": "uuid_input", "foreignField": "externalid", "as": "event_rated"}},
		bson.M{"$unwind": bson.M{"path": "$event_rated", "preserveNullAndEmptyArrays": true, "includeArrayIndex": "event_rated_index"}},
	}

	pipe := con.Pipe(pipeline)
	resp := []bson.M{}
	err = pipe.All(&resp)

	if err != nil {
		fmt.Println("Errored: %#v \n", err)
	}

	fmt.Println(reflect.TypeOf(result))

	/* loop result Struct; get customer data */
	subscriptionCon := session.DB("orbits-mdm").C("subscription")
	for i, data := range resp {

		//get subscription
		subscriptionPipeline := []bson.M{
			bson.M{"$match": bson.M{"payment_means_number": data["iddevice"]}},
			bson.M{"$lookup": bson.M{"from": "customer-account", "localField": "customer_account_id", "foreignField": "_id", "as": "customer_account"}},
			bson.M{"$unwind": bson.M{"path": "$customer_account", "preserveNullAndEmptyArrays": true, "includeArrayIndex": "customer_account_index"}},
			bson.M{"$lookup": bson.M{"from": "customer", "localField": "customer_account.customer_id", "foreignField": "_id", "as": "customer"}},
			bson.M{"$unwind": bson.M{"path": "$customer", "preserveNullAndEmptyArrays": true, "includeArrayIndex": "customer_index"}},
		}

		pipe2 := subscriptionCon.Pipe(subscriptionPipeline)
		resp2 := []bson.M{}
		err2 := pipe2.All(&resp2)

		if err2 != nil {
			fmt.Println("Errored: %#v \n", err2)
		}

		//push subscription to event_input
		resp[i]["subscription"] = resp2

		//fmt.Println(resp2)

		//fmt.Println(i)
		// fmt.Println(data["iddevice"])
	}

	/* loop result Struct; get plaza data */
	plazaCon := session.DB("orbits-mdm").C("plaza")
	for i, data := range resp {

		//get subscription
		plazaPipeline := []bson.M{
			bson.M{"$match": bson.M{"code": data["plaza_code"]}},
		}

		pipe3 := plazaCon.Pipe(plazaPipeline)
		plaza := Plaza{}
		err3 := pipe3.One(&plaza)

		if err3 != nil {
			fmt.Println("Errored: %#v \n", err3)
		}

		//push subscription to event_input
		resp[i]["plaza_name"] = plaza.Name

		//fmt.Println(resp2)

		//fmt.Println(i)
		// fmt.Println(data["iddevice"])
	}

	//fmt.Println(resp)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "body": resp})

}

/* TIME FORMAT */
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan _2"))
	return []byte(stamp), nil
}

type RecapDataGroup struct {
	Day  string `json:"day" bson:"day"`
	Lane string `json:"lane" bson:"lane"`
	//Month      string `json:"month" bson:"month"`
	Plaza_code string `json:"plaza_code" bson:"plaza_code"`
	Shift      string `json:"shift" bson:"shift"`
	//Year       string `json:"year" bson:"year"`
	//Date string `json:"date" bson:"date"`
}
type RecapData struct {
	RecapDataGroup   `bson:"_id" binding:"dive"`
	Count            int      `json:"count" bson:"count"`
	Plaza_name       string   `json:"plaza_name" bson:"plaza_name"`
	Total_trx_amount int      `json:"total_trx_amount" bson:"total_trx_amount"`
	Datetime         JSONTime `json:"datetime" bson:"datetime"`
}

type Plaza struct {
	Name string `json:"name" bson:"name"`
	Code string `json:"code" bson:"code"`
}

func GetTransactionRecapByDate(c *gin.Context) {
	/* Get mgo session & specify database and collection */
	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB("orbits-transaction").C("event_input")

	// result := []EventInput{}

	var reqp RequestPayload
	/*Bind request payload*/
	errreq := c.Bind(&reqp)
	if errreq != nil {
		fmt.Println(errreq.Error())
	}

	/* parse date string to time object */
	//Start date
	fromDate, err := time.Parse(time.RFC3339, reqp.Start_date)

	if err != nil {
		fmt.Println(err)
	}

	//End date
	toDate, err := time.Parse(time.RFC3339, reqp.End_date)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fromDate)
	fmt.Println(toDate)

	//CHECK PLAZA CODE
	match := bson.M{}

	if reqp.Plaza_code == "all" {
		match = bson.M{"$match": bson.M{"timestamp_tr.time": bson.M{"$gte": fromDate, "$lte": toDate}}}

	} else {
		match = bson.M{"$match": bson.M{"timestamp_tr.time": bson.M{"$gte": fromDate, "$lte": toDate}, "plaza_code": reqp.Plaza_code}}
	}
	fmt.Println(match)
	pipeline := []bson.M{

		match, // validate between date
		bson.M{"$sort": bson.M{"timestamp_tr.time": 1}},
		bson.M{"$lookup": bson.M{"from": "event_rated", "localField": "uuid_input", "foreignField": "externalid", "as": "event_rated"}},
		bson.M{"$unwind": bson.M{"path": "$event_rated", "preserveNullAndEmptyArrays": true, "includeArrayIndex": "event_rated_index"}},
		bson.M{
			"$match": bson.M{
				"event_rated_index": bson.M{"$ne": nil},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"day":        bson.M{"$dayOfYear": "$timestamp_tr.time"},
					"shift":      "$shift",
					"lane":       "$lane",
					"plaza_code": "$plaza_code",
				},
				"total_trx_amount": bson.M{"$sum": "$trx_amount"},
				"count":            bson.M{"$sum": 1},
				"datetime":         bson.M{"$first": "$timestamp_tr.time"},
			},
		},
	}

	dataRecap := []RecapData{}

	pipe := con.Pipe(pipeline)
	resp := []bson.M{}
	err = pipe.All(&dataRecap)

	if err != nil {
		fmt.Println("Errored: %#v \n", err)
	}

	/* loop result Struct */
	plazaCon := session.DB("orbits-mdm").C("plaza")
	for i, data := range dataRecap {
		fmt.Println(data.Plaza_code)
		//get subscription
		plazaPipeline := []bson.M{
			bson.M{"$match": bson.M{"code": data.Plaza_code}},
		}

		plaza := Plaza{}

		pipe2 := plazaCon.Pipe(plazaPipeline)
		err2 := pipe2.One(&plaza)

		if err2 != nil {
			fmt.Println("Errored: %#v \n", err2)
		}
		fmt.Println(plaza)
		fmt.Println(i)
		//push plaza to plaza_name
		dataRecap[i].Plaza_name = plaza.Name
	}

	fmt.Println(resp)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "body": dataRecap})
}
