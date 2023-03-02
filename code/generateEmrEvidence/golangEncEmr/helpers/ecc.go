package helpers

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
)

func GeneratePriAndPubKeys(curve elliptic.Curve) {
	priKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(fmt.Sprintf("generate private key fail: %s", err))
	}

	derTxt, err := x509.MarshalECPrivateKey(priKey)
	if err != nil {
		panic(fmt.Sprintf("marshal private key fail: %s", err))
	}

	block := pem.Block{Type: "ECC PRIVATE KEY", Bytes: derTxt}
	file, err := os.Create("ecc_private.pem")
	if err != nil {
		panic(fmt.Sprintf("create private key file fail: %s", err))
	}

	defer file.Close()

	if err := pem.Encode(file, &block); err != nil {
		panic(fmt.Sprintf("write private key fail: %s", err))

	}

	pubDerTxt, err := x509.MarshalPKIXPublicKey(&priKey.PublicKey)
	if err != nil {
		panic(fmt.Sprintf("marshal public key fail: %s", err))
	}

	pubBlock := pem.Block{Type: "ECC PRIVATE KEY", Bytes: pubDerTxt}
	pubFile, err := os.Create("ecc_public.pem")
	if err != nil {
		panic(fmt.Sprintf("create private key file fail: %s", err))

	}
	defer pubFile.Close()

	if err := pem.Encode(pubFile, &pubBlock); err != nil {
		panic(fmt.Sprintf("write private key fail: %s", err))
	}
}

func SignByEcc(priFile string, plaintext []byte) (rTxt, sTxt []byte) {
	content, err := ioutil.ReadFile(priFile)
	if err != nil {
		panic(fmt.Sprintf("read public key file fail: %s", err))
	}

	block, _ := pem.Decode(content)
	priKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic(fmt.Sprintf("parse private key fail: %s", err))
	}

	h := sha256.New()
	h.Write(plaintext)

	r, s, err := ecdsa.Sign(rand.Reader, priKey, h.Sum(nil))
	if err != nil {
		panic(fmt.Sprintf("sign fail: %s", err))
	}

	rTxt, _ = r.MarshalText()
	sTxt, _ = s.MarshalText()

	return rTxt, sTxt

}

func VerifyByEcc(pubFile string, plaintext, rTxt, sTxt []byte) bool {

	content, err := ioutil.ReadFile(pubFile)

	if err != nil {
		panic(fmt.Sprintf("read public key file fail: %s", err))
	}
	block, _ := pem.Decode(content)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(fmt.Sprintf("parse public key fail: %s", err))
	}
	pubKey := pubInterface.(*ecdsa.PublicKey)
	r, s := new(big.Int), new(big.Int)
	if err := r.UnmarshalText(rTxt); err != nil {
		panic(fmt.Sprintf("unmarshal r fail: %s", err))
	}
	if err := s.UnmarshalText(sTxt); err != nil {
		panic(fmt.Sprintf("unmarshal s fail: %s", err))
	}

	h := sha256.New()
	h.Write(plaintext)

	return ecdsa.Verify(pubKey, h.Sum(nil), r, s)
}
