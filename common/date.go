package common

/*
@Time : 2022/5/8 11:57
@Author : onns
@File : common/date.go
*/

type AnniversaryType int8

const (
	Birthday      AnniversaryType = 1
	LunarBirthday AnniversaryType = 2
	SpecialDay    AnniversaryType = 3
	OneDay        AnniversaryType = 4
	Countdown     AnniversaryType = 5
)

type MemDays struct {
	Type AnniversaryType `json:"type"`
	Name string          `json:"name"`
	Date string          `json:"date"`
}

const DefaultDateFormat = "2006-01-02"
