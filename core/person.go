package core

import (
	"fmt"
	"math/rand"
	"time"
)

func main1() {
	rand.Seed(time.Now().Unix())

	//var xuzhou = map[string]string{
	//	"320300": "徐州市",
	//	"320302": "鼓楼区",
	//	"320303": "云龙区",
	//	"320305": "贾汪区",
	//	"320311": "泉山区",
	//	"320312": "铜山区",
	//	"320321": "丰县",
	//	"320322": "沛县",
	//	"320324": "睢宁县",
	//	"320381": "新沂市",
	//	"320382": "邳州市",
	//}

	for i:=0;i<90;i++ {
		fmt.Println(GenUnifyCode("91" + "320300"))
	}
}
