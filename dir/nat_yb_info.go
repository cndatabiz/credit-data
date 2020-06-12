package dir

import (
	"credit_data/core"
	"fmt"
	"time"
)

var NatYbDesc = "自然人医保信息"

type NatYb struct {
	name      string
	id_type   string
	id_code   string
	yb_date   time.Time
	yb_amount float64
	bz        string
}

var NatYbTitle = [...]string{"name","id_type","id_code","yb_date","yb_amount","bz"}


func (s * NatYb) NatYbValues() [] string  {
	var line = make([] string ,len(NatYbTitle))
	line[0] = s.name
	line[1] = s.id_type
	line[2] = s.id_code
	line[3] = s.yb_date.Format("2006-01-02")
	line[4] = fmt.Sprintf("%.2f", s.yb_amount)
	line[5] = s.bz

	return line
}

func Get3YearsYbRecord(natInfo[3] string) [] *NatYb{
	persons := [] *NatYb{}
	startDate := time.Now().AddDate(-3, 0, 0)
	var curentDate time.Time
	for i := 0; i < 36; i++ {
		curentDate = startDate.AddDate(0,i,0)
		amount := core.RandFloats(300.0,1024.99,1,2)[0]
		ybRow := getNatYbRecordWithMainInfo(natInfo, curentDate, amount)
		persons = append(persons, ybRow)
	}
	return persons
}

func getNatYbRecordWithMainInfo(natInfo[3] string, ybDate time.Time, amount float64) *NatYb {
	person := NatYb{}

	person.name = natInfo[0]
	person.id_type = natInfo[1]
	person.id_code = natInfo[2]

	person.yb_date = ybDate
	person.yb_amount = amount

	return &person
}

