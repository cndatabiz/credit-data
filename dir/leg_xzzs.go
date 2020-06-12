package dir

import (
	"credit_data/core"
	"credit_data/utils"
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

var XzzsDesc = "行政征收信息"

type LegXzzs struct {
	comp_name string
	id_type   string
	id_code   string
	fddbrxm   string
	fddbrzjlx string
	fddbrzjhm string
	xzzszl    string
	xzzsnr    string
	szzsrq    time.Time
	xzzsjg    string
	bz        string
}

var xzzsTitle = [...]string{"comp_name","id_type","id_code","fddbrxm","fddbrzjlx","fddbrzjhm","xzzszl","xzzsnr","szzsrq","xzzsjg","bz"}

func (s * LegXzzs) xzValues() [] string  {
	var line = make([] string ,len(xzzsTitle))
	line[0] = s.comp_name
	line[1] = s.id_type
	line[2] = s.id_code
	line[3] = s.fddbrxm
	line[4] = s.fddbrzjlx
	line[5] = s.fddbrzjhm
	line[6] = s.xzzszl
	line[7] = s.xzzsnr
	line[8] = s.szzsrq.Format("2006-01-02")
	line[9] = s.xzzsjg
	line[10] = s.bz
	return line
}

func GenLegZsData(rows int)  {

	file, err := os.Create("legxzzs.csv")
	utils.CheckError(XzzsDesc,"Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(xzzsTitle[:])
	utils.CheckError(XzzsDesc,"Cannot write to file", err)

	for i := 0; i < rows; i++ {
		person := LegXzzs{}

		person.comp_name = core.GenLegName("徐州市")
		person.id_type = "L1"
		person.id_code = core.GenUnifyCode("91" + "320300")

		person.fddbrzjlx = "N1"
		var xm,idCode,_,_ = core.GenIdCodeAndName("320300")
		person.fddbrxm = xm
		person.fddbrzjhm = idCode

		person.xzzszl = "因违反行政法的规定而引起的征收"
		var zsnr = [...] string{"排污费","滞纳金"}
		person.xzzsnr = core.RandDictItem(zsnr[:])

		person.szzsrq = core.RandDate("2018-01-01","2020-05-31",1)[0]
		var jgmc = [...] string {"环保局","城市管理局","税务局"}
		person.xzzsjg = core.RandDictItem(jgmc[:])

		person.bz = "备注" + strconv.Itoa(i+1)

		err := writer.Write(person.xzValues())
		utils.CheckError(XzzsDesc,"Cannot write to file", err)
	}
}

func getXzzsRecord() * LegXzzs {
	person := LegXzzs{}

	person.comp_name = core.GenLegName("徐州市")
	person.id_type = "L1"
	person.id_code = core.GenUnifyCode("91" + "320300")

	person.fddbrzjlx = "N1"
	var xm,idCode,_,_ = core.GenIdCodeAndName("320300")
	person.fddbrxm = xm
	person.fddbrzjhm = idCode

	zsInfo(&person)

	return &person
}

func getXzzsRecordWithMainInfo(legInfo[3] string, netInfo[3] string) * LegXzzs {
	person := LegXzzs{}

	person.comp_name = legInfo[0]
	person.id_type = legInfo[1]
	person.id_code = legInfo[2]

	person.fddbrxm = netInfo[0]
	person.fddbrzjlx = netInfo[1]
	person.fddbrzjhm = netInfo[2]
	zsInfo(&person)

	return &person
}

func zsInfo(person *LegXzzs) {
	person.xzzszl = "因违反行政法的规定而引起的征收"
	var zsnr = [...]string{"排污费", "滞纳金"}
	person.xzzsnr = core.RandDictItem(zsnr[:])

	person.szzsrq = core.RandDate("2018-01-01", "2020-05-31", 1)[0]
	var jgmc = [...]string{"环保局", "城市管理局", "税务局"}
	person.xzzsjg = core.RandDictItem(jgmc[:])
}