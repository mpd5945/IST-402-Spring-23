package main

import (
	"crypto/aes"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"gitlab.com/elktree/ecc"
	"golang.org/x/crypto/chacha20"
	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/andreburgaud/crypt2go/padding"
)

func main() {
	// ECC encryption
	pub, priv, _ := ecc.GenerateKeys(elliptic.P384())
	plaintext := "secret secrets are no fun, secret secrets hurt someone"
	encrypted, _ := pub.Encrypt([]byte(plaintext))
	decrypted, _ := priv.Decrypt(encrypted)

	// Hashing
	hasher := sha256.New()
	hasher.Write(decrypted)
	hashed := hasher.Sum(nil)

	// ChaCha20 encryption
	nonce := make([]byte, chacha20.NonceSizeX)
	_, err := rand.Read(nonce)
	if err != nil {
		panic(err.Error())
	}
	key := make([]byte, chacha20.KeySize)
	_, err = rand.Read(key)
	if err != nil {
		panic(err.Error())
	}
	chachaCipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		panic(err.Error())
	}
	encryptedHashed := make([]byte, len(hashed))
	chachaCipher.XORKeyStream(encryptedHashed, hashed)

	// ECB encryption
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBEncrypter(block)

	// Pad encryptedHashed before encrypting with ECB
	padder := padding.NewPkcs7Padding(block.BlockSize())
	paddedEncryptedHashed, err := padder.Pad(encryptedHashed)
	if err != nil {
		panic(err.Error())
	}

	finalEncrypted := make([]byte, len(paddedEncryptedHashed))
	mode.CryptBlocks(finalEncrypted, paddedEncryptedHashed)

	// Print the final encrypted message
	finalEncryptedStr := base64.URLEncoding.EncodeToString(finalEncrypted)
	fmt.Println("Final encrypted message:", finalEncryptedStr)
}
