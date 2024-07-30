package cryptoUtil

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 伪造一个读取器，用来模拟产生错误的情况
type badRandomReader struct{}

func (r *badRandomReader) Read([]byte) (int, error) {
	return 0, errors.New("fake error")
}

func TestGenerateRSAKeys(t *testing.T) {
	testCases := []struct {
		name           string
		before         func(t *testing.T)
		after          func(t *testing.T)
		wantPrivateKey bool
		wantPublicKey  bool
		wantErr        error
	}{
		{
			name:           "生成成功",
			before:         func(t *testing.T) {},
			after:          func(t *testing.T) {},
			wantPrivateKey: true,
			wantPublicKey:  true,
			wantErr:        nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(t)

			privateKeyPEM, publicKeyPEM, err := GenerateRSAKeys()
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantPrivateKey, privateKeyPEM != "")
			assert.Equal(t, tc.wantPublicKey, publicKeyPEM != "")
		})
	}
}

// 测试 rsa.GenerateKey 生产失败
func TestGenerateRSAKeysError(t *testing.T) {
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

func TestGenerateRSAKeys2(t *testing.T) {
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

// 验证加密和解密后的数据是否一致
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

// 验证 DecryptRSA 方法
func TestDecryptRSA(t *testing.T) {
	privateKeyPEM, _, _ := GenerateRSAKeys()
	privateKeyBlock, _ := pem.Decode([]byte(privateKeyPEM))
	if privateKeyBlock == nil {
		t.Fatal("Failed to parse private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		t.Fatalf("Error parsing private key: %s", err)
	}

	message := []byte("Hello, World!")

	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, &privateKey.PublicKey, message)
	if err != nil {
		t.Fatalf("Error encrypting message: %s", err)
	}

	// Test decryption with correct private key
	plaintext, err := DecryptRSA(privateKeyPEM, ciphertext)
	if err != nil {
		t.Fatalf("Error decrypting ciphertext: %s", err)
	}
	if string(plaintext) != string(message) {
		t.Fatalf("Decrypted message doesn't match original message")
	}

	// Test decryption with incorrect private key
	_, err = DecryptRSA("InvalidPrivateKey", ciphertext)
	if err == nil {
		t.Fatal("Expected error decrypting with invalid private key")
	}

	// Test decryption with incorrect ciphertext
	_, err = DecryptRSA(privateKeyPEM, []byte("InvalidCiphertext"))
	if err == nil {
		t.Fatal("Expected error decrypting invalid ciphertext")
	}
}
