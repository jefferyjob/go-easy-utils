package cryptoUtil

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"testing"
)

// 伪造一个读取器，用来模拟产生错误的情况
type badRandomReader struct{}

func (r *badRandomReader) Read([]byte) (int, error) {
	return 0, errors.New("fake error")
}

func TestGenerateRSAKeys(t *testing.T) {
	// 测试 GenerateRSAKeys 是否能够正常工作
	privateKeyPEM, publicKeyPEM, err := GenerateRSAKeys()
	if err != nil {
		t.Errorf("GenerateRSAKeys() error = %v", err)
		return
	}

	// 检查返回的 PEM 字符串是否有效
	if privateKeyPEM == "" {
		t.Error("Expected non-empty private key PEM string")
	}

	if publicKeyPEM == "" {
		t.Error("Expected non-empty public key PEM string")
	}

	// 尝试解析生成的私钥
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

	// 尝试解析生成的公钥
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
	if !privateKey.PublicKey.Equal(rsaPublicKey) {
		t.Errorf("Generated private key does not match parsed public key")
		return
	}
}

func TestEncryptRSAError(t *testing.T) {
	// 测试 EncryptRSA 函数在提供无效的公钥时是否返回错误
	plaintext := []byte("Hello, RSA!")
	invalidPublicKey := "invalid_public_key"
	ciphertext, err := EncryptRSA(invalidPublicKey, plaintext)
	if err == nil {
		t.Errorf("Expected error but got ciphertext: %v", ciphertext)
	}

	// 测试 EncryptRSA 函数在解析公钥失败时是否返回错误
	invalidKeyPEM := "invalid_key_pem"
	_, err = EncryptRSA(invalidKeyPEM, plaintext)
	if err == nil {
		t.Error("Expected error but got nil")
	}
}

func TestDecryptRSAError(t *testing.T) {
	// 测试 DecryptRSA 函数在提供无效的私钥时是否返回错误
	ciphertext := []byte("fake_ciphertext")
	invalidPrivateKey := "invalid_private_key"
	plaintext, err := DecryptRSA(invalidPrivateKey, ciphertext)
	if err == nil {
		t.Errorf("Expected error but got plaintext: %v", plaintext)
	}

	// 测试 DecryptRSA 函数在解析私钥失败时是否返回错误
	invalidKeyPEM := "invalid_key_pem"
	_, err = DecryptRSA(invalidKeyPEM, ciphertext)
	if err == nil {
		t.Error("Expected error but got nil")
	}
}

func TestRSAEncryptionAndDecryption(t *testing.T) {
	// 生成 RSA 密钥对
	privateKeyPEM, publicKeyPEM, err := GenerateRSAKeys()
	if err != nil {
		t.Fatalf("Failed to generate RSA keys: %s", err)
	}

	// 加密数据
	plaintext := []byte("Hello, RSA!")
	ciphertext, err := EncryptRSA(publicKeyPEM, plaintext)
	if err != nil {
		t.Fatalf("Failed to encrypt data: %s", err)
	}

	// 解密数据
	decryptedText, err := DecryptRSA(privateKeyPEM, ciphertext)
	if err != nil {
		t.Fatalf("Failed to decrypt data: %s", err)
	}

	// 检查解密后的数据是否与原始数据一致
	if !bytes.Equal(decryptedText, plaintext) {
		t.Fatal("Decrypted text does not match the original plaintext")
	}
}
