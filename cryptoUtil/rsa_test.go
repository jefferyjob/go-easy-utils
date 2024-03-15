package cryptoUtil

import (
	"testing"
)

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
