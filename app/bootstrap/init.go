package bootstrap

import "github.com/gin-gonic/gin"

func Init() {
	loadEnv()
	db := connectDb()
	defer db.Close()

	router := gin.Default()

	userController := InitUserController(db)
	linkController := InitLinkController(db)
	linkHistoryService := InitLinkHistoryService(db)
	statController := InitStatController(db)

	router.GET("/redirect/:shortUrl", linkController.RedirectToUrl, linkHistoryService.WriteHistoryLinkConversion)

	v1PublicApi := router.Group("/api/v1")

	//Access without auth
	{
		v1PublicApi.POST("/register", userController.Register)
	}

	basicAuth := InitBasicAuth(db)
	v1Api := router.Group("/api/v1", basicAuth.BasicAuth())

	{
		v1Api.GET("/users/self", userController.Self)
		v1Api.POST("/links", linkController.CreateShortLink)
		v1Api.GET("/links", linkController.GetSelfLinks)
		v1Api.GET("/links/:linkId", linkController.GetLink)
		v1Api.DELETE("/links/:linkId", linkController.DeleteLink)
		v1Api.GET("/stat/top", statController.Top)
		v1Api.GET("/stat/conversion/graph", statController.Conversion)
	}

	_ = router.Run()
}
