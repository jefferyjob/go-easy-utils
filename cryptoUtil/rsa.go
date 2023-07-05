package cryptoUtil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// GenerateRSAKeys 生成RSA私钥和公钥
//
//	func GenerateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
//		// 生成私钥
//		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
//		if err != nil {
//			return nil, nil, err
//		}
//
//		// 从私钥中提取公钥
//		publicKey := &privateKey.PublicKey
//
//		return privateKey, publicKey, nil
//	}
func GenerateRSAKeys() (string, string, error) {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}

	// 从私钥中提取公钥
	publicKey := &privateKey.PublicKey

	//return privateKey, publicKey, nil

	// 将私钥转换为PEM格式
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	// 将公钥转换为PEM格式
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	return string(privateKeyPEM), string(publicKeyPEM), nil
}

// EncryptRSA RSA加密数据
//
//	func EncryptRSA(publicKey *rsa.PublicKey, plaintext []byte) ([]byte, error) {
//		ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)
//		if err != nil {
//			return nil, err
//		}
//
//		return ciphertext, nil
//	}
func EncryptRSA(publicKeyStr string, message []byte) ([]byte, error) {
	// 解码公钥
	publicKeyBlock, _ := pem.Decode([]byte(publicKeyStr))
	if publicKeyBlock == nil {
		return nil, fmt.Errorf("failed to parse public key")
	}

	// 解析公钥
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}

	// 类型断言为RSA公钥
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to convert public key to RSA public key")
	}

	// 加密消息
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, message)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}

// DecryptRSA RSA解密数据
//
//	func DecryptRSA(privateKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
//		plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
//		if err != nil {
//			return nil, err
//		}
//
//		return plaintext, nil
//	}
func DecryptRSA(privateKeyStr string, ciphertext []byte) ([]byte, error) {
	// 解码私钥
	privateKeyBlock, _ := pem.Decode([]byte(privateKeyStr))
	if privateKeyBlock == nil {
		return nil, fmt.Errorf("failed to parse private key")
	}

	// 解析私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}

	// 解密密文
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
