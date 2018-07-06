package mongodb

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	//"log"
	"math"
	"net/http"
	//"orbits-master-api/models/customer"
)

type Options struct {
	Filter       []Filter `json:"filter"`
	Current_page int      `json:"current_page" binding:"required"`
	Limit        int      `json:"limit" binding:"required"`
}
type Filter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ResultPagination struct {
	Body        interface{} `bson:",inline"`
	TotalRows   int
	TotalPages  float64
	CurrentPage int
	Perpage     int
	LastPage    float64
	Indexs      []int
	Code        int `json:"code"`
}
type ResultCount struct {
	Count int
}

// func (r *ResultPagination) SetBody(body *Person) {
// 	r.Body = body
// }

func (r *ResultPagination) SetTotalRows(totalrows int) {
	r.TotalRows = totalrows
}

func (r *ResultPagination) SetTotalPages(totalpages float64) {
	r.TotalPages = totalpages
}

func (r *ResultPagination) SetCurrentPage(currentpage int) {
	r.CurrentPage = currentpage
}

func (r *ResultPagination) SetPerpage(perpage int) {
	r.Perpage = perpage
}

func (r *ResultPagination) SetLastPage(lastpage float64) {
	r.LastPage = lastpage
}

func FindPagination(c *gin.Context, db string, collection string, model interface{}) {
	json, _ := ioutil.ReadAll(c.Request.Body)

	//key := gjson.Get(string(json), "options.filter")
	Current_page := gjson.Get(string(json), "options.current_page").Int()
	Limit := gjson.Get(string(json), "options.limit").Int()
	Filter := gjson.Get(string(json), "options.filter")

	session := c.MustGet("mongoSession").(*mgo.Session)
	con := session.DB(db).C(collection)

	var current_page float64
	var limit float64
	var offset float64
	current_page = float64(Current_page)
	limit = float64(Limit)
	offset = math.Floor((current_page * limit) - limit)

	/* Query goes here */
	//prepare filter
	param := bson.M{}

	/* if Filter Not empty then push to param */
	if Filter.String() != "" {
		/* push value to param */
		Filter.ForEach(func(key, value gjson.Result) bool {
			param[key.String()] = bson.RegEx{value.String(), ""}
			// fmt.Println(key.String())
			return true // keep iterating
		})
	}
	// result := model

	err := con.Find(param).Limit(int(limit)).Skip(int(offset)).All(model)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}
	/* Count document */
	count, err := con.Find(param).Count()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}

	// fmt.Println(count)

	/* Prepare result */
	r := ResultPagination{}
	r.Body = model
	r.SetTotalRows(count)
	r.SetTotalPages(math.Ceil(float64(count) / float64(limit)))
	r.SetCurrentPage(int(current_page))
	r.SetPerpage(int(limit))
	r.SetLastPage(math.Ceil(float64(count) / float64(limit)))

	range_index := 5
	total_pages := int(r.TotalPages)
	var index []int
	if total_pages <= range_index {
		for p := 1; p <= total_pages; p++ {
			index = append(index, p)
		}
		r.Indexs = index
	} else {
		minp := 1
		maxp := 0
		//set minp
		if r.CurrentPage > range_index {
			minp = r.CurrentPage - range_index
		}
		//set maxp
		if (total_pages - r.CurrentPage) > range_index {
			maxp = r.CurrentPage + range_index
		} else {
			maxp = (r.CurrentPage + (total_pages - r.CurrentPage))
		}

		for p := minp; p <= maxp; p++ {
			index = append(index, p)
		}
		r.Indexs = index
	}

	r.Code = 200
	c.JSON(200, r)
	//fmt.Println("Phone:", result.Phone)

}
