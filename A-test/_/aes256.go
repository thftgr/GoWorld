package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func main() {
	key := "nerianwqerj139req9rwer1949512013"
	iv := "0123456789abcdef"
	ciphertext, err := base64.StdEncoding.DecodeString("LYD086NXY5Iw32R99XgYqHdr9utfHkAbokmy31e9BOwOpK48hPcoXo59jXaWgkrq")
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, ciphertext)

	fmt.Printf("%s\n", ciphertext)
	//b, _ := base64.StdEncoding.DecodeString(string(ciphertext))
	//fmt.Printf(string(b))
}
