package dir

import (
	"credit_data/core"
	"credit_data/utils"
	"strings"
	"time"
)

var NatXzjlDesc = "行政奖励信息"

type NatXzjl struct {
	name    string
	id_type string
	id_code string
	rdwsh   string
	jlmc    string
	jlsx    string
	jlxs    string
	jlsyrq  time.Time
	yxjzq   time.Time
	syjg    string
	bz      string
}

var NatXzjlTitle = [...]string{"name","id_type","id_code","rdwsh","jlmc","jlsx","jlxs","jlsyrq","yxjzq","syjg","bz"}


func (s * NatXzjl) NatXzjlValues() [] string  {
	var line = make([] string ,len(NatXzjlTitle))
	line[0] = s.name
	line[1] = s.id_type
	line[2] = s.id_code
	line[3] = s.rdwsh
	line[4] = s.jlmc
	line[5] = s.jlsx
	line[6] = s.jlxs
	line[7] = s.jlsyrq.Format("2006-01-02")
	line[8] = s.yxjzq.Format("2006-01-02")
	line[9] = s.syjg
	line[10] = s.bz

	return line
}


func getNatXzjlRecordWithMainInfo(natInfo [3] string) *NatXzjl {
	person := NatXzjl{}

	person.name = natInfo[0]
	person.id_type = natInfo[1]
	person.id_code = natInfo[2]

	NatXzjlInfo(&person)

	return &person
}

func NatXzjlInfo(person *NatXzjl) {
	person.rdwsh = utils.PadWithZero(11, int(core.RandInteger(1, 1000000000, 1)[0]))
	var jlmc = [...]string{"科技发明创新奖", "帮扶市民贡献奖", "抗洪防旱特出奖", "环保特出贡献奖"}
	person.jlmc = core.RandDictItem(jlmc[:])

	if strings.Contains(person.jlmc, "科技") {
		person.syjg = "科技局"
		person.rdwsh = "T" + person.rdwsh
	}else if strings.Contains(person.jlmc, "帮扶"){
		person.syjg = "民政局"
		person.rdwsh = "C" + person.rdwsh
	}else if strings.Contains(person.jlmc, "抗洪"){
		person.syjg = "水利局"
		person.rdwsh = "W" + person.rdwsh
	}else if strings.Contains(person.jlmc, "环保"){
		person.syjg = "环保局"
		person.rdwsh = "E" + person.rdwsh
	}

	var jlsx = [...]string{"一等奖", "二等奖", "三等奖"}
	person.jlsx = person.jlmc + core.RandDictItem(jlsx[:])

	var jlxs = [...]string{"0", "1", "2","3", "4", "5", "9"}
	person.jlxs = core.RandDictItem(jlxs[:])

	person.jlsyrq = core.RandDate("2018-01-01", "2020-05-31", 1)[0]
	person.yxjzq = person.jlsyrq.Add(time.Hour*24*365*3)
}