package controller

import (
	"encoding/binary"
	"encoding/hex"
	"log"
	"net/http"
	"time"
	"witness-on-chain/lib/utils"
	"witness-on-chain/model"

	"github.com/gin-gonic/gin"
)

// SignTimestamp
func SignTimestamp(ctx *gin.Context) {
	log.Printf("SignTimestamp enter")

	digest := make([]byte, 4)
	now := time.Now().UTC()
	binary.LittleEndian.PutUint32(digest, uint32(now.Unix()))

	sig, padding := rb.Sign(digest)
	if len(sig) == 0 {
		log.Printf("padding over 256 times.")
		ctx.JSON(http.StatusInternalServerError, model.Response{Code: -1, Msg: "padding over 256 times"})
		return
	}
	ctx.JSON(http.StatusOK, model.TimestampResponse{
		Timestamp: int(now.Unix()),
		UTC:       now.Format("2006-01-02 15:04:05"),
		Digest:    hex.EncodeToString(digest),
		Signatures: model.Signatures{
			Rabin: model.Signature{
				PublicKey: "",
				// Signature: hex.EncodeToString(sig), // SigBE
				Signature: hex.EncodeToString(utils.ReverseBytes(sig)), // SigLE
				Padding:   hex.EncodeToString(padding),
			},
		},
	})
}
