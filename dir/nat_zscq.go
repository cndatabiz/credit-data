package dir

import (
	"credit_data/core"
	"time"
)

var NatZscqDesc = "评价信息"

type NatZscq struct {
	name      string
	id_type   string
	id_code   string
	zscqdjzh  string
	zscqmc    string
	zscqzl    string
	dljg      string
	dlr       string
	djjg      string
	djrq      time.Time
}

var NatZscqTitle = [...]string{"name","id_type","id_code","zscqdjzh","zscqmc","zscqzl","dljg","dlr","djjg","djrq"}

func (s * NatZscq) NatZscqValues() [] string  {
	var line = make([] string ,len(NatZscqTitle))
	line[0] = s.name
	line[1] = s.id_type
	line[2] = s.id_code
	line[3] = s.zscqdjzh
	line[4] = s.zscqmc
	line[5] = s.zscqzl
	line[6] = s.dljg
	line[7] = s.dlr
	line[8] = s.djjg
	line[9] = s.djrq.Format("2006-01-02")

	return line
}

func getNatZscqRecordWithMainInfo(natInfo[3] string) *NatZscq {
	person := NatZscq{}

	person.name = natInfo[0]
	person.id_type = natInfo[1]
	person.id_code = natInfo[2]

	natZscqInfo(&person)

	return &person
}

func natZscqInfo(person *NatZscq) {
	person.zscqdjzh = core.GenIpNo()

	var zscqmc = [...]string{"xxx改进算法", "xxx处理机制", "xxx发明装置"}
	person.zscqmc = core.RandDictItem(zscqmc[:])
	person.zscqzl = "1"

	var dljg = [...]string{"巨匠知识产权服务中心", "华创知识产权服务机构", "创客知识产权服务机构"}
	person.dljg = core.RandDictItem(dljg[:])
	name,_ := core.GenNameSex()
	person.dlr = name
	person.djjg = "市科技局"

	person.djrq = core.RandDate("2018-01-01", "2020-05-31", 1)[0]
}