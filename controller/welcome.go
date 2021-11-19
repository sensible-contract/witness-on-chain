package controller

import (
	"log"
	"net/http"
	"witness-on-chain/model"

	"github.com/gin-gonic/gin"
)

// Witness-On-Chain
func WitnessOnChain(ctx *gin.Context) {
	log.Printf("Witness-On-Chain enter")

	ctx.JSON(http.StatusOK, model.Welcome{
		PublicKey: model.PublicKey{
			Rabin: rb.PubKeyHex,
		},
		Contact: "hi@witnessonchain.com",
		Message: "Welcome to WitnessOnChain!",
	})
}
