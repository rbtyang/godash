package dashtime

import "time"

const (
	Loc_UTC     = "UTC"                 //世界协调时间 标准时区
	Loc_China   = "Asia/Shanghai"       //亚洲中国上海 标准时区
	Loc_Germany = "Europe/Berlin"       //欧洲德国柏林 标准时区
	Loc_America = "America/Los_Angeles" //美洲美国洛杉矶 标准时区

	Fmt_Y_m_D_H_M_S     = "2006-01-02 15:04:05"
	Fmt_H_M_S           = "15:04:05"
	Fmt_Y_m_D_H_M_S_CST = "2006-01-02 15:04:05 +0800 CST"
	Fmt_YmDHMS          = "20060102150405"
	Fmt_YmDhMS          = "20060102030405"
	Fmt_Y_m_D           = "2006-01-02"
	Fmt_YmD             = "20060102"
	Fmt_Zh_Y_m_D_H_M    = "2006年01月02日15:04"
	Fmt_Zh_Y_m_D_H_M_S  = "2006年01月02日15:04:05"
	Fmt_Y_m_D_H_M       = "2006-01-02 15:04"
	Fmt_Date            = Fmt_Y_m_D
	Fmt_Time            = Fmt_H_M_S
	Fmt_DateTime        = Fmt_Y_m_D_H_M_S
)

/*
ParseToChina @Editor robotyang at 2023

# ParseToChina 格式化 文本时间 (默认:2006-01-02 15:04:05) 为time.Time
*/
func ParseToChina(timeStr string, format ...string) (time.Time, error) {
	defFormat := Fmt_Y_m_D_H_M_S
	for _, value := range format {
		defFormat = value
		break
	}
	var cstSh, _ = time.LoadLocation(Loc_China)
	timeUnix, err := time.ParseInLocation(defFormat, timeStr, cstSh)
	if err != nil {
		return time.Time{}, err
	}
	return timeUnix, nil
}

/*
Format @Editor robotyang at 2023

# Format  格式化时间(time.Time)为指定时间格式文本

@Param time：待格式化时间

@Param format：可选格式，否则默认格式 Fmt_Y_m_D_H_M_S
*/
func Format(times time.Time, format ...string) string {
	defFormat := Fmt_Y_m_D_H_M_S
	for _, value := range format {
		defFormat = value
		break
	}
	return times.Format(defFormat)
}

/*
CurrentTimeString @Editor robotyang at 2023

# CurrentTimeString  获取 当前时间 并转换成 自定格式
*/
func CurrentTimeString(defaultFormat ...string) string {
	format := Fmt_Y_m_D_H_M_S
	for _, value := range defaultFormat {
		format = value
	}
	loc, _ := time.LoadLocation(Loc_China)
	return time.Now().In(loc).Format(format)
}

/*
RestNextDawn @Editor robotyang at 2023

# RestNextDawn  获取 now 对应到 凌晨的时间(到明天凌晨零点的时间)
*/
func RestNextDawn(times time.Time) time.Duration {
	nextDay := times.AddDate(0, 0, 1)
	nextDay = time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), 0, 0, 0, 0, nextDay.Location())
	return nextDay.Sub(times)
}

/*
LocChina @Editor robotyang at 2023

# LocChina  设置time为 Asia/Shanghai默认时区
*/
func LocChina(times time.Time) time.Time {
	var loc, _ = time.LoadLocation(Loc_China)
	return times.In(loc)
}
