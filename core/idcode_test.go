package core

import (
	"log"
	"testing"
)

func TestGenIdCodeAndName(t *testing.T) {
	//rand.Seed(time.Now().Unix())
	//name, idCode, sex, birth := GenIdCodeAndName("320300")
	unifyCode := GenUnifyCode("320300")
	log.Printf("unifyCode:%s", unifyCode)

	//fmt.Print(strconv.Atoi("3"))
}
