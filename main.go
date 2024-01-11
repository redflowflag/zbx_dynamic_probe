package main

import (
	"cmder"
	"fmt"
	"os"
	"path/filepath"
	"yfcrypt"
)

var pub_key = ``

var helpStr = `
Usage: yfprobe [OPTION]... [FILE]...
examples:
yfprobe cmd <encrypted cmd string>            Run Command
yfprobe lua <encrypted lua string>            Run Lua Script
yfprobe encrypt <private key> <filename>      Encrypt Command File
yfprobe decrypt <public key> <filename>       Decrypt Command File
`

func main() {
	lenArgs := len(os.Args)
	if lenArgs < 3 {
		fmt.Println("err: arguments invalid")
		fmt.Println(helpStr)
		return
	}

	if os.Args[1] == "encrypt" {
		cipher, err := yfcrypt.EncryptFileWithPrivateKey(os.Args[2], os.Args[3])
		if err != nil {
			fmt.Println("err: encrypt error", err)
		}
		fmt.Println(cipher)
		return
	}
	exePath, _ := os.Executable()
	curPath := filepath.Dir(exePath)
	if os.Args[1] == "decrypt" {
		plaintext, err := yfcrypt.DecryptStringWithPublicKey(curPath+"./rsa_pub.pem", os.Args[3], pub_key)
		if err != nil {
			fmt.Println("err: decrypt error", err)
		}
		fmt.Println(plaintext)
		return
	}

	//decrypt args3 [command/script]
	plaintext, err1 := yfcrypt.DecryptStringWithPublicKey(curPath+"./rsa_pub.pem", os.Args[2], pub_key)
	if err1 != nil {
		fmt.Println("err: decrypt error", err1)
		return
	}

	if os.Args[1] == "cmd" {

		rlt, err := cmder.RunCommand(plaintext)
		if err != nil {
			fmt.Println("err: run cmd error", err)
			return
		}
		fmt.Println(rlt)
		return
	}
	if os.Args[1] == "lua" {
		cmder.RunLuaScript(plaintext)
		return
	}
}
