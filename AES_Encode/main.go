package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func main() {
	key := []byte("jfprj  pjtpj2594") //key與iv必須為128(16字),256,512bits的字串
	//       	   1234567890123456
	// AES加密過都固定21個A，不管怎麼換key都一樣
	msg := "pkhgdr"

	//BASE64
	//str := base64.StdEncoding.EncodeToString([]byte(msg))
	//fmt.Println(str)
	//decode, _ := base64.StdEncoding.DecodeString(str)
	//fmt.Println(string(decode))

	encrypt, err := EncryptAes(key, msg)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(encrypt)
	//URLEncoded
	//escapeurl := url.QueryEscape(encrypt)
	//fmt.Println(escapeurl)
	//unescapeurl, _ := url.QueryUnescape(escapeurl)
	//fmt.Println(unescapeurl)

	decrypt, err := DecryptAes(key, encrypt)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(decrypt)

	//xorcipherkey := "hongyi"
	//data := "112"
	//encrypt = EncryptDecrypt(data, xorcipherkey)
	//fmt.Println(encrypt)
	//decrypt = EncryptDecrypt(encrypt, xorcipherkey)
	//fmt.Println(decrypt)
}

func EncryptAes(key []byte, message string) (string, error) {
	byteMsg := []byte(message)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	ciphertext := make([]byte, aes.BlockSize+len(byteMsg))
	iv := ciphertext[:aes.BlockSize]

	stream := cipher.NewCFBEncrypter(block, []byte(iv))
	stream.XORKeyStream(ciphertext[aes.BlockSize:], byteMsg)
	// byte陣列 ascii碼
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptAes(key []byte, message string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", fmt.Errorf("could not base64 decode: %v", err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("could not create new cipher: %v", err)
	}
	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("invalid ciphertext block size")
	}

	//128,256,512bits才可用(iv,key)
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	return string(cipherText), nil
}

//xor cipher
func EncryptDecrypt(input, key string) (output string) {
	kL := len(key)
	for i := range input {
		output += string(input[i] ^ key[i%kL])
	}
	return output
}
