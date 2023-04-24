package main

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "encoding/hex"
    "fmt"
)

func main() {
    // Generate a private key
    privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Get the public key
    publicKey := &privateKey.PublicKey

    // Get the user input
    userInput := "Hello, World!"

    // Convert the user input to bytes
    inputBytes := []byte(userInput)

    // Encrypt the user input
    encryptedBytes, sharedSecret, err := encrypt(inputBytes, publicKey)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Convert the encrypted bytes to a hex string
    encryptedHex := hex.EncodeToString(encryptedBytes)

    fmt.Println("Encrypted:", encryptedHex)

    // Decrypt the encrypted bytes
    decryptedBytes, err := decrypt(encryptedBytes, privateKey, sharedSecret)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Convert the decrypted bytes to a string
    decryptedString := string(decryptedBytes)

    fmt.Println("Decrypted:", decryptedString)
}

func encrypt(input []byte, publicKey *ecdsa.PublicKey) ([]byte, []byte, error) {
    // Generate a random number
    k, err := rand.Int(rand.Reader, publicKey.Params().N)
    if err != nil {
        return nil, nil, err
    }

    // Calculate the x and y coordinates of the point k*G
    x, y := publicKey.Curve.ScalarBaseMult(k.Bytes())

    // Calculate the shared secret
    sharedSecretX, _ := publicKey.Curve.ScalarMult(x, y, k.Bytes())
    sharedSecret := sharedSecretX.Bytes()

    // Generate a random number for the nonce
    nonce, err := rand.Int(rand.Reader, publicKey.Params().N)
    if err != nil {
        return nil, nil, err
    }

    // Calculate the x and y coordinates of the point nonce*G
    x, y = publicKey.Curve.ScalarBaseMult(nonce.Bytes())

    // Calculate the ciphertext
    ciphertextX, ciphertextY := publicKey.Curve.ScalarMult(x, y, publicKey.Params().Gx, publicKey.Params().Gy)
    x, y = publicKey.Curve.ScalarMult(x, y, sharedSecret)
    ciphertextX, ciphertextY = publicKey.Curve.Add(ciphertextX, ciphertextY, x, y)
    ciphertext := make([]byte, 2*publicKey.Params().BitSize/8)
    copy(ciphertext[:], ciphertextX.Bytes())
    copy(ciphertext[publicKey.Params().BitSize/8:], ciphertextY.Bytes())

    // Calculate the tag
    tagX, tagY := publicKey.Curve.ScalarBaseMult(nonce.Bytes())
    tagX, tagY = publicKey.Curve.ScalarMult(tagX, tagY, sharedSecret)
    tag := make([]byte, 2*publicKey.Params().BitSize/8)
    copy(tag[:], tagX.Bytes())
    copy(tag[publicKey.Params().BitSize/8:], tagY.Bytes())

    // Concatenate the ciphertext and tag
    output := make([]byte, len(ciphertext)+len(tag))
    copy(output, ciphertext)
    copy(output[len(ciphertext):], tag)

    return output, sharedSecret, nil
}

func decrypt(input []byte, privateKey *ecdsa.PrivateKey, sharedSecret []byte) ([]byte, error) {
    // Split the input into the ciphertext and tag
    ciphertext := input[:len(input)/2]
    tag := input[len(input)/2:]

    // Calculate the x and y coordinates of the point nonce*G
    x, y := privateKey.PublicKey.Curve.ScalarBaseMult(tag[:privateKey.PublicKey.Params().BitSize/8])
    x2, y2 := privateKey.PublicKey.Curve.ScalarMult(tag[privateKey.PublicKey.Params().BitSize/8:], x, privateKey.D)
    x, y = privateKey.PublicKey.Curve.Add(x, y, x2, y2)

    // Calculate the x and y coordinates of the point ciphertext*sharedSecret
    x, y = privateKey.PublicKey.Curve.ScalarMult(ciphertext[:privateKey.PublicKey.Params().BitSize/8], ciphertext[privateKey.PublicKey.Params().BitSize/8:], sharedSecret)

    // Calculate the plaintext
    plaintext := make([]byte, len(ciphertext)/2)
    copy(plaintext, x.Bytes())

    return plaintext, nil
}