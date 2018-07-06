package customer

import (
	"github.com/gin-gonic/gin"
	"orbits-master-api/models/customer"
	"orbits-master-api/usecases/helper/mongodb"
)

func FindPagination(c *gin.Context) {
	result := []customer.Customer{}
	mongodb.FindPagination(c, "orbits-mdm", "customer", &result)
}
