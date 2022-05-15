package idvalidator

import (
	"strconv"
	"strings"
	"time"
)

var weight = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var validValue = [11]byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

var validProvince = map[string]string{
	"11": "北京市",
	"12": "天津市",
	"13": "河北省",
	"14": "山西省",
	"15": "内蒙古自治区",
	"21": "辽宁省",
	"22": "吉林省",
	"23": "黑龙江省",
	"31": "上海市",
	"32": "江苏省",
	"33": "浙江省",
	"34": "安徽省",
	"35": "福建省",
	"36": "江西省",
	"37": "山东省",
	"41": "河南省",
	"42": "湖北省",
	"43": "湖南省",
	"44": "广东省",
	"45": "广西壮族自治区",
	"46": "海南省",
	"50": "重庆市",
	"51": "四川省",
	"52": "贵州省",
	"53": "云南省",
	"54": "西藏自治区",
	"61": "陕西省",
	"62": "甘肃省",
	"63": "青海省",
	"64": "宁夏回族自治区",
	"65": "新疆维吾尔自治区",
	"71": "台湾省",
	"81": "香港特别行政区",
	"91": "澳门特别行政区",
}

func Check(id string) bool {
	// 身份证位数错误
	if len(id) != 15 && len(id) != 18 {
		return false
	}

	// 转大写
	id = strings.ToUpper(id)
	switch len(id) {
	case 15:
		// 出生日期年份加‘19’
		idRune := []rune(id)
		tmp := append(idRune[:6], '1', '9')
		idRune = append(tmp, idRune[6:12]...)
		id = string(idRune)
	case 18:
		// 验证最后一位校验码
		if !checkValidNo18(id) {
			return false
		}
	}

	// 验证出生日期
	if !checkBirthdayCode(id[6:14]) {
		return false
	}

	// 验证地区(省)
	if !checkAddressCode(id[:2]) {
		return false
	}

	return true
}

// 验证18位身份证校验码
func checkValidNo18(id string) bool {
	sum := 0
	for i := 0; i < 17; i++ {
		n, _ := strconv.Atoi(string(id[i]))
		sum += n * weight[i]
	}
	// mod得出18位身份证校验码
	mod := sum % 11
	if validValue[mod] == id[17] {
		return true
	}

	return false
}

// 验证生日
func checkBirthdayCode(birthday string) bool {
	year, _ := strconv.Atoi(birthday[:4])
	month, _ := strconv.Atoi(birthday[4:6])
	day, _ := strconv.Atoi(birthday[6:])

	curYear, curMonth, curDay := time.Now().Date()
	// 出生日期大于现在的日期
	if year < 1900 || year > curYear || month <= 0 || month > 12 || day <= 0 || day > 31 {
		return false
	}

	if year == curYear {
		if month > int(curMonth) {
			return false
		} else if month == int(curMonth) && day > curDay {
			return false
		}
	}

	// 出生日期在2月份
	if month == 2 {
		// 闰年2月只有29天
		if isLeapYear(year) && day > 29 {
			return false
		} else if day > 28 { //非闰年2月只有28天
			return false
		}
	} else if 4 == month || 6 == month || 9 == month || 11 == month { //小月只有30天
		if day > 30 {
			return false
		}
	}

	return true
}

// 判断是否为闰年
func isLeapYear(year int) bool {
	if year <= 1900 {
		return false
	}
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}

// 验证地区(验证省)
func checkAddressCode(address string) bool {
	// 校验省，前两位
	if _, ok := validProvince[address]; ok {
		return true
	}

	return false
}
