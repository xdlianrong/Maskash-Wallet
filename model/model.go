package model

import (
	"wallet/ELGamal"
	"wallet/controllers"
)

type NewWallet struct {
	Name string `json:"name" form:"name"`
	Id 	 string `json:"id" form:"id"`
	Str  string `json:"str" form:"str"`
}

type BctoEx struct {
	G1     string `json:"g1"`
	G2     string `json:"g2"`
	P      string `json:"p"`
	H      string `json:"h"`
	Amount ELGamal.Account `json:"amount"`
}

type ExchangeCoin struct {
	Receipt controllers.Receipt `json:"receipt"`
	Priv	ELGamal.PrivateKey	`json:"priv"`
}