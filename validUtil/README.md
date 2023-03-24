# ValidUtil

## Install

```bash
go get -u github.com/jefferyjob/go-easy-utils/validUtil
```

## Functions

```go
// IsTime 验证是否为时间格式（HH:mm:ss）
func IsTime(str string) bool

// IsDate 验证是否为日期格式（yyyy-MM-dd）
func IsDate(str string) bool

// IsDateTime 验证是否为日期时间格式（yyyy-MM-dd HH:mm:ss）
func IsDateTime(str string) bool

// IsIDCard 验证身份证号(18或15位)
func IsIDCard(str string) bool

// IsIDCard18 验证18位身份证号
func IsIDCard18(id string) bool

// IsIDCard15 验证15位身份证号
func IsIDCard15(idCard string) bool

// IsMobile 验证是否为手机号码
func IsMobile(mobileNum string) bool 

// IsTelephone 验证是否为座机号码
func IsTelephone(telephone string) bool 

// IsPostalCode 验证是否为邮编号码
func IsPostalCode(str string) bool 

// IsDecimal 验证给定的字符串小数点后是否最多两位
func IsDecimal(input string) bool 

// IsNumber 验证是否全部为数字
func IsNumber(input string) bool

// IsBankCardNo 验证是否为银行卡号
func IsBankCardNo(str string) bool

// IsAllChinese 验证给定的字符串全部为中文
func IsAllChinese(input string) bool

// IsContainChinese 验证给定的字符串包含中文
func IsContainChinese(input string) bool

// IsEmail 是否为email
func IsEmail(input string) bool

// IsIPv4 是否为ipv4地址
func IsIPv4(input string) bool

// IsIPv6 是否为ipv6地址
func IsIPv6(input string) bool

// IsURL 是否为URL地址
func IsURL(input string) bool

// IsJSON 是否为Json
func IsJSON(input string) bool

// IsChineseName 验证是否为中文名
func IsChineseName(name string) bool

// IsEnglishName 验证是否为英文名
func IsEnglishName(name string) bool

// IsQQ 验证是否为QQ号
func IsQQ(qq string) bool 

// IsWeChat 验证是否为微信号
func IsWeChat(wechat string) bool

// IsWeibo 验证是否为微博ID
func IsWeibo(weibo string) bool

// IsPassword 验证密码是否合法
// 密码长度在6-20个字符之间，只包含数字、字母和下划线
func IsPassword(password string) bool
```