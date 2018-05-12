// Package infra constains Structs and Methods which represent the actual implementation of all the
// interfaces, it also has controllers and routes specific methods and structs which help in connecting
// the main app(infrastructure) with the rest of the applications' usecases and domain
package infra

import (
	"github.com/gin-gonic/gin"
)

// AddRoutes takes in the router and the app objects and add the app's specific handlers
// to different routes.
func AddRoutes(router *gin.Engine, baseUrl string, app *App) {
	// Index route
	router.GET("/", app.IndexHandler)

	// Transaction related group Routes
	transactionGroup := router.Group(baseUrl + "/transaction")
	{
		transactionGroup.POST("/", app.CheckTransactionHandler)
	}
}
