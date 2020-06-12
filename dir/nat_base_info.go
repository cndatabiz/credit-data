package dir

import (
	"credit_data/core"
	"time"
)

var NatBaseInfoDesc = "自然人基本信息"

type NatBaseInfo struct {
	name    string
	id_type string
	id_code string
	xb      string
	csrq    time.Time
	gjdq    string
	txdz    string
	qfrq    time.Time
	qfjg    string
	djzt    string
}

var NatBaseInfoTitle = [...]string{"name","id_type","id_code","xb","csrq","gjdq","txdz","qfrq","qfjg","djzt"}


func (s * NatBaseInfo) NatBaseInfoValues() [] string  {
	var line = make([] string ,len(NatBaseInfoTitle))

	line[0] = s.name
	line[1] = s.id_type
	line[2] = s.id_code
	line[3] = s.xb
	line[4] = s.csrq.Format("2006-01-02")
	line[5] = s.gjdq
	line[6] = s.txdz
	line[7] = s.qfrq.Format("2006-01-02")
	line[8] = s.qfjg
	line[9] = s.djzt

	return line
}

func getNatBaseInfoRecordWithMainInfo(netInfo[3] string, xb, birthday string ) *NatBaseInfo {
	person := NatBaseInfo{}

	person.name = netInfo[0]
	person.id_type = netInfo[1]
	person.id_code = netInfo[2]

	if xb == "男" {
		person.xb = "1"
	} else {
		person.xb = "2"
	}
	person.csrq, _ = time.Parse("2006-01-02",birthday)
	natBaseInfo(&person)

	return &person
}

func natBaseInfo(person *NatBaseInfo) {
	person.gjdq = "156"
	person.txdz = core.GenAddress()
	person.qfrq = core.RandDate("2005-01-01","2018-12-31",1)[0]
	person.qfjg = "市公安局"
	person.djzt = "1"
}