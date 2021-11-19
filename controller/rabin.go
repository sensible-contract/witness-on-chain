package controller

import (
	"os"
	"witness-on-chain/lib/rabin"
)

var rb = new(rabin.Rabin)

func init() {
	pString := os.Getenv("PINT")
	qString := os.Getenv("QINT")

	rb.Init(pString, qString)
}
