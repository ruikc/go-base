package main


import (
	"flag"
	"github.com/gin-gonic/gin"
)

func main()  {
	flag.Parse()
	gin.SetMode(Sc.Model)
	router := gin.Default()
	NewRouter(router)
	router.Run(":" + Sc.Port)
}
