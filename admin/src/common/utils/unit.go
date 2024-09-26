package utils

/*
 * 自定义单位换算器
 * 依照标准字典换算
 * 默认(毫秒) '0' 秒 '1' 分 '2' 时 '3' 日 '4' 月 '5' 年 '6'
 */

var unitMap = map[string]uint64{
	"0": 1,
	"1": 1000,
	"2": 60000,       // 60 * 1000,
	"3": 3600000,     // 60 * 60 * 1000,
	"4": 86400000,    // 24 * 60 * 60 * 1000,
	"5": 2592000000,  // 30 * 24 * 60 * 60 * 1000,
	"6": 30758400000, // 356 * 24 * 60 * 60 * 1000,
}

// GetMilliseconds 获取毫秒单位周期
func GetMilliseconds(num uint16, unit string) (res uint64) {
	unitDur := unitMap[unit]
	res = uint64(num) * unitDur
	return
}
