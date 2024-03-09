package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "server/docs"

	swaggerFiles "github.com/swaggo/files"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
//	@schemes					http https
func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET(":id", ShowAccount)
			accounts.GET("", ListAccounts)
		}
		//...
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := r.Run(); err != nil { // listen and serve on 0.0.0.0:8080
		fmt.Printf("%+v\n", err)
	}
}

// ShowAccount godoc
//
//	@Summary		Show an account
//	@Description	get string by ID
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"
//	@Success		200	{object}	Account
//	@Failure		400	{string}	string	"error"
//	@Failure		404	{string}	string	"error"
//	@Failure		500	{string}	string	"error"
//	@Router			/accounts/{id} [get]
func ShowAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

// ListAccounts godoc
//
//	@Summary		List accounts
//	@Description	get accounts
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			q	query		string	false	"name search by q"	Format(email)
//	@Success		200	{array}		Account
//	@Failure		400	{string}	string	"error"
//	@Failure		404	{string}	string	"error"
//	@Failure		500	{string}	string	"error"
//	@Router			/accounts [get]
func ListAccounts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

type Account struct {
	ID   int    `json:"id" example:"1"`
	Name string `json:"name" example:"account name"`
}
