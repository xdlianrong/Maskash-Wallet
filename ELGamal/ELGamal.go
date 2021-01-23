package ELGamal

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

type PrivStr struct {
	G1         string `json:"G1"`
	G2         string `json:"G2"`
	P          string `json:"P"`
	Publickey  string `json:"publickey"`
	Privatekey string `json:"privatekey"`
}

// 学长写的
// PublicKey 公钥
type PublicKey struct {
	G1, G2, P, H *big.Int
}

// PrivateKey 私钥
type PrivateKey struct {
	PublicKey
	X *big.Int
}

type Account struct {
	Pub  PublicKey  `json:"Pub"`
	Priv PrivateKey `json:"Priv"`
	Info struct {
		Name    string `json:"Name"`
		ID      string `json:"ID"`
		Hashky  string `json:"Hashky"`
		ExtInfo string `json:"ExtInfo"`
	} `json:"Info"`
}

func GenerateAccount(randString string, name string, id string, extInfo string) Account {
	pub, priv, _ := GenerateKeys(randString)
	fmt.Println("生成账户"+name, "私钥：", priv.X.String())
	return Account{
		Pub:  pub,
		Priv: priv,
		Info: struct {
			Name    string `json:"Name"`
			ID      string `json:"ID"`
			Hashky  string `json:"Hashky"`
			ExtInfo string `json:"ExtInfo"`
		}{
			Name:    name,
			ID:      id,
			Hashky:  pub.H.String(),
			ExtInfo: extInfo,
		},
	}
}

func GenerateKeys(info string) (pub PublicKey, priv PrivateKey, err error) {
	// 本函数用于根据用户信息 string 生成一对公私钥 pub 和 priv
	// 从质数表中随机选择大质数P
	var error_bool bool
	pub.P, error_bool = new(big.Int).SetString(select_prime(), 16)
	if !error_bool {
		return
	}
	priv.P = pub.P

	// 使用string Hash生成G1
	pub.G1 = new(big.Int)
	HashInfoBuf := sha256.Sum256([]byte(info))
	HashInfo := HashInfoBuf[:]
	pub.G1.SetBytes(HashInfo)
	pub.G1.Mod(pub.G1, pub.P)
	for {
		gcd := new(big.Int).GCD(nil, nil, pub.G1, pub.P).Int64()
		if gcd == 1 {
			break
		}
		pub.G1.Sub(pub.G1, new(big.Int).SetInt64(1))
	}
	priv.G1 = pub.G1

	// 使用string time Hash生成G2
	pub.G2 = new(big.Int)
	now := time.Now().Unix()
	stringNow := []byte(strconv.FormatInt(now, 10))
	HashInfo = append(HashInfo, stringNow...)
	pub.G2.SetBytes(HashInfo)
	pub.G2.Mod(pub.G2, pub.P)
	for {
		gcd := new(big.Int).GCD(nil, nil, pub.G2, pub.P).Int64()
		if gcd == 1 {
			break
		}
		pub.G2.Sub(pub.G2, new(big.Int).SetInt64(1))
	}
	priv.G2 = pub.G2

	// 随机选择私钥 X
	priv.X = new(big.Int)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	priv.X.Rand(rnd, pub.P)

	// 计算公钥 H
	pub.H = new(big.Int)
	pub.H.Exp(pub.G2, priv.X, pub.P)
	priv.H = pub.H

	return
}

func (account Account) KeyToString() (privStr PrivStr) {
	privStr.G1 = fmt.Sprintf("%0*x", 64, account.Pub.G1)
	privStr.G2 = fmt.Sprintf("%0*x", 64, account.Pub.G2)
	privStr.P = fmt.Sprintf("%0*x", 64, account.Pub.P)
	privStr.Publickey = fmt.Sprintf("%0*x", 64, account.Pub.H)
	privStr.Privatekey = fmt.Sprintf("%0*x", 64, account.Priv.X)
	return
}
