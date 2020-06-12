package core

import (
	"math"
	"math/rand"
	"regexp"
	"time"
)

/**
 * 生成随机小数
 * n 生成数据个数
 * precision 小数精度
 */
func RandFloats(min, max float64, n int, precision int) []float64 {
	if precision < 1 {
		precision = 1
	}

	pData := math.Pow10(precision)

	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64() * (max - min)
		res[i] = math.Round(res[i]*pData)/pData
	}
	return res
}

/**
 * 生成随机整数
 * n 生成数据个数
 */
func RandInteger(min, max int64,n int) [] int64 {
	res := make([] int64, n)
	for i := range res{
		res[i] = min + rand.Int63n(max-min)
	}
	return res
}

/**
 * 在指定范围内随机生成日期
 * start,end 输入格式 2006-01-02
 * n 生成日期个数
 */
var re = regexp.MustCompile("((19|20)\\d\\d)-(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])")

func RandDate(start, end string, n int) [] time.Time {
	res := make([] time.Time, n)
	if !re.MatchString(start) || !re.MatchString(end){
		panic("日期格式不正确,需输入yyyy-MM-dd格式")
	}

	sDate,_ := time.Parse("2006-01-02", start)
	eDate,_ := time.Parse("2006-01-02", end)

	max := eDate.Unix()
	min := sDate.Unix()

	for i := range res{
		sec := rand.Int63n(max - min) + min
		res[i] = time.Unix(sec,0)
	}

	return res
}

var numberRunes = []rune("1234567890")
var lowercaseRunes = []rune("abcdefghijklmnopqrstuvwxyz")
var capitalRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var letterRunes  = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var charRunes    = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
/**
 * 在英文字母
 * charType 是否为大写，true为大写，false小写
 * n 生成日期个数
 */
func RandLetterOrNum(charType CharType, n int) string {
	res := make([]rune, n)
	switch charType {
	case Number:
		for i := range res {
			res[i] = numberRunes[rand.Intn(len(numberRunes))]
		}
	case Lowercase:
		for i := range res {
			res[i] = lowercaseRunes[rand.Intn(len(lowercaseRunes))]
		}
	case Capital:
		for i := range res {
			res[i] = capitalRunes[rand.Intn(len(capitalRunes))]
		}
	case Letter:
		for i := range res {
			res[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
	case Character:
		for i := range res {
			res[i] = charRunes[rand.Intn(len(charRunes))]
		}
	}

	return string(res)
}

func RandDictItem(dictItems []string) string {
	return dictItems[rand.Intn(len(dictItems))]
}

