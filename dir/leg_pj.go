package dir

import (
	"credit_data/core"
	"credit_data/utils"
	"strconv"
	"time"
)

var PjDesc = "评价信息"

type LegPj struct {
	comp_name string
	id_type   string
	id_code   string
	pjnr      string
	pjxz      string
	pjjgmc    string
	pjrq      time.Time
	bz        string
}

var pjTitle = [...]string{"comp_name","id_type","id_code","pjnr","pjxz","pjjgmc","pjrq","bz"}


func (s * LegPj) pjValues() [] string  {
	var line = make([] string ,len(pjTitle))
	line[0] = s.comp_name
	line[1] = s.id_type
	line[2] = s.id_code
	line[3] = s.pjnr
	line[4] = s.pjxz
	line[5] = s.pjjgmc
	line[6] = s.pjrq.Format("2006-01-02")
	line[7] = s.bz

	return line
}

func GenLegPjData(rows int)  {
	file,writer := utils.FileWriter("legPj.csv")
	defer file.Close()
	defer writer.Flush()

	var err = writer.Write(pjTitle[:])
	utils.CheckError(PjDesc,"Cannot write to file", err)

	for i := 0; i < rows; i++ {
		var legInfo = [3]string{
			core.GenLegName("徐州市"),"L1",core.GenUnifyCode("91" + "320300"),
		}
		person := getPjRecordWithMainInfo(legInfo)
		person.bz = "备注" + strconv.Itoa(i+1)
		err := writer.Write(person.pjValues())
		utils.CheckError(PjDesc,"Cannot write to file", err)
	}
}

func getPjRecord() *LegPj {
	person := LegPj{}

	person.comp_name = core.GenLegName("徐州市")
	person.id_type = "L1"
	person.id_code = core.GenUnifyCode("91" + "320300")

	pjInfo(&person)

	return &person
}

func getPjRecordWithMainInfo(legInfo[3] string) *LegPj {
	person := LegPj{}

	person.comp_name = legInfo[0]
	person.id_type = legInfo[1]
	person.id_code = legInfo[2]

	pjInfo(&person)

	return &person
}

func pjInfo(person *LegPj) {
	var nr = [...]string{"表现良好", "表现优秀", "表现一般"}
	person.pjnr = core.RandDictItem(nr[:])

	var xz = [...]string{"行业评价", "监管评价", "官方评价"}
	person.pjxz = core.RandDictItem(xz[:])

	var jgmc = [...]string{"行业协会", "工商局", "税务局"}
	person.pjjgmc = core.RandDictItem(jgmc[:])

	person.pjrq = core.RandDate("2018-01-01", "2020-05-31", 1)[0]
}