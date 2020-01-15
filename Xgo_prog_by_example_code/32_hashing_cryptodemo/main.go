package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {

	demoHash_md5()
	demoHash_sha256()
	demoHashKey("mykey", "Hello world, go!")

	// symmetric crypto
	key := "this is key 1234"
	message := "Hello world, go!"
	encrypted := encrypt_symmetric_crpyto(key, message)
	fmt.Println("message:")
	fmt.Println(message)
	fmt.Println("key:")
	fmt.Println(key)
	fmt.Println("encrypted:")
	fmt.Println(encrypted)

	decrypted := decrypt_symmetric_crpyto(key, encrypted)
	fmt.Println("encrypted message:")
	fmt.Println(encrypted)
	fmt.Println("key:")
	fmt.Println(key)
	fmt.Println("decrypted:")
	fmt.Println(decrypted)

	// asymmetric crypto
	generateRSAkeys()
	plainText := "Hello world, go!"
	fmt.Println("plainText:")
	fmt.Println(plainText)
	rsa_encrypted := encrypt_asymmetric_crypto(plainText)
	fmt.Println("encrypted:")
	fmt.Println(rsa_encrypted)
	rsa_decrypted := decrypt_asymmetric_crypto(rsa_encrypted)
	fmt.Println("decrypted:")
	fmt.Println(rsa_decrypted)
}

func demoHash_md5() {

	fmt.Println("--------Demo encoding hash using md5--------")

	message := "Hello world, go!"
	fmt.Println("plaintext:")
	fmt.Println(message)

	h := md5.New()
	h.Write([]byte(message))
	hash_message := hex.EncodeToString(h.Sum(nil))
	fmt.Println("hashing message:")
	fmt.Println(hash_message)

}
func demoHash_sha256() {

	fmt.Println("--------Demo encoding hash using sha256--------")

	message := "Hello world, go!"
	fmt.Println("plaintext:")
	fmt.Println(message)

	h := sha256.New()
	h.Write([]byte(message))
	hash_message := hex.EncodeToString(h.Sum(nil))
	fmt.Println("hashing message:")
	fmt.Println(hash_message)

}
func demoHashKey(key, message string) {

	fmt.Println("--------Demo encoding hash with key: HMAC and sha256--------")

	fmt.Println("key:")
	fmt.Println(key)
	fmt.Println("plaintext:")
	fmt.Println(message)

	hmacKey := []byte(key)
	h := hmac.New(sha256.New, hmacKey)
	h.Write([]byte(message))
	hash_message := hex.EncodeToString(h.Sum(nil))
	fmt.Println("hashing message:")
	fmt.Println(hash_message)

}

func encrypt_symmetric_crpyto(key, message string) string {

	fmt.Println("--------Demo encrypt encrypt_symmetric_crpyto--------")
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		fmt.Println("key must 16,24,32 byte length")
		return ""
	}
	bc, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	text := []byte(message)

	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(bc, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)

	return base64.StdEncoding.EncodeToString(ciphertext)
}

func decrypt_symmetric_crpyto(key, message string) string {

	fmt.Println("--------Demo decrypt decrypt_symmetric_crpyto--------")
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		fmt.Println("key must 16,24,32 byte length")
		return ""
	}
	encrypted, _ := base64.StdEncoding.DecodeString(message)

	bc, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(bc, iv)
	cfb.XORKeyStream(encrypted, encrypted)

	return string(encrypted)
}

func generateRSAkeys() {
	fmt.Println("Generating RSA keys....")

	// cambiar archivos y sus rutas
	privKeyFile := "./private.rsa.key"
	pubKeyFile := "./public.rsa.key"

	// generar RSA keys
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}

	// extraer claves privadas y públicas de RSA keys
	privASN1 := x509.MarshalPKCS1PrivateKey(privateKey)
	pubASN1, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}

	// almacenar claves privadas y públicas en archivos
	privBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privASN1,
	})

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	ioutil.WriteFile(privKeyFile, privBytes, 0644)
	ioutil.WriteFile(pubKeyFile, pubBytes, 0644)
	fmt.Println("Done")

}

func encrypt_asymmetric_crypto(message string) string {

	fmt.Println("--------Demo encrypt encrypt_asymmetric_crypto--------")

	// public key file
	pubKeyFile := "./public.rsa.key"

	// leer la clave pública del archivo
	pubBytes, err := ioutil.ReadFile(pubKeyFile)
	if err != nil {
		panic(err)
	}

	pubBlock, _ := pem.Decode(pubBytes)
	if pubBlock == nil {
		fmt.Println("Failed to load public key file")
		return ""
	}

	// Decode RSA public key
	publicKey, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		fmt.Printf("bad public key: %s", err)
		return ""
	}

	// encrypt message
	msg := []byte(message)
	encryptedmsg, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), msg)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(encryptedmsg)
}
func decrypt_asymmetric_crypto(message string) string {

	fmt.Println("--------Demo decrypt decrypt_asymmetric_crypto--------")

	// private key file
	privKeyFile := "./private.rsa.key"

	// leer la clave privada del archivo
	privBytes, err := ioutil.ReadFile(privKeyFile)
	if err != nil {
		panic(err)
	}

	privBlock, _ := pem.Decode(privBytes)
	if privBlock == nil {
		fmt.Println("Failed to load private key file")
		return ""
	}

	// Decode RSA private key
	privateKey, err := x509.ParsePKCS1PrivateKey(privBlock.Bytes)
	if err != nil {
		fmt.Printf("bad public key: %s", err)
		return ""
	}

	// encrypt message
	encrypted, _ := base64.StdEncoding.DecodeString(message)
	decrypteddmsg, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encrypted)
	if err != nil {
		panic(err)
	}

	return string(decrypteddmsg)
}
