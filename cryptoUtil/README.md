# cryptoUtil

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/cryptoUtil
```

## Import

```go
import (
	"github.com/jefferyjob/go-easy-utils/v2/cryptoUtil"
)
```

## Functions

```go
// HashSHA256 hash加密
func HashSHA256(str string) string

// Md5 MD5加密
func Md5(str string) string

// GenerateRSAKeys 生成RSA私钥和公钥
func GenerateRSAKeys() (string, string, error)

// EncryptRSA RSA加密数据
func EncryptRSA(publicKeyStr string, message []byte) ([]byte, error)

// DecryptRSA RSA解密数据
func DecryptRSA(privateKeyStr string, ciphertext []byte) ([]byte, error)
```