package dashtime

import (
	"errors"
	"time"
)

const (
	Loc_UTC     = "UTC"                 //世界协调时间 标准时区
	Loc_Local   = "Local"               //系统所在当地 标准时区
	Loc_China   = "Asia/Shanghai"       //亚洲中国上海 标准时区
	Loc_Germany = "Europe/Berlin"       //欧洲德国柏林 标准时区
	Loc_America = "America/Los_Angeles" //美洲美国洛杉矶 标准时区
)

const (
	Lay_YmDHMS_CST = "2006-01-02 15:04:05 +0800 CST"
	Lay_YmDHMS     = "2006-01-02 15:04:05"
	Lay_YmDHM      = "2006-01-02 15:04"
	Lay_YmDH       = "2006-01-02 15"
	Lay_YmD        = "2006-01-02"
	Lay_HMS        = "15:04:05"

	Lay_YmDHMS_S = "20060102150405"
	Lay_YmDHM_S  = "200601021504"
	Lay_YmDH_S   = "2006010215"
	Lay_YmD_S    = "20060102"
	Lay_HMS_S    = "150405"

	Lay_YmDHMS_ZH = "2006年01月02日15点04分05秒"
	Lay_YmDHM_ZH  = "2006年01月02日15点04分"
	Lay_YmDH_ZH   = "2006年01月02日15点"
	Lay_YmD_ZH    = "2006年01月02日"
	Lay_HMS_ZH    = "15点04分05秒"
	Lay_YmDHMS_zh = "2006年01月02日15:04:05"
	Lay_YmDHM_zh  = "2006年01月02日15:04"
	Lay_YmDH_zh   = "2006年01月02日15"
	Lay_YmD_zh    = "2006年01月02日"
	Lay_HMS_zh    = "15:04:05"

	Lay_Date        = Lay_YmD
	Lay_Time        = Lay_HMS
	Lay_DateTime    = Lay_YmDHMS
	Lay_Date_S      = Lay_YmD_S
	Lay_Time_S      = Lay_HMS_S
	Lay_DateTime_S  = Lay_YmDHMS_S
	Lay_Date_ZH     = Lay_YmD_ZH
	Lay_Time_ZH     = Lay_HMS_ZH
	Lay_DateTime_ZH = Lay_YmDHMS_ZH
	Lay_Date_zh     = Lay_YmD_zh
	Lay_Time_zh     = Lay_HMS_zh
	Lay_DateTime_zh = Lay_YmDHMS_zh

	Lay_Layout      = time.Layout
	Lay_ANSIC       = time.ANSIC
	Lay_UnixDate    = time.UnixDate
	Lay_RubyDate    = time.RubyDate
	Lay_RFC822      = time.RFC822
	Lay_RFC822Z     = time.RFC822Z
	Lay_RFC850      = time.RFC850
	Lay_RFC1123     = time.RFC1123
	Lay_RFC1123Z    = time.RFC1123Z
	Lay_RFC3339     = time.RFC3339
	Lay_RFC3339Nano = time.RFC3339Nano
	Lay_Kitchen     = time.Kitchen
	Lay_Stamp       = time.Stamp
	Lay_StampMilli  = time.StampMilli
	Lay_StampMicro  = time.StampMicro
	Lay_StampNano   = time.StampNano
)

/*
Layout @Editor robotyang at 2023

# Layout 解析任意时间字符串的布局类型

@Param timStr 待解析时间字符串
*/
func Layout(timStr string) (string, error) {
	layouts := []string{
		Lay_YmDHMS_CST,
		Lay_YmDHMS,
		Lay_YmDHM,
		Lay_YmDH,
		Lay_YmD,
		Lay_HMS,
		Lay_YmDHMS_S,
		Lay_YmDHM_S,
		Lay_YmDH_S,
		Lay_YmD_S,
		Lay_HMS_S,
		Lay_YmDHMS_ZH,
		Lay_YmDHM_ZH,
		Lay_YmDH_ZH,
		Lay_YmD_ZH,
		Lay_HMS_ZH,
		Lay_YmDHMS_zh,
		Lay_YmDHM_zh,
		Lay_YmDH_zh,
		Lay_YmD_zh,
		Lay_HMS_zh,
		time.Layout,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	for _, layout := range layouts {
		if _, err := time.Parse(layout, timStr); err == nil {
			return layout, nil
		}
	}
	return "", errors.New("not yet supported layout")
}

/*
Parse @Editor robotyang at 2023

# Parse 解析任意时间字符串为 *time.Time

@Param timStr 待解析时间字符串
*/
func Parse(timStr string) (*time.Time, error) {
	if lay, err := Layout(timStr); err != nil {
		return nil, err
	} else {
		tim, err := time.Parse(lay, timStr)
		return &tim, err
	}
}

/*
ParseLoc @Editor robotyang at 2023

# ParseLoc 解析任意时间字符串为 *time.Time（并指定时区）

@Param timStr 待解析时间字符串

@Param loc 指定时区
*/
func ParseLoc(timStr string, locStr string) (*time.Time, error) {
	if lay, err := Layout(timStr); err != nil {
		return nil, err
	} else {
		var loc, _ = time.LoadLocation(locStr)
		tim, err := time.ParseInLocation(lay, timStr, loc)
		return &tim, err
	}
}

/*
DuraNextDawn @Editor robotyang at 2023

# DuraNextDawn  获取 times 对应到 凌晨的时间(到明天凌晨零点的时间)
*/
func DuraNextDawn(tim *time.Time) *time.Duration {
	next := tim.AddDate(0, 0, 1)
	nextDawn := time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
	dura := nextDawn.Sub(*tim)
	return &dura
}
