package helpers

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/* func Ed25519Sign(msg string) []byte {
	var publicKey []byte
	var privateKey []byte

	publicKey, privateKey, _ = ed25519.GenerateKey(rand.Reader)

	msgByte := []byte(msg)

	signature := ed25519.Sign(privateKey, msgByte)
	verify := ed25519.Verify(publicKey, msgByte, signature)

	fmt.Println("公钥：", publicKey)
	fmt.Println("私钥：", privateKey)
	fmt.Println("签名：", signature)
	fmt.Println("验签：", verify)
	return signature
}

func GenerateSaveEd25519(fb string) error {
	var (
		err   error
		b     []byte
		block *pem.Block
		pub   ed25519.PublicKey
		priv  ed25519.PrivateKey
	)
	pub, priv, err = ed25519.GenerateKey(rand.Reader)

	fmt.Printf("priv: %v\n", priv)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	b, err = x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return err
	}

	block = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: b,
	}

	err = ioutil.WriteFile(fb+".priv", pem.EncodeToMemory(block), 0600)
	if err != nil {
		return err
	}

	b, err = x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: b,
	}
	filename := fb + ".pub"
	err = ioutil.WriteFile(filename, pem.EncodeToMemory(block), 0644)
	return err
}

func UseKeyToVerify(fb string) error {
	var (
		err error
		//b     []byte
		block *pem.Block
		//pub   ed25519.PublicKey
		//priv  ed25519.PrivateKey
	)
	pubKeyPem := fb + ".pub"
	content, err := ioutil.ReadFile(pubKeyPem)
	if err != nil {
		panic(fmt.Sprintf("read private key file fail: %s", err))
	}

	block, _ = pem.Decode(content)
	prikey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic(fmt.Sprintf("parse private key fail: %s", err))
	}
	println(prikey)

	return nil
} */

/*
func GenerateNewEd25519() error {

	pub, priv, _ := ed25519.GenerateKey(rand.Reader)
	publicKey, _ := ssh.NewPublicKey(pub)
	fmt.Printf("priv: %v\n", edkey.MarshalED25519PrivateKey(priv))
	pemKey := &pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: edkey.MarshalED25519PrivateKey(priv),
	}

	privateKey := pem.EncodeToMemory(pemKey)
	authorizedKey := ssh.MarshalAuthorizedKey(publicKey)

	_ = ioutil.WriteFile("id_ed25519", privateKey, 0600)
	err := ioutil.WriteFile("id_ed25519.pub", authorizedKey, 0644)

	return err
} */

func GenerateED25519AndWrite(private, public string) {
	//GenerateKey will generate the private and public key pairs using
	//rand.Rander as source of entropy
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	//Create file and write public key
	pubOut, err := os.OpenFile(public, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Failed to create %s file: %s", private, err)
	}
	pubBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		log.Fatalf("Unable to marshal public key: %v", err)
	}

	// Encode public key using PEM format
	if err := pem.Encode(pubOut, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	}); err != nil {
		log.Fatalf("Failed to write data to %s file: %s", public, err)
	}
	if err := pubOut.Close(); err != nil {
		log.Fatalf("Error closing %s file: %s", public, err)
	}

	//Create file and write private key
	keyOut, err := os.OpenFile(private, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Failed to create %s file: %s", private, err)
	}
	privBytes, err := x509.MarshalPKCS8PrivateKey(privKey)
	if err != nil {
		log.Fatalf("Unable to marshal private key: %v", err)
	}
	if err := pem.Encode(keyOut, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privBytes,
	}); err != nil {
		log.Fatalf("Failed to write data to %s file: %s", private, err)
	}
	if err := keyOut.Close(); err != nil {
		log.Fatalf("Error closing %s file: %s", private, err)
	}

}

func ED25519DecodePEMFile(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	p, _ := pem.Decode(buf)
	if p == nil {
		return nil, fmt.Errorf("no pem block found")
	}
	return p.Bytes, nil
}

// PrivateKey is ed25519.PrivateKey.
type PrivateKey ed25519.PrivateKey

//GetPrivateKey reads the private key from input file and
//returns the initialized PrivateKey.
func GetPrivateKey(privateKey string) (PrivateKey, error) {
	p, _ := ED25519DecodePEMFile(privateKey)
	key, err := x509.ParsePKCS8PrivateKey(p)
	if err != nil {
		return nil, err
	}
	edKey, ok := key.(ed25519.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("key is not ed25519 key")
	}
	return PrivateKey(edKey), nil
}

//Sign reads the input file and computes the ED25519 signature
//using the private key
func (p PrivateKey) SignByED25519(s string) (string, error) {
	/* f, err := os.Open(path)
	if err != nil {
		return "Open error", err
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return "buffer error", err
	} */
	signature := ed25519.Sign(ed25519.PrivateKey(p), []byte(s))
	return hex.EncodeToString(signature), nil
}

//PublicKey is ed25519.PublicKey
type PublicKey ed25519.PublicKey

//return the initialized PublicKey
func GetPublicKey(publicKey string) (PublicKey, error) {
	p, _ := ED25519DecodePEMFile(publicKey)
	key, err := x509.ParsePKIXPublicKey(p)
	if err != nil {
		return nil, err
	}
	edKey, ok := key.(ed25519.PublicKey)
	if !ok {
		return nil, fmt.Errorf("key is not ed25519 key")
	}
	return PublicKey(edKey), nil
}

//Verify checks that input singature is valid. That is, if
//input file was signed by private key corresponding to input
//public key.
func (p PublicKey) VerifyByED25519(decryptsignature []byte, signature string) (bool, error) {
	/* f, err := os.Open(file)
	if err != nil {
		return false, err
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return false, err
	} */
	byteSign, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}
	ok := ed25519.Verify(ed25519.PublicKey(p), decryptsignature, byteSign)
	return ok, nil
}
                      