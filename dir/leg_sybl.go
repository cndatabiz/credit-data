package dir

import (
	"credit_data/core"
	"credit_data/utils"
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

var BlDesc = "公用事业不良信息"

type LegBl struct {
	comp_name string
	id_type   string
	id_code   string
	fddbrxm   string
	fddbrzjlx string
	fddbrzjhm string
	wgxxnr    string
	rdjg      string
	rdrq      time.Time
	bz        string
}

var blTitle = [...]string{"comp_name","id_type","id_code","fddbrxm","fddbrzjlx","fddbrzjhm","wgxxnr","rdjg","rdrq","bz"}


func (s * LegBl) blValues() [] string  {
	var line = make([] string ,len(blTitle))
	line[0] = s.comp_name
	line[1] = s.id_type
	line[2] = s.id_code
	line[3] = s.fddbrxm
	line[4] = s.fddbrzjlx
	line[5] = s.fddbrzjhm
	line[6] = s.wgxxnr
	line[7] = s.rdjg
	line[8] = s.rdrq.Format("2006-01-02")
	line[9] = s.bz
	return line
}

func GenLegBlData(rows int)  {

	file, err := os.Create("legblxx.csv")
	utils.CheckError(BlDesc,"Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(blTitle[:])
	utils.CheckError(BlDesc,"Cannot write to file", err)

	for i := 0; i < rows; i++ {
		person := LegBl{}

		person.comp_name = core.GenLegName("徐州市")
		person.id_type = "L1"
		person.id_code = core.GenUnifyCode("91" + "320300")

		person.fddbrzjlx = "N1"
		var xm,idCode,_,_ = core.GenIdCodeAndName("320300")
		person.fddbrxm = xm
		person.fddbrzjhm = idCode

		var wfnr = [...] string{"欠费","欠费及滞纳金"}
		person.wgxxnr = core.RandDictItem(wfnr[:])

		var jgmc = [...] string {"水务局","供热公司","燃气公司","供电局","电信公司"}
		person.rdjg = core.RandDictItem(jgmc[:])
		person.rdrq = core.RandDate("2018-01-01","2020-05-31",1)[0]

		person.bz = "备注" + strconv.Itoa(i+1)

		err := writer.Write(person.blValues())
		utils.CheckError(BlDesc,"Cannot write to file", err)
	}
}

func getBlRecordWithMainInfo(legInfo[3] string, netInfo[3] string) * LegBl {
	person := LegBl{}

	person.comp_name = legInfo[0]
	person.id_type = legInfo[1]
	person.id_code = legInfo[2]

	person.fddbrxm = netInfo[0]
	person.fddbrzjlx = netInfo[1]
	person.fddbrzjhm = netInfo[2]
	blInfo(&person)

	return &person
}

func blInfo(person *LegBl) {
	var wfnr = [...] string{"欠费","欠费及滞纳金"}
	person.wgxxnr = core.RandDictItem(wfnr[:])

	var jgmc = [...] string {"水务局","供热公司","燃气公司","供电局","电信公司"}
	person.rdjg = core.RandDictItem(jgmc[:])
	person.rdrq = core.RandDate("2018-01-01","2020-05-31",1)[0]
}
