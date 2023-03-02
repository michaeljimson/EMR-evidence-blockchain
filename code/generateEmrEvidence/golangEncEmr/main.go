package main

import (
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"michaeljimson.com/golangencemr/hashEmr"
	"michaeljimson.com/golangencemr/helpers"
	tojson "michaeljimson.com/golangencemr/toJson"
)

/* var hashMethod = flag.String("s", "sha256", "请输入哈希版本")

func printHash(flag string, str string) {
	switch flag {
	case "SHA256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(str)))
		return
	case "SHA512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(str)))
		return
	case "SHA384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(str)))
		return
	}
}

// Checksum returns the checksum of some data, using a specified algorithm.
// It only returns an error when an invalid algorithm is used. The valid ones
// are MD5, SHA1, SHA224, SHA256, SHA384, SHA512, SHA3224, SHA3256, SHA3384,
// and SHA3512.
func Checksum(algorithm string, data []byte) (checksum string, err error) {
	// default
	var hasher hash.Hash
	switch strings.ToUpper(algorithm) {
	case "MD5":
		hasher = md5.New()
	case "SHA1":
		hasher = sha1.New()
	case "SHA224":
		hasher = sha256.New224()
	case "SHA256":
		hasher = sha256.New()
	case "SHA384":
		hasher = sha512.New384()
	case "SHA512":
		hasher = sha512.New()
	case "SHA3224":
		hasher = sha3.New224()
	case "SHA3256":
		hasher = sha3.New256()
	case "SHA3384":
		hasher = sha3.New384()
	case "SHA3512":
		hasher = sha3.New512()
	default:
		msg := "Invalid algorithm parameter passed go Checksum: %s"
		return checksum, fmt.Errorf(msg, algorithm)
	}
	hasher.Write(data)
	str := hex.EncodeToString(hasher.Sum(nil))
	return str, nil
} */
//decodePEMFile reads and decodes generic PEM files.
func decodePEMFile(filePath string) ([]byte, error) {
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
	fmt.Printf("p: %v\n", p)
	if p == nil {
		return nil, fmt.Errorf("no pem block found")
	}
	fmt.Printf("p: %v\n", p.Bytes)
	return p.Bytes, nil
}

