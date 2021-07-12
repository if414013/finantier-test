package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encryption-service/exception"
	"encryption-service/model"
	"fmt"
	"io"
	"os"
)

func NewEncryptionService() EncryprionService {
	return &encryptionServiceImpl{}
}

type encryptionServiceImpl struct{}

func (service *encryptionServiceImpl) EncryptStockInfo(stockInfo model.StockInfo) (response *model.StockInfo) {
	// assuming that we only need to encrypt the field value
	response = &model.StockInfo{
		Symbol:    encrypt(stockInfo.Symbol),
		Timestamp: encrypt(stockInfo.Timestamp),
		Open:      encrypt(stockInfo.Open),
		High:      encrypt(stockInfo.High),
		Low:       encrypt(stockInfo.Low),
		Close:     encrypt(stockInfo.Close),
		Volume:    encrypt(stockInfo.Volume),
	}
	return response
}

func encrypt(stringToEncrypt string) (encryptedString string) {

	keyString := os.Getenv("ENCRYPTION_KEY")

	//Since the key is in string, we need to convert decode it to bytes
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		exception.PanicIfNeeded(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		exception.PanicIfNeeded(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		exception.PanicIfNeeded(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}
