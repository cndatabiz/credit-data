package dir

import (
	core "credit_data/core"
	"time"
)

var TsDesc = "法人投诉信息"

type LegTousu struct {
	comp_name string
	id_type   string
	id_code   string
	fddbrxm   string
	fddbrzjlx string
	fddbrzjhm string
	tsrq      time.Time
	tssy      string
	tscljg    string
	tsresult  string
	tsclrq    time.Time
	bz        string
}
var tsTitle = [...]string{"comp_name","id_type","id_code","fddbrxm","fddbrzjlx","fddbrzjhm","tsrq","tssy","tscljg","tsresult","tsclrq","bz"}

func (s * LegTousu) tsValues() [] string  {
	var line = make([] string ,len(tsTitle))

	line[0] = s.comp_name
	line[1] = s.id_type
	line[2] = s.id_code
	line[3] = s.fddbrxm
	line[4] = s.fddbrzjlx
	line[5] = s.fddbrzjhm
	line[6] = s.tsrq.Format("2006-01-02")
	line[7] = s.tssy
	line[8] = s.tscljg
	line[9] = s.tsresult
	line[10] = s.tsclrq.Format("2006-01-02")
	line[11] = s.bz

	return line
}

func getTouSuRecordWithMainInfo(legInfo[3] string, netInfo[3] string) *LegTousu {
	person := LegTousu{}

	person.comp_name = legInfo[0]
	person.id_type = legInfo[1]
	person.id_code = legInfo[2]

	person.fddbrxm = netInfo[0]
	person.fddbrzjlx = netInfo[1]
	person.fddbrzjhm = netInfo[2]
	tsInfo(&person)

	return &person
}

func tsInfo(person *LegTousu) {
	person.tsrq = core.RandDate("2018-01-01", "2019-12-31", 1)[0]

	var tssy = [...]string{"破坏绿化", "服务态度","恶意竞争","违规竞标","行贿投标","虚假广告","虚假宣传","拖欠工资"}
	person.tssy = core.RandDictItem(tssy[:])

	var tscljg = [...]string{"环保局", "行政中心","工商局","发改委","社会保障局"}
	person.tscljg = core.RandDictItem(tscljg[:])

	var tsresult = [...]string{"已受理", "已受理并按属实处理","处理教育被投诉方","恶意投诉","已转相关部门"}
	person.tsresult = core.RandDictItem(tsresult[:])

	person.tsclrq = core.RandDate("2019-01-01", "2020-07-31", 1)[0]
}