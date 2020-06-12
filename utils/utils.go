package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

func CheckError(infoName, message string, err error) {
	if err != nil {
		log.Fatal(infoName, message, err)
	}
}

/*
 * 生成三位数字符串，不足前缀补零 11 -> 011
 */
func PadWithZero(bitNum, data int) string {
	format := "%0" + strconv.Itoa(bitNum) + "d"
	return fmt.Sprintf(format, data)
}


func FileWriter(fileName string) (*os.File, *csv.Writer){
	file, err := os.Create(fileName)
	CheckError("", "Cannot create file", err)

	writer := csv.NewWriter(file)

	return file, writer
}


func Fields(structInfo interface{}) [] string{
	fields := reflect.TypeOf(structInfo)
	num := fields.NumField()
	fieldNames := make([]string, num)

	for i := 0; i < num; i++ {
		field := fields.Field(i)
		fieldNames = append(fieldNames, field.Name)
	}
	return fieldNames
}

func WriteTitle(csvWriter *csv.Writer, title []string, infoName string) {
	var err = csvWriter.Write(title)
	CheckError(infoName,"Cannot write to %s File", err)
}

func WriteInfo(csvWriter *csv.Writer, data [] string, infoName string) {
	pjErr := csvWriter.Write(data)
	CheckError(infoName,"Cannot write to pjFile", pjErr)
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return false, err
}

func TimeMeasurement(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s", elapsed)
}