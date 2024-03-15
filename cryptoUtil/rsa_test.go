package cryptoUtil

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"reflect"
	"testing"
)

func TestGenerateRSAKeys(t *testing.T) {
	privateKeyPEM, publicKeyPEM, err := GenerateRSAKeys()
	if err != nil {
		t.Errorf("GenerateRSAKeys() error = %v", err)
		return
	}

	// 解码私钥 PEM 格式
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		t.Errorf("Invalid private key PEM")
		return
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		t.Errorf("Failed to parse private key: %v", err)
		return
	}

	// 解码公钥 PEM 格式
	block, _ = pem.Decode([]byte(publicKeyPEM))
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		t.Errorf("Invalid public key PEM")
		return
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		t.Errorf("Failed to parse public key: %v", err)
		return
	}
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		t.Errorf("Failed to convert public key to RSA public key")
		return
	}

	// 比较生成的密钥与解析的密钥
	if !reflect.DeepEqual(privateKey.Public(), rsaPublicKey) {
		t.Errorf("Generated private key does not match parsed public key")
		return
	}
}

func TestRSAEncryptionAndDecryption(t *testing.T) {
	// 生成RSA密钥对
	privateKey, publicKey, err := GenerateRSAKeys()
	if err != nil {
		t.Errorf("Failed to generate RSA keys: %s", err)
		return
	}

	// 加密数据
	plaintext := []byte("Hello, RSA!")
	ciphertext, err := EncryptRSA(publicKey, plaintext)
	if err != nil {
		t.Errorf("Failed to encrypt data: %s", err)
		return
	}

	// 解密数据
	decryptedText, err := DecryptRSA(privateKey, ciphertext)
	if err != nil {
		t.Errorf("Failed to decrypt data: %s", err)
		return
	}

	// 检查解密后的数据是否与原始数据一致
	if string(decryptedText) != string(plaintext) {
		t.Errorf("Decrypted text does not match the original plaintext")
	}
}
