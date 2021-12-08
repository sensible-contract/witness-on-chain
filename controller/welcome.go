package controller

import (
	"encoding/hex"
	"log"
	"net/http"
	"witness-on-chain/lib/utils"
	"witness-on-chain/model"

	"github.com/gin-gonic/gin"
)

// Witness-On-Chain
func WitnessOnChain(ctx *gin.Context) {
	log.Printf("Witness-On-Chain enter")

	ctx.JSON(http.StatusOK, model.Welcome{
		PublicKey: model.PublicKey{
			Rabin: hex.EncodeToString(utils.ReverseBytes(rb.PubKey)),
		},
		Contact: "hi@witnessonchain.com",
		Message: "Welcome to WitnessOnChain!",
	})
}
