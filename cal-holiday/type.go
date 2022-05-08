package main

import (
	"errors"

	"github.com/onns/go-tool/cal-holiday/cal"
)

/*
@Time : 2022/5/8 11:49
@Author : onns
@File : cal-holiday/type.go
*/

const (
	DefaultStart = 1900
	DefaultEnd   = 2100
)

type HolidayType int8

const (
	MothersDay HolidayType = 1
	FathersDay HolidayType = 2
)

type Holiday struct {
	TypeRaw string `json:"type"`
	Type    HolidayType
	Start   int `json:"start"`
	End     int `json:"end"`
	Cal     cal.HolidayCalculator
}

func (m *Holiday) Format() (err error) {
	err = m.ParseDuration()
	if err != nil {
		return
	}
	err = m.ParseType()
	if err != nil {
		return
	}
	return
}

func (m *Holiday) ParseType() (err error) {
	switch m.TypeRaw {
	case "母亲节":
		m.Type = MothersDay
		m.Cal = &cal.MothersDayCalculator{
			Name: m.TypeRaw,
		}
	case "父亲节":
		m.Type = FathersDay
		m.Cal = &cal.FathersDayCalculator{
			Name: m.TypeRaw,
		}
	default:
		err = errors.New("未定义的节日类型")
		return
	}
	return
}

func (m *Holiday) ParseDuration() (err error) {
	if m.Start == 0 {
		m.Start = DefaultStart
	}
	if m.End == 0 {
		m.End = DefaultEnd
	}
	return
}


