package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.com/banking-app/helpers"
	"google.com/banking-app/models"
	"google.com/banking-app/server"
)

type Controller struct{}

func (c *Controller) LoginController(ctx *gin.Context) {

	reqBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		er := &helpers.Error{ErrorMessage: "cannot read request body", Err: err}
		er.HandleErr()
	}

	req := new(models.LoginRequest)

	if err := json.Unmarshal(reqBody, req); err != nil {
		er := &helpers.Error{ErrorMessage: "invalid request body", Err: err}
		er.HandleErr()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})

	}

	srv := new(server.Server)

	if response := srv.LoginService(req.Username, req.Password); response["details"] == "" || response["token"] == "" {
		er := &helpers.Error{ErrorMessage: "internal server error"}
		er.HandleErr()
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})

	} else {
		ctx.JSON(http.StatusOK, response)

	}
}

func (c *Controller) RegisterController(ctx *gin.Context) {
	reqBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		er := &helpers.Error{ErrorMessage: "cannot read request body", Err: err}
		er.HandleErr()
	}

	req := new(models.RegisterRequest)

	if err := json.Unmarshal(reqBody, req); err != nil {
		er := &helpers.Error{ErrorMessage: "invalid request body", Err: err}
		er.HandleErr()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
	}

	srv := new(server.Server)

	if response := srv.RegisterService(req.Username, req.Email, req.Password); response["details"] == "" || response["token"] == "" {
		er := &helpers.Error{ErrorMessage: "internal server error"}
		er.HandleErr()
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})

	} else {
		ctx.JSON(http.StatusOK, response)

	}
}

func (c *Controller) GetUserController(ctx *gin.Context) {

	id := ctx.Param("id")
	token := ctx.GetHeader("Authorization")

	if isValid := helpers.ValidateToken(id, token); isValid {
		srv := new(server.Server)
		if response := srv.GetUserService(id, token); response["details"] == "" || response["token"] == "" {
			er := &helpers.Error{ErrorMessage: "internal server error"}
			er.HandleErr()
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})

		} else {
			ctx.JSON(http.StatusOK, response)
		}
	} else {
		log.Println("invalid token")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid token",
		})
	}
}

func (c *Controller) TransactionController(ctx *gin.Context) {

	reqBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		er := &helpers.Error{ErrorMessage: "cannot read request body", Err: err}
		er.HandleErr()
	}

	token := ctx.GetHeader("Authorization")

	req := new(models.TransactionRequest)

	if err := json.Unmarshal(reqBody, req); err != nil {
		er := &helpers.Error{ErrorMessage: "invalid request body", Err: err}
		er.HandleErr()
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
	}
	if isValid := helpers.ValidateToken(fmt.Sprint(req.UserId), token); isValid {
		srv := new(server.Server)

		if response := srv.TransactService(req.UserId, req.From, req.To, req.Amount, token); response["details"] == "" || response["token"] == "" {
			er := &helpers.Error{ErrorMessage: "internal server error"}
			er.HandleErr()
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})

		} else {
			ctx.JSON(http.StatusOK, response)

		}
	} else {
		log.Println("invalid token")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid token",
		})
	}
}

func (c *Controller) GetMyTransactionsController(ctx *gin.Context) {

	token := ctx.GetHeader("Authorization")
	userId := ctx.Param("id")

	if isValid := helpers.ValidateToken(userId, token); isValid {
		srv := new(server.Server)

		if response := srv.GetMyTransactionsService(userId, token); response["details"] == "" || response["token"] == "" {
			er := &helpers.Error{ErrorMessage: "internal server error"}
			er.HandleErr()
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})

		} else {
			ctx.JSON(http.StatusOK, response)
		}
	} else {
		log.Println("invalid token")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid token",
		})
	}
}
