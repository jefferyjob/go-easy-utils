package cryptoUtil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"reflect"
	"testing"
)

// 伪造一个读取器，用来模拟产生错误的情况
type badRandomReader struct{}

func (r *badRandomReader) Read([]byte) (int, error) {
	return 0, errors.New("fake error")
}

// 测试RSA公钥和私钥生成
func TestGenerateRSAKeys(t *testing.T) {
	// 保存原始的 rand.Reader
	originalRandReader := rand.Reader

	// 替换全局的 rand.Reader 为一个模拟错误的 Reader
	rand.Reader = &badRandomReader{}

	// 恢复原始的 rand.Reader
	defer func() {
		rand.Reader = originalRandReader
	}()

	privateKeyPEM, publicKeyPEM, err := GenerateRSAKeys()

	// 检查返回的错误是否符合预期
	expectedErr := errors.New("fake error")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Expected error '%v' but got '%v'", expectedErr, err)
	}

	// 检查返回的 PEM 字符串是否为空
	if privateKeyPEM != "" {
		t.Error("Expected empty private key PEM string")
	}

	if publicKeyPEM != "" {
		t.Error("Expected empty public key PEM string")
	}
}

// 测试 RSA 加密和解密功能
func TestRSAEncryptionAndDecryption(t *testing.T) {
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

	// 加密数据
	plaintext := []byte("Hello, RSA!")
	ciphertext, err := EncryptRSA(publicKeyPEM, plaintext)
	if err != nil {
		t.Errorf("Failed to encrypt data: %s", err)
		return
	}

	// 解密数据
	decryptedText, err := DecryptRSA(privateKeyPEM, ciphertext)
	if err != nil {
		t.Errorf("Failed to decrypt data: %s", err)
		return
	}

	// 检查解密后的数据是否与原始数据一致
	if string(decryptedText) != string(plaintext) {
		t.Errorf("Decrypted text does not match the original plaintext")
	}
}
