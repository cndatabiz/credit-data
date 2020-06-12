package dir

import (
	"credit_data/core"
	"time"
)

var ZscqDesc = "评价信息"

type LegZscq struct {
	comp_name string
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

var ZscqTitle = [...]string{"comp_name","id_type","id_code","zscqdjzh","zscqmc","zscqzl","dljg","dlr","djjg","djrq"}

func (s * LegZscq) zscqValues() [] string  {
	var line = make([] string ,len(ZscqTitle))
	line[0] = s.comp_name
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

func getZscqRecordWithMainInfo(legInfo[3] string) *LegZscq {
	person := LegZscq{}

	person.comp_name = legInfo[0]
	person.id_type = legInfo[1]
	person.id_code = legInfo[2]

	zscqInfo(&person)

	return &person
}

func zscqInfo(person *LegZscq) {
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