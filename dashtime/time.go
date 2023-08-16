package dashtime

import "time"

const (
	TimeLocDefault           = "Asia/Shanghai" // 默认时区
	TimeFmt_Y_m_D_H_M_S      = "2006-01-02 15:04:05"
	TimeFmt_H_M_S            = "15:04:05"
	TimeFmt_Y_m_D_H_M_S_CST  = "2006-01-02 15:04:05 +0800 CST"
	TimeFmt_YmDHMS           = "20060102150405"
	TimeFmt_YmDhMS           = "20060102030405"
	TimeFmt_Y_m_D            = "2006-01-02"
	TimeFmt_YmD              = "20060102"
	TimeFmtChinese_Y_m_D_H_M = "2006年01月02日15:04"
	TimeFmt_Y_m_D_H_M        = "2006-01-02 15:04"
	TimeFmtDate              = TimeFmt_Y_m_D
	TimeFmtTime              = TimeFmt_H_M_S
	TimeFmtDateTime          = TimeFmt_Y_m_D_H_M_S
)

/*
Format  格式化时间(time.Time)为指定时间格式文本
*/
func Format(timeSuk *time.Time, defaultFormat ...string) string {
	if timeSuk == nil {
		return ""
	}
	format := TimeFmt_Y_m_D_H_M_S
	for _, value := range defaultFormat {
		format = value
	}
	return timeSuk.Format(format)
}

/*
ParseString  格式化 文本时间 (默认:2006-01-02 15:04:05) 为time.Time
*/
func ParseString(timeStr string, defaultFormat ...string) (time.Time, error) {
	format := TimeFmt_Y_m_D_H_M_S
	for _, value := range defaultFormat {
		format = value
	}
	var cstSh, _ = time.LoadLocation(TimeLocDefault)
	timeUnix, err := time.ParseInLocation(format, timeStr, cstSh)
	if err != nil {
		return time.Time{}, err
	}
	return timeUnix, nil
}

/*
CurrentTimeString  获取 当前时间 并转换成 自定格式
*/
func CurrentTimeString(defaultFormat ...string) string {
	format := TimeFmt_Y_m_D_H_M_S
	for _, value := range defaultFormat {
		format = value
	}
	loc, _ := time.LoadLocation(TimeLocDefault)
	return time.Now().In(loc).Format(format)
}

/*
RestNextDawn  获取 now 对应到 凌晨的时间(到明天凌晨零点的时间)
*/
func RestNextDawn(now time.Time) time.Duration {
	nextDay := now.AddDate(0, 0, 1)
	nextDay = time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), 0, 0, 0, 0, nextDay.Location())
	return nextDay.Sub(now)
}

/*
CurrentTimePointer  获取 now时间的 *time.Time 格式
*/
func CurrentTimePointer() *time.Time {
	now := time.Now()
	now = SetLocDefault(&now)
	return &now
}

/*
SetLocDefault  设置time为 Asia/Shanghai默认时区
*/
func SetLocDefault(timeSuk *time.Time) time.Time {
	var loc, _ = time.LoadLocation(TimeLocDefault)
	return timeSuk.In(loc)
}
