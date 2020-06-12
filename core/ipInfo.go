package core

import (
	"credit_data/utils"
	"strconv"
)

// 生成专利信息
var ipFactor = [...]int{2,3,4,5,6,7,8,9}

func GenIpNo() string{
	var ipSn string
	ipCreate := RandDate("2000-01-01","2020-01-01",1)
	ipYear := ipCreate[0].Format("2006")
	ipSn += ipYear

	ipTypeDict := [...]string {"1","2","3","1","2","3","1","2","3","1","2","3","1","2","3","1","2","3","8","9"}
	ipType := RandDictItem(ipTypeDict[:])
	ipSn += ipType

	sn := utils.PadWithZero(7, int(RandInteger(1, 10000000, 1)[0]))
	ipSn += sn

	return ipSn + getIpCheckNo(ipSn)
}

func getIpCheckNo(ipSn string) string {
	var result int
	for i, item := range ipSn{
		if i < 8 {
			result += int(item) * ipFactor[i]
		}
	}
	result = result % 11
	if result == 10 {
		return ".X"
	}else {
		return "." + strconv.Itoa(result)
	}
}
