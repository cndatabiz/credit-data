package dir

import (
	"credit_data/core"
	"fmt"
	"time"
)

var LegBaseInfoDesc = "行政征收信息"

type LegBaseInfo struct {
	comp_name string
	id_shxym   string
	id_gszc   string
	id_zzjg   string
	id_swdj   string
	id_shzz   string

	fddbrxm   string
	fddbrzjlx string
	fddbrzjhm string

	djzt    string
	djjg    string
	clrq    time.Time
	zczb    float64
	zcdz    string
	hydm    string
}

var LegBaseInfoTitle = [...]string{"comp_name","id_shxym","id_gszc","id_zzjg","id_swdj","id_shzz","fddbrxm","fddbrzjlx","fddbrzjhm","djzt","djjg","clrq","zczb","zcdz","hydm"}

func (s * LegBaseInfo) LegBaseInfoValues() [] string  {
	var line = make([] string ,len(LegBaseInfoTitle))
	line[0] = s.comp_name
	line[1] = s.id_shxym
	line[2] = s.id_gszc
	line[3] = s.id_zzjg
	line[4] = s.id_swdj
	line[5] = s.id_shzz
	line[6] = s.fddbrxm
	line[7] = s.fddbrzjlx
	line[8] = s.fddbrzjhm
	line[9] = s.djzt
	line[10] = s.djjg
	line[11] = s.clrq.Format("2006-01-02")
	line[12] = fmt.Sprintf("%.2f", s.zczb)
	line[13] = s.zcdz
	line[13] = s.hydm

	return line
}

func getLegBaseRecordWithMainInfo(legInfo[3] string, natInfo[3] string) *LegBaseInfo {
	person := LegBaseInfo{}

	person.comp_name = legInfo[0]
	person.id_shxym = legInfo[2]

	person.fddbrxm  = natInfo[0]
	person.fddbrzjlx  = natInfo[1]
	person.fddbrzjhm  = natInfo[2]

	LegBaseRecord(&person)

	return &person
}

func LegBaseRecord(person *LegBaseInfo) {
	person.id_gszc = ""
	person.id_zzjg = core.GenOrgCode()
	person.id_swdj = ""
	person.id_swdj = ""

	// 行业类别字段没有值
	person.djzt = "1"
	person.djjg = "市工商局"
	person.clrq = core.RandDate("2000-01-01","2020-12-31",1)[0]
	person.zczb = core.RandFloats(10.00,1000000.00,1,2)[0]
	person.zcdz = core.GenAddress()
	person.hydm = core.GenHyCode()
}

