package dir

import (
	"credit_data/core"
	"time"
)

var NatTouSuDesc = "投诉信息"

type NatTousu struct {
	name      string
	id_type   string
	id_code   string
	tsrq      time.Time
	tssy      string
	tscljg    string
	tsresult  string
	tsclrq    time.Time
	bz        string
}
var NatTouSuTitle = [...]string{"name","id_type","id_code","tsrq","tssy","tscljg","tsresult","tsclrq","bz"}

func (s * NatTousu) NatTsValues() [] string  {
	var line = make([] string ,len(NatTouSuTitle))

	line[0] = s.name
	line[1] = s.id_type
	line[2] = s.id_code
	line[3] = s.tsrq.Format("2006-01-02")
	line[4] = s.tssy
	line[5] = s.tscljg
	line[6] = s.tsresult
	line[7] = s.tsclrq.Format("2006-01-02")
	line[8] = s.bz

	return line
}

func getNatTouSuRecordWithMainInfo(netInfo[3] string) *NatTousu {
	person := NatTousu{}

	person.name = netInfo[0]
	person.id_type = netInfo[1]
	person.id_code = netInfo[2]

	natTouSuInfo(&person)

	return &person
}

func natTouSuInfo(person *NatTousu) {
	person.tsrq = core.RandDate("2018-01-01", "2019-12-31", 1)[0]

	var tssy = [...]string{"破坏绿化", "服务态度","恶意竞争","违规竞标","行贿投标","虚假广告","虚假宣传","拖欠工资"}
	person.tssy = core.RandDictItem(tssy[:])

	var tscljg = [...]string{"环保局", "行政中心","工商局","发改委","社会保障局"}
	person.tscljg = core.RandDictItem(tscljg[:])

	var tsresult = [...]string{"已受理", "已受理并按属实处理","处理教育被投诉方","恶意投诉","已转相关部门"}
	person.tsresult = core.RandDictItem(tsresult[:])

	person.tsclrq = core.RandDate("2019-01-01", "2020-07-31", 1)[0]
}