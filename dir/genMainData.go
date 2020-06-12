package dir

import (
	"credit_data/core"
	"credit_data/utils"
	"strconv"
	"strings"
)

func GenLegInData(rows int)  {
	natBaseFile,natBaseInfoWriter := utils.FileWriter("result/natBaseInfo.csv")
	defer natBaseFile.Close()
	defer natBaseInfoWriter.Flush()

	legBaseFile,legBaseInfoWriter := utils.FileWriter("result/legBaseInfo.csv")
	defer legBaseFile.Close()
	defer legBaseInfoWriter.Flush()


	pjFile,pjWriter := utils.FileWriter("result/legPj.csv")
	defer pjFile.Close()
	defer pjWriter.Flush()

	zsFile,zsWriter := utils.FileWriter("result/legxzzs.csv")
	defer zsFile.Close()
	defer zsWriter.Flush()

	blFile,blWriter := utils.FileWriter("result/legsybl.csv")
	defer blFile.Close()
	defer blWriter.Flush()

	tsFile,tsWriter := utils.FileWriter("result/legtousu.csv")
	defer tsFile.Close()
	defer tsWriter.Flush()

	cqFile,cqWriter := utils.FileWriter("result/legzscq.csv")
	defer cqFile.Close()
	defer cqWriter.Flush()

	tdFile,tdWriter := utils.FileWriter("result/legtudi.csv")
	defer tdFile.Close()
	defer tdWriter.Flush()

	natTsFile,natTsWriter := utils.FileWriter("result/nattousu.csv")
	defer natTsFile.Close()
	defer natTsWriter.Flush()

	natCqFile,natCqWriter := utils.FileWriter("result/natzscq.csv")
	defer natCqFile.Close()
	defer natCqWriter.Flush()

	natPjFile,natPjWriter := utils.FileWriter("result/natpj.csv")
	defer natPjFile.Close()
	defer natPjWriter.Flush()

	natXzjlFile,natXzjlWriter := utils.FileWriter("result/natxzjl.csv")
	defer natXzjlFile.Close()
	defer natXzjlWriter.Flush()

	natYbFile,natYbWriter := utils.FileWriter("result/natYb.csv")
	defer natYbFile.Close()
	defer natYbWriter.Flush()

	utils.WriteTitle(natBaseInfoWriter, NatBaseInfoTitle[:], NatBaseInfoDesc)
	utils.WriteTitle(legBaseInfoWriter, LegBaseInfoTitle[:], LegBaseInfoDesc)
	utils.WriteTitle(pjWriter, pjTitle[:], PjDesc)
	utils.WriteTitle(zsWriter, xzzsTitle[:], XzzsDesc)
	utils.WriteTitle(blWriter, blTitle[:], BlDesc)
	utils.WriteTitle(tsWriter, tsTitle[:], TsDesc)
	utils.WriteTitle(cqWriter, ZscqTitle[:], ZscqDesc)
	utils.WriteTitle(tdWriter, TdTitle[:], TdDesc)

	utils.WriteTitle(natTsWriter, NatTouSuTitle[:], NatTouSuDesc)
	utils.WriteTitle(natCqWriter, NatZscqTitle[:], NatZscqDesc)
	utils.WriteTitle(natPjWriter, NatPjTitle[:], NatPjDesc)
	utils.WriteTitle(natXzjlWriter, NatXzjlTitle[:], NatXzjlDesc)
	utils.WriteTitle(natYbWriter, NatYbTitle[:], NatYbDesc)

	var legNameSet = make(map[string] bool)

	for i := 0; i < rows; i++ {
		var regions = [30]string{
			"320300:徐州市","320382:邳州市","320381:新沂市","320100:南京市","320200:无锡市",
			"320400:常州市","320481:溧阳市","320581:常熟市","320830:南通市","321200:泰州市",
			"321300:宿迁市","321100:镇江市","321181:丹阳市","321183:句容市","320811:昆山市",
			"370100:济南市","370200:青岛市","370285:莱西市","370300:淄博市","370400:枣庄市",
			"370500:东营市","370600:烟台市","370700:潍坊市","370800:泰安市","371000:威海市",
			"371083:乳山市","371100:日照市","371600:滨州市","371300:临沂市","371400:德州市",
		}
		region := core.RandDictItem(regions[:])
		regionInfo := strings.Split(region,":")
		cityName := regionInfo[1]
		cityCode := regionInfo[0]
		legName := core.GenLegName(cityName)

		for legNameSet[legName] {
			legName = core.GenLegName(cityName)
			if !legNameSet[legName] {
				break
			}
		}
		legNameSet[legName] = true

		var legInfo = [3]string{
			legName, "L1", core.GenUnifyCode("91" + cityCode),
		}

		var xm, idCode, sex, birthday = core.GenIdCodeAndName(cityCode)
		var natInfo = [3]string{
			xm, "N1", idCode,
		}

		natBaseInfo := getNatBaseInfoRecordWithMainInfo(natInfo,sex, birthday)
		utils.WriteInfo(natBaseInfoWriter, natBaseInfo.NatBaseInfoValues(), NatBaseInfoDesc)

		legBaseInfo := getLegBaseRecordWithMainInfo(legInfo, natInfo)
		utils.WriteInfo(legBaseInfoWriter, legBaseInfo.LegBaseInfoValues(), NatBaseInfoDesc)

		pjInfo := getPjRecordWithMainInfo(legInfo)
		pjInfo.bz = PjDesc + "备注" + strconv.Itoa(i+1)
		utils.WriteInfo(pjWriter, pjInfo.pjValues(), PjDesc)

		zsInfo := getXzzsRecordWithMainInfo(legInfo, natInfo)
		zsInfo.bz = XzzsDesc + "备注" + strconv.Itoa(i+1)
		utils.WriteInfo(zsWriter, zsInfo.xzValues(), XzzsDesc)

		blInfo := getBlRecordWithMainInfo(legInfo, natInfo)
		blInfo.bz = BlDesc + "备注" + strconv.Itoa(i+1)
		utils.WriteInfo(blWriter, blInfo.blValues(), BlDesc)

		tsInfo := getTouSuRecordWithMainInfo(legInfo, natInfo)
		tsInfo.bz = TsDesc + "备注" + strconv.Itoa(i+1)
		utils.WriteInfo(tsWriter, tsInfo.tsValues(), BlDesc)

		cqInfo := getZscqRecordWithMainInfo(legInfo)
		utils.WriteInfo(cqWriter, cqInfo.zscqValues(), ZscqDesc)

		tdInfo := getTdRecordWithMainInfo(legInfo)
		utils.WriteInfo(tdWriter, tdInfo.tdValues(), TdDesc)

		natTsInfo := getNatTouSuRecordWithMainInfo(natInfo)
		natTsInfo.bz = NatTouSuDesc + "备注" + strconv.Itoa(i+1)
		utils.WriteInfo(natTsWriter, natTsInfo.NatTsValues(), NatTouSuDesc)

		natCqInfo := getNatZscqRecordWithMainInfo(natInfo)
		utils.WriteInfo(natCqWriter, natCqInfo.NatZscqValues(), NatZscqDesc)

		natPjInfo := getNatPjRecordWithMainInfo(natInfo)
		natPjInfo.bz = NatPjDesc + "备注" + strconv.Itoa(i+1)
		utils.WriteInfo(natPjWriter, natPjInfo.NatPjValues(), NatPjDesc)

		natXzjlInfo := getNatXzjlRecordWithMainInfo(natInfo)
		natXzjlInfo.bz = NatXzjlDesc + "备注" + strconv.Itoa(i+1)
		utils.WriteInfo(natXzjlWriter, natXzjlInfo.NatXzjlValues(), NatXzjlDesc)

		records := Get3YearsYbRecord(natInfo)
		for _,item := range records{
			utils.WriteInfo(natYbWriter, item.NatYbValues(), NatYbDesc)
		}
	}
}

