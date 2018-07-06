package inventory

import (
	"github.com/gin-gonic/gin"
	"orbits-master-api/models/inventory"
	"orbits-master-api/usecases/helper/mongodb"
)

func FindPagination(c *gin.Context) {
	result := []inventory.Inventory{}
	mongodb.FindPagination(c, "orbits-mdm", "inventory", &result)
}
