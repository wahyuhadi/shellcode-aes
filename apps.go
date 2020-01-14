package main

import (
	b64 "encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	aes "shellcode-aes/aes-crypt"
)

var (
	key  = flag.String("key", "1234567812345678", "key for enc or dec shellcode")
	do   = flag.String("do", "enc", "example --do=enc or --do=dec")
	file = flag.String("file", "shellcode", "location file shellcode with hex format")
)

func main() {
	// check Key
	flag.Parse()
	keyLen := len(*key)
	if keyLen > 0 && keyLen != 16 && keyLen != 24 && keyLen != 32 {
		fmt.Println("not a valid key. lenght should be 16, 24 or 32")
		return
	}

	isKey := []byte(*key)
	// Read shellcode from file
	data, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println("[!] Error read shellcode file, Please use -file options")
		return
	}
	shellcode := []byte(data)

	if *do == "enc" {
		sEnc, err := aes.Encrypt(isKey, shellcode)
		if err != nil {
			fmt.Println(err)
		}

		b64Enc := b64.StdEncoding.EncodeToString(sEnc)
		fmt.Println("[+] Key : ", *key)
		fmt.Println("[!] Decode methode is base64(aes)")
		fmt.Println("[+] Shellcode -->", b64Enc)
	}

	if *do == "dec" {
		bs64Dec, _ := b64.StdEncoding.DecodeString(string(data))
		isAes := []byte(bs64Dec)
		sDec, err := aes.Decrypt(isKey, isAes)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("[+] success decode shellcode in file : ", *file)
		fmt.Println("[+] with Key : ", *key)
		fmt.Println("[+] Real Shellcode -->", string(sDec), "\n\n")
	}

}
