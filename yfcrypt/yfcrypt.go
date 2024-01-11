package yfcrypt

import (
	"fmt"
	"os"

	"github.com/wenzhenxi/gorsa"
)

func EncryptStringWithPrivateKey(privateKeyPath string, plaintext string) (string, error) {
	keyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return "", fmt.Errorf("cannot read key file%v", err)
	}

	cipherText, err := gorsa.PriKeyEncrypt(plaintext, string(keyBytes))

	if err != nil {
		return "", err
	}

	return cipherText, nil
}

func EncryptFileWithPrivateKey(privateKeyPath string, cmdFile string) (string, error) {
	keyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return "", fmt.Errorf("cannot read key file%v", err)
	}

	fileContent, err := os.ReadFile(cmdFile)
	if err != nil {
		return "", fmt.Errorf("cannot read key file%v", err)
	}

	cipherText, err := gorsa.PriKeyEncrypt(string(fileContent), string(keyBytes))
	if err != nil {
		return "", err
	}
	return cipherText, nil
}

func DecryptStringWithPublicKey(publicKeyPath string, cipherText string, publicKey string) (string, error) {
	pub_key := publicKey
	if _, err := os.Stat(publicKeyPath); err == nil || os.IsExist(err) {
		keyBytes, err := os.ReadFile(publicKeyPath)
		if err == nil {
			pub_key = string(keyBytes)
		}

	}

	plaintext, err := gorsa.PublicDecrypt(cipherText, pub_key)
	if err != nil {
		return "", err
	}

	return plaintext, nil
}

func DecryptFileWithPublicKey(PublicKeyPath string, cipherFile string, publicKey string) (string, error) {
	pub_key := publicKey
	if _, err := os.Stat(PublicKeyPath); err == nil || os.IsExist(err) {
		keyBytes, err := os.ReadFile(PublicKeyPath)
		if err == nil {
			pub_key = string(keyBytes)
		}
	}

	fileContent, err := os.ReadFile(cipherFile)
	if err != nil {
		return "", fmt.Errorf("cannot read key file%v", err)
	}

	plaintext, err := gorsa.PublicDecrypt(string(fileContent), pub_key)
	if err != nil {
		return "", err
	}

	return plaintext, nil
}
