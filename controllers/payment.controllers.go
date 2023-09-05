package controllers

import (
	"net/http"
	"project1/interfaces"
	"project1/models"


	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService interfaces.Ipayment
}

func InitTransactionController(profileService interfaces.Ipayment) TransactionController {
	return TransactionController{profileService}
}

func (t *TransactionController) CreatePayment(ctx *gin.Context) {
	var transactions models.Paymentscard
	if err := ctx.ShouldBindJSON(&transactions); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newProfile, err := t.TransactionService.CreatePayment(float64(transactions.CardNo), int(transactions.Cvvverified), transactions.Balance)

	if err != nil {


		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProfile})
}
