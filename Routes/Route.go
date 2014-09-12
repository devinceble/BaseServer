package Routes

import "github.com/gin-gonic/gin"

//Main Route for Server
func Route() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{"Welcome": "To BaseStructure"})
	})
	r.GET("/api/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{"user": "Sample"})
	})
	r.Static("/client", "Client/")
	return r
}
