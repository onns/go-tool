package cal

import (
	"time"

	"github.com/onns/go-tool/common"
)

/*
@Time : 2022/5/8 12:04
@Author : onns
@File : cal-holiday/cal/cal.go
*/

type HolidayCalculator interface {
	Output(start, end int) (res []*common.MemDays, err error)
}

type MothersDayCalculator struct {
	Name string
}

func (m *MothersDayCalculator) Output(start, end int) (res []*common.MemDays, err error) {
	res = make([]*common.MemDays, 0)
	for i := start; i <= end; i++ {
		day := time.Date(i, 5, 1, 0, 0, 0, 0, time.Local)
		count := 0
		for {
			if day.Weekday() == time.Sunday {
				count++
			}
			if count == 2 {
				break
			}
			day = day.AddDate(0, 0, 1)
		}
		res = append(res, &common.MemDays{
			Type: common.OneDay,
			Name: m.Name,
			Date: day.Format(common.DefaultDateFormat),
		})
	}
	return
}

type FathersDayCalculator struct {
	Name string
}

func (m *FathersDayCalculator) Output(start, end int) (res []*common.MemDays, err error) {
	res = make([]*common.MemDays, 0)
	for i := start; i <= end; i++ {
		day := time.Date(i, 6, 1, 0, 0, 0, 0, time.Local)
		count := 0
		for {
			if day.Weekday() == time.Sunday {
				count++
			}
			if count == 3 {
				break
			}
			day = day.AddDate(0, 0, 1)
		}
		res = append(res, &common.MemDays{
			Type: common.OneDay,
			Name: m.Name,
			Date: day.Format(common.DefaultDateFormat),
		})
	}
	return
}
