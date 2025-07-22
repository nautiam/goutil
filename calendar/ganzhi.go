package calendar

import (
	"time"
)

var (
	// 1900 2月春分后为 庚子年
	GanTable    = []string{"Giáp", "Ất", "Bính", "Đinh", "Mậu", "Kỷ", "Canh", "Tân", "Nhâm", "Quý"}
	ZhiTable    = []string{"Tý", "Sửu", "Dần", "Mão", "Thìn", "Tỵ", "Ngọ", "Mùi", "Thân", "Dậu", "Tuất", "Hợi"}
	AnimalTable = []string{"Chuột", "Trâu", "Hổ", "Mèo", "Rồng", "Rắn", "Ngựa", "Dê", "Khỉ", "Gà", "Chó", "Lợn"}
)

// Tian Gan
func Gan(x int) string {
	return GanTable[(x)%10]
}

// Di Zhi
func Zhi(x int) string {
	return ZhiTable[(x)%12]
}

// Tian Gan && Di Zhi
func GanZhi(x int) string {
	return Gan(x) + Zhi(x)
}

// Sheng Xiao
func AnimalYear(year int) string {
	return AnimalTable[(year-MinYear)%12]
}

// Input: solar (year,month,day)
func GanZhiYMD(sYear, sMonth, sDay int) (gzYear, gzMonth, gzDay string) {
	if sMonth < 2 {
		gzYear = GanZhi(sYear - MinYear + 36 - 1)
	} else {
		gzYear = GanZhi(sYear - MinYear + 36)
	}

	// 立春日期
	jq3 := JieQi(sYear, 3)

	// 月柱 1900年1月小寒以前为 丙子月(60进制12)
	firstNode := JieQi(sYear, sMonth*2-1) // 返回当月「节」为几日开始
	gzMonth = GanZhi((sYear-MinYear)*12 + sMonth + 11)

	// 依节气调整二月分的年柱, 以立春为界
	if sMonth == 2 && sDay >= jq3 {
		gzYear = GanZhi(sYear - MinYear + 36)
	}

	// 依节气月柱, 以「节」为界
	if sDay >= firstNode {
		gzMonth = GanZhi((sYear-MinYear)*12 + sMonth + 12)
	}

	// 当月一日与 1900/1/1 相差天数
	// 1900/1/1 日柱为甲戌日(60进制10)
	base2 := time.Date(MinYear, 1, 1, 0, 0, 0, 0, time.UTC)
	now2 := time.Date(sYear, time.Month(sMonth), 1, 0, 0, 0, 0, time.UTC)
	offset := int(now2.Sub(base2).Seconds()/86400) + 10

	// 日柱
	gzDay = GanZhi(offset + sDay - 1)

	return
}

// hour range [0,23]
// return: start point DiZhi
func ZhiHour(hour int) (zHour string) {
	zHour = ZhiTable[int((hour+1)/2)%12]
	return
}
