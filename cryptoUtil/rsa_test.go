package cryptoUtil

import (
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

func TestEncryptRSA(t *testing.T) {
	_, pubKeyStr, _ := GenerateRSAKeys()

	testCases := []struct {
		name         string
		publicKeyStr string
		message      []byte
		expectedErr  bool
	}{
		{
			name:         "正常加密",
			publicKeyStr: pubKeyStr,
			message:      []byte("test message"),
			expectedErr:  false,
		},
		{
			name:         "无效公钥",
			publicKeyStr: "invalid public key",
			message:      []byte("test message"),
			expectedErr:  true,
		},
		{
			name:         "无效公钥格式",
			publicKeyStr: "-----BEGIN PUBLIC KEY-----\nInvalidKey\n-----END PUBLIC KEY-----",
			message:      []byte("test message"),
			expectedErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := EncryptRSA(tc.publicKeyStr, tc.message)
			if tc.expectedErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

func TestDecryptRSA(t *testing.T) {
	privateKey, publicKey, _ := GenerateRSAKeys()

	// 使用公钥加密一条消息，以便测试解密
	message := []byte("test message")
	ciphertext, err := EncryptRSA(publicKey, message)
	if err != nil {
		t.Fatalf("failed to encrypt message: %v", err)
	}

	testCases := []struct {
		name          string
		privateKeyStr string
		ciphertext    []byte
		expected      []byte
		expectedErr   bool
	}{
		{
			name:          "正常解密",
			privateKeyStr: privateKey,
			ciphertext:    ciphertext,
			expected:      message,
			expectedErr:   false,
		},
		{
			name:          "无效私钥",
			privateKeyStr: "invalid private key",
			ciphertext:    ciphertext,
			expected:      nil,
			expectedErr:   true,
		},
		{
			name:          "无效私钥格式",
			privateKeyStr: "-----BEGIN RSA PRIVATE KEY-----\nInvalidKey\n-----END RSA PRIVATE KEY-----",
			ciphertext:    ciphertext,
			expected:      nil,
			expectedErr:   true,
		},
		{
			name:          "无效密文",
			privateKeyStr: privateKey,
			ciphertext:    []byte("invalid ciphertext"),
			expected:      nil,
			expectedErr:   true,
		},
		{
			name:          "无效私钥数据",
			privateKeyStr: "-----BEGIN RSA PRIVATE KEY-----\n" + "A" + "\n-----END RSA PRIVATE KEY-----",
			ciphertext:    ciphertext,
			expected:      nil,
			expectedErr:   true,
		},
		{
			name: "不完整的私钥",
			privateKeyStr: "-----BEGIN RSA PRIVATE KEY-----\n" +
				"MIICWwIBAAKBgQDEkzKS0u5p6kwl9m0g3g4mMI09S8QOAbW5aBMbDWZ5R0pUtH5h" +
				"J9mQFt8Uu4FJ8Yc9C5ZiM5F9pV5J2V4SeKk3RbKjFG2iD6rzO/OMrMZ3/1H8n02" +
				"eZ/D14SvnPBNhYnb8Ysdd4kS8A==\n-----END RSA PRIVATE KEY-----",
			ciphertext:  ciphertext,
			expected:    nil,
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := DecryptRSA(tc.privateKeyStr, tc.ciphertext)
			if tc.expectedErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}
