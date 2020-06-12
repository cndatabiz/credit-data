package main

import (
	"credit_data/config"
	"credit_data/dir"
	"credit_data/utils"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

var (
	cfg config.AppConfig
)

func init()  {
	cfg = getConf()
}

func getConf() config.AppConfig {
	conf := config.AppConfig{}
	viper.SetConfigName("app-settings")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v", err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}

func main() {
	defer utils.TimeMeasurement(time.Now())

	if len(os.Args) > 1 {
		cfg.RcdNumber,_ = strconv.Atoi(os.Args[1])
	}else{
		log.Println("Default output 29 row records.")
	}
	createOutDir()
	log.Println("Start generator data.")

	rand.Seed(time.Now().Unix())
	dir.GenLegInData(cfg.RcdNumber)

	log.Printf("Success generator %d rows of data.", cfg.RcdNumber)
}

func createOutDir() {
	outPath := path.Join("./", cfg.OutputPath)
	if flag,_ := utils.Exists(outPath); !flag {
		_ = os.Mkdir(outPath, 755)
	}
}

