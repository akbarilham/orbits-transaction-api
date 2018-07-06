package mongodb

import (
	//"github.com/tidwall/gjson"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func FindCollectionById(c *gin.Context, id bson.ObjectId, db string, collection string, model interface{}) interface{} {
	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB(db).C(collection)

	err := con.Find(bson.M{"_id": id}).One(model)

	if err != nil {
		log.Panicln(err)
	}

	return model

}
