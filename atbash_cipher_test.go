package atbash_cipher

import (
	"fmt"
	"testing"
)

func TestDecrypt(t *testing.T) {
	plaintext := "abcdefg"
	encryptResult := Encrypt(plaintext)
	fmt.Println(encryptResult)
	decryptResult := Decrypt(encryptResult)
	fmt.Println(decryptResult)
}
