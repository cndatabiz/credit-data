package dir

import (
	"credit_data/core"
	"credit_data/utils"
	"fmt"
	"time"
)

var TdDesc = "评价信息"

type LegTd struct {
	comp_name string
	id_type   string
	id_code   string
	tdzh      string
	zld       string
	djh       string
	syqlx     string
	symj      float64
	yxqssj    time.Time
	yxzzsj    time.Time
	sfdy      string
	fzjg      string
	fzrq      time.Time
}

var TdTitle = [...]string{"comp_name","id_type","id_code","tdzh","zld","djh","syqlx","symj","yxqssj","yxzzsj","sfdy","fzjg","fzrq"}


func (s * LegTd) tdValues() [] string  {
	var line = make([] string ,len(TdTitle))
	line[0] = s.comp_name
	line[1] = s.id_type
	line[2] = s.id_code
	line[3] = s.tdzh
	line[4] = s.zld
	line[5] = s.djh
	line[6] = s.syqlx
	line[7] = fmt.Sprintf("%.2f", s.symj)
	line[8] = s.yxqssj.Format("2006-01-02")
	line[9] = s.yxzzsj.Format("2006-01-02")
	line[10] = s.sfdy
	line[11] = s.fzjg
	line[12] = s.fzrq.Format("2006-01-02")

	return line
}

func getTdRecordWithMainInfo(legInfo[3] string) *LegTd {
	person := LegTd{}

	person.comp_name = legInfo[0]
	person.id_type = legInfo[1]
	person.id_code = legInfo[2]

	tdInfo(&person)

	return &person
}

func tdInfo(person *LegTd) {
	tdYear := fmt.Sprintf("%d",core.RandInteger(1990,2020,1)[0])
	djh    := utils.PadWithZero(7, int(core.RandInteger(1, 10000, 1)[0]))
	person.tdzh = tdYear + djh

	person.zld = core.GenAddress()
	person.djh = djh

	var syqlx = [...]string{"0", "1", "9"}
	person.syqlx = core.RandDictItem(syqlx[:])

	person.symj = core.RandFloats(100.00,100000.00,2,1)[0]
	person.yxqssj = core.RandDate("2005-01-01","2009-12-31",1)[0]
	person.yxzzsj = core.RandDate("2010-01-01","2020-12-31",1)[0]

	var sfdy = [...]string{"0", "1"}
	person.sfdy = core.RandDictItem(sfdy[:])
	person.fzjg = "市国土局"

	person.fzrq = core.RandDate("2005-01-01","2009-12-31",1)[0]
}