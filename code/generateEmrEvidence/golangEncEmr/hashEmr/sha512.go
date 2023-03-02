package hashEmr

import (
	"crypto/sha512"
	"fmt"
	"io"
	"os"
)

func Hash_sha512_file(filepath string) (string, error) {
	var hash_value string

	file, err := os.Open(filepath)
	if err != nil {
		return hash_value, err
	}
	defer file.Close()

	my_hash := sha512.New()

	_, err2 := io.Copy(my_hash, file)
	if err2 != nil {
		return hash_value, err2
	}
	hash_byte := my_hash.Sum(nil)

	hash_value = fmt.Sprintf("%x", hash_byte)

	return hash_value, nil

}
