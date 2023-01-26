package atbash_cipher

import (
	"unicode/utf8"
)

// 因为是对称的，所以只需要做一张映射表就可以了，映射表是固定的
var mappingTable = make(map[rune]rune, 0)

func init() {
	for index := 0; index < 26; index++ {
		fromCharacter := rune('A' + index)
		toCharacter := rune('A' + 25 - index)
		// 同时做一个大小写的映射，通过空间换时间避免掉判断大小写
		mappingTable[fromCharacter] = toCharacter
		mappingTable[fromCharacter+32] = toCharacter + 32
	}
}

// Encrypt 对明文进行atbash加密，英文字符会进行映射，其它字符会保持原样
func Encrypt(plaintext string) string {
	runeSlice := make([]rune, utf8.RuneCountInString(plaintext))
	for index, character := range plaintext {
		mappingToCharacter, exists := mappingTable[character]
		if exists {
			runeSlice[index] = mappingToCharacter
		} else {
			runeSlice[index] = character
		}
	}
	return string(runeSlice)
}

// Decrypt 解密
func Decrypt(ciphertext string) string {
	return Encrypt(ciphertext)
}
