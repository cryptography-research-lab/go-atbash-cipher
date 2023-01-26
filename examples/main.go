package main

import (
	"fmt"
	atbash_cipher "github.com/cryptography-research-lab/go-atbash-cipher"
)

func main() {

	// 对明文进行加密
	encrypt := atbash_cipher.Encrypt("HELLO")
	fmt.Println(encrypt) // Output: SVOOL

	// 对加密结果进行解密
	decrypt := atbash_cipher.Decrypt(encrypt)
	fmt.Println(decrypt) // Output: HELLO

}
