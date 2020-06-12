package dir

import (
	"credit_data/core"
	"time"
)

var NatPjDesc = "评价信息"

type NatPj struct {
	name string
	id_type   string
	id_code   string
	pjjg      string
	pjjgmc    string
	pjrq      time.Time
	pjjgyxqs  time.Time
	pjjgyxqz  time.Time
	dqzt      string
	bz        string
}

var NatPjTitle = [...]string{"name","id_type","id_code","pjjg","pjjgmc","pjrq","pjjgyxqs","pjjgyxqz","dqzt","bz"}


func (s * NatPj) NatPjValues() [] string  {
	var line = make([] string ,len(NatPjTitle))
	line[0] = s.name
	line[1] = s.id_type
	line[2] = s.id_code
	line[3] = s.pjjg
	line[4] = s.pjjgmc
	line[5] = s.pjrq.Format("2006-01-02")
	line[6] = s.pjjgyxqs.Format("2006-01-02")
	line[7] = s.pjjgyxqz.Format("2006-01-02")
	line[8] = s.dqzt
	line[9] = s.bz

	return line
}


func getNatPjRecordWithMainInfo(natInfo [3] string) *NatPj {
	person := NatPj{}

	person.name = natInfo[0]
	person.id_type = natInfo[1]
	person.id_code = natInfo[2]

	NatPjInfo(&person)

	return &person
}

func NatPjInfo(person *NatPj) {
	var pjjg = [...]string{"表现良好", "表现优秀", "表现一般"}
	person.pjjg = core.RandDictItem(pjjg[:])

	var jgmc = [...]string{"行业协会", "大公国际", "百行征信"}
	person.pjjgmc = core.RandDictItem(jgmc[:])

	person.pjrq = core.RandDate("2018-01-01", "2020-05-31", 1)[0]
	person.pjjgyxqs = person.pjrq.Add(time.Hour*24)
	person.pjjgyxqz = person.pjjgyxqs.Add(time.Hour*24*365*3)
	person.dqzt = "1"
}