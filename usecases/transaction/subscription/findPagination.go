package subscription

import (
	"github.com/gin-gonic/gin"
	"orbits-master-api/models/subscription"
	"orbits-master-api/usecases/helper/mongodb"
)

func FindPagination(c *gin.Context) {
	result := []subscription.Subscription{}
	mongodb.FindPagination(c, "orbits-mdm", "subscription", &result)
}