func main() {

	/* flag.Parse()
	var i string
	fmt.Print("请输入要加密的字符串:")
	fmt.Scanln(&i)
	printHash(strings.ToUpper(*hashMethod), i) */

	/* Testfile := "encryptFile_test"
	infile, inerr := os.Owpen(Testfile)
	if inerr == nil {
		/* startTime := time.Now()
		md5h := md5.New()
		io.Copy(md5h, infile)
		fmt.Printf("%x %s\n", md5h.Sum([]byte("")), Testfile)
		fmt.Printf("time.Since(startTime): %v\n", time.Since(startTime))
		/* infile.Seek(0, 0)
		startTime := time.Now()
		for i := 0; i < 10; i++ {
			sha512h := sha512.New()
			io.Copy(sha512h, infile)
			fmt.Printf("%x %s\n", sha512h.Sum(

				[]byte("")), Testfile)
			defer fmt.Printf("time.Since(startTime): %v\n", time.Since(startTime))
		}
	} else {
		fmt.Println(inerr)
		os.Exit(1)
	} */

	/* startTime := time.Now()
	helpers.EncryptFile("test1.xml", "test")
	helpers.DecryptFile("encryptFile_test", "qtest1.xml")
	fmt.Printf("time.Since(startTime): %v\n", time.Since(startTime))
	*/

	/* startTime := time.Now()
	var i string
	fmt.Print("请输入要加密的字符串:")
	fmt.Scanln(&i)
	var data []byte = []byte(i)
	fmt.Println(Checksum(strings.ToUpper(*hashMethod), data))
	fmt.Printf("time.Since(startTime): %v\n", time.Since(startTime)) */
	file_path := "EmrTest1.xml"

	hash, err := hashEmr.Hash_sha512_file(file_path)
	fmt.Printf("hash: %T\n", hash)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s sha512 is:\n%s\n", file_path, hash)
	//fmt.Printf("time.Since(startTime): %v\n", time.Since(startTime))
	/*
		helpers.GeneratePriAndPubKeys(elliptic.P256())
		txt := []byte(hash)
		r, s := helpers.SignByEcc("ecc_private.pem", txt)
		fmt.Printf("time.Since(startTime): %v\n", time.Since(startTime))
		fmt.Printf("r: %v\n", string(r))
		fmt.Printf("s: %v\n", string(s))
		fmt.Printf("sign verify result: %v\n", helpers.VerifyByEcc("ecc_public.pem", txt, r, s))
	*/

	// helpers.Ed25519Sign(hash)

	/* 	var e tojson.EmrToJsonslice
	   	e.EmrToJsons = append(e.EmrToJsons, tojson.EmrToJson{EmrName: "张三", EmrHash: sign, EmrSignPubKey: })
	*/

	/* if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("%s : generate pen formatted ed25519 keys\n", os.Args[0])
		fmt.Println("    provide a single argument for the private key name")
		fmt.Println("    the public key name will have '.pub' appended")
		os.Exit(1)
	} */
	// 生成ed25519 公私钥
	/* 	FileBaseName := "ec25519"

	   	if err := helpers.GenerateSaveEd25519(FileBaseName); err != nil {
	   		fmt.Printf("err: %v\n", err)
	   		os.Exit(1)
	   	}

	   	var priKey []byte
	   	priKey, err = decodePEMFile("ec25519.priv") */
	//Create Pub/PrivKey and signature
	//helpers.GenerateED25519AndWrite("ED25519privKey.pem", "ED25519pubKey.pem")
	privKey, _ := helpers.GetPrivateKey("ED25519privKey.pem")

	signature, _ := privKey.SignByED25519(hash)
	//signature := helpers.Ed25519Sign(hash)
	println(fmt.Sprintf("signature is %s\n", signature))

	/* decryptSignature, _ := helpers.AesDecrypt(encryptSignature, []byte("ABCDABCDABCDABCD"))
	println(fmt.Sprintf("decryptSignature is %s\n", string(decryptSignature))) */
	/* //Verify signature
	pubKey, _ := helpers.GetPublicKey("ED25519pubKey.pem")

	ok, _ := pubKey.VerifyByED25519([]byte(hash), signature)

	if ok {
		fmt.Println("古德古德")
	} else {
		fmt.Println("完了完了")
	}
	*/

	//encrySignature
	encryptSignature, _ := helpers.AesEncrypt([]byte(signature), []byte("ABCDABCDABCDABCD"))
	println(fmt.Sprintf("encryptSignature is %x\n", encryptSignature))
	jsonEmrHash := hex.EncodeToString(encryptSignature)
	//write to json file
	jsonemr := tojson.EmrToJson{}
	jsonemr.ID = "111111111111111111"
	jsonemr.EmrHash = jsonEmrHash
	jsonemr.EmrName = "zhangsan"
	jsonemr.DoctorID = "123123123123"
	jsonemr.DoctorName = "lisi"
	jsonemr.EmrSignPubKey = "MCowBQYDK2VwAyEAScigIpHiqYjpQ8UbbtwPDAG+bojmbT/tr1gPGthCnPg="
	jsonemr.EmrDetail = "This electronic medical is belong to zhangsan "
	jsonemr.EmrTime = time.Now()

	jsonemrByte, err := json.MarshalIndent(jsonemr, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(jsonemr.EmrName+".json", jsonemrByte, 0644)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(jsonemrByte))

	decodeJsonFile, err := ioutil.ReadFile(jsonemr.EmrName + ".json")
	if err != nil {
		log.Fatal(err)
	}
	var jsonDecodeEmr = tojson.EmrToJson{}
	err = json.Unmarshal(decodeJsonFile, &jsonDecodeEmr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jsonDecodeEmr)
}
