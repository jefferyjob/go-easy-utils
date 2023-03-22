package validUtil

// IsBankCardNo 验证是否为大陆银行卡号
func IsBankCardNo(cardNumber string) bool {
	if len(cardNumber) != 16 && len(cardNumber) != 19 {
		return false
	}
	var cardArr []int
	for _, c := range cardNumber {
		if c < '0' || c > '9' {
			return false
		}
		cardArr = append(cardArr, int(c-'0'))
	}
	if len(cardArr) == 16 {
		sum := 0
		for i := len(cardArr) - 1; i >= 0; i-- {
			if i%2 == 0 {
				cardArr[i] *= 2
				if cardArr[i] > 9 {
					cardArr[i] -= 9
				}
			}
			sum += cardArr[i]
		}
		return sum%10 == 0
	} else {
		sum := 0
		for i := len(cardArr) - 1; i >= 0; i-- {
			if (len(cardArr)-i)%2 == 0 {
				cardArr[i] *= 2
				if cardArr[i] > 9 {
					cardArr[i] -= 9
				}
			}
			sum += cardArr[i]
		}
		return sum%10 == 0
	}
}
