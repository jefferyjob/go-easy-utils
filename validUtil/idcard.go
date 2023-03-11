package validUtil

import (
	"regexp"
	"strconv"
	"time"
)

// IsIDCard 验证身份证号(18或15位)
func IsIDCard(str string) bool {
	if len(str) != 15 && len(str) != 18 {
		return false
	}
	if len(str) == 18 {
		return IsIDCard18(str)
	} else {
		return IsIDCard15(str)
	}
}

// IsIDCard18 验证18位身份证号
func IsIDCard18(id string) bool {
	// 18位身份证号码正则表达式，根据规则来编写
	regExp := "^[1-9]\\d{5}(19|20)\\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]$"
	// 利用正则表达式匹配身份证号码
	match, err := regexp.MatchString(regExp, id)
	if err != nil {
		// 匹配过程出错，返回false
		return false
	}
	if !match {
		// 身份证号码不符合规则，返回false
		return false
	}
	// 解析身份证号码中的年、月、日
	year, _ := strconv.Atoi(id[6:10])
	month, _ := strconv.Atoi(id[10:12])
	day, _ := strconv.Atoi(id[12:14])
	// 判断年份是否合法
	if year < 1900 || year > time.Now().Year() {
		return false
	}
	// 判断月份是否合法
	if month < 1 || month > 12 {
		return false
	}
	// 判断日期是否合法
	if day < 1 || day > 31 {
		return false
	}
	// 对身份证号码的最后一位进行校验
	// 根据身份证号码的规则，最后一位可能是数字0-9，也可能是字符X（表示10）
	// 将字符X转换成数字10进行校验
	lastChar := id[len(id)-1]
	var lastNum int
	if lastChar == 'X' || lastChar == 'x' {
		lastNum = 10
	} else {
		lastNum, _ = strconv.Atoi(string(lastChar))
	}
	// 对身份证号码的前17位进行加权和校验
	// 加权系数，根据规则固定
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	// 计算加权和
	sum := 0
	for i := 0; i < len(weights); i++ {
		num, _ := strconv.Atoi(string(id[i]))
		sum += num * weights[i]
	}
	// 计算校验码
	checkCode := sum % 11
	checkCodeMap := map[int]string{
		0:  "1",
		1:  "0",
		2:  "X",
		3:  "9",
		4:  "8",
		5:  "7",
		6:  "6",
		7:  "5",
		8:  "4",
		9:  "3",
		10: "2",
	}
	// 校验身份证号码的最后一位
	return checkCodeMap[checkCode] == strconv.Itoa(lastNum)
}

// IsIDCard15 验证15位身份证号
func IsIDCard15(idCard string) bool {
	// 验证是否为15位数字
	if match, _ := regexp.MatchString(`^\d{15}$`, idCard); !match {
		return false
	}

	// 将身份证号前两位转换成省份代码
	provinceCode, err := strconv.Atoi(idCard[:2])
	if err != nil || provinceCode < 11 || provinceCode > 91 {
		return false
	}

	// 验证生日是否正确
	year := strconv.Itoa(1900 + int(idCard[6]-'0')*10 + int(idCard[7]-'0'))
	month := string(idCard[8:10])
	day := string(idCard[10:12])
	if match, _ := regexp.MatchString(`^(19|20)\d{2}(0[1-9]|1[0-2])(0[1-9]|[12]\d|3[01])$`, year+month+day); !match {
		return false
	}

	// 验证校验位
	factors := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	checkCodes := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
	sum := 0
	for i := 0; i < 17; i++ {
		num, _ := strconv.Atoi(string(idCard[i]))
		sum += num * factors[i]
	}
	checkCode := checkCodes[sum%11]
	if string(idCard[14]) != checkCode {
		return false
	}

	return true
}
