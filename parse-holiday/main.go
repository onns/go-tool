package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/onns/go-tool/config-load"
)

/*
@Time : 2021/10/16 00:58
@Author : onns
@File : parse-holiday/main.go
*/

type Config struct {
	S string `json:"s"`
}

func main() {
	var (
		b   []byte
		err error
	)
	cfg := &Config{
	}
	b, err = config_load.LoadConfig("")
	if err != nil {
		log.Printf("%v", err)
	}
	if err = json.Unmarshal(b, cfg); err != nil {
		log.Printf("%v", err)
	}
	// log.Printf("%v", cfg)
	res := parseHoliday(cfg.S)
	log.Println(res)
	genIcs(res)
}


type Year struct {
	Workday   []*Day // 工作日
	MakeupDay []*Day // 补班
	Holiday   []*Day // 假期
	Weekend   []*Day // 周末
}
type Day struct {
	Time time.Time
	Name string
	Desc string
}

func parseHoliday(s string) (res *Year) {
	res = &Year{
		Workday:   make([]*Day, 0),
		MakeupDay: make([]*Day, 0),
		Holiday:   make([]*Day, 0),
		Weekend:   make([]*Day, 0),
	}
	year := findYear(s)
	// log.Println(year)
	if year == 0 {
		return
	}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		line = strings.Trim(line, " \b\n\r\t")
		if len(line) <= 0 {
			continue
		}
		name := findName(line)
		// log.Println(name)
		k, e, o, c := findTime(line, year)
		i := 1
		for t := k; !t.After(e); t = t.Add(time.Hour * 24) {
			res.Holiday = append(res.Holiday, &Day{
				Time: t,
				Name: formatH(name, i, c),
				Desc: line[2:],
			})
			i += 1
		}
		if o {
			t, c := findOvertime(line, year)
			i = 1
			for _, day := range t {
				res.MakeupDay = append(res.MakeupDay, &Day{
					Time: day,
					Name: formatO(name, i, c),
					Desc: line[2:],
				})
				i += 1
			}
		}
	}
	return
}

func findYear(s string) (year int) {
	re, err := regexp.Compile(`([0-9]{4})年`)
	if err != nil {
		return
	}
	t := re.FindStringSubmatch(s)
	// log.Print(t)
	year, err = strconv.Atoi(t[1])
	if err != nil {
		return
	}
	return
}

func findName(s string) (name string) {
	re, err := regexp.Compile(`、([\S]{2,6})：`)
	if err != nil {
		return
	}
	name = re.FindStringSubmatch(s)[1]
	return
}

func findTime(s string, year int) (start, end time.Time, isDaysOff bool, count int) {
	var (
		sm, sd, em, ed int
	)
	re, err := regexp.Compile(`([0-9]{1,2})月([0-9]{1,2})日至(([0-9]{1,2})月)?([0-9]{1,2})日放假(调休)?，共([0-9]{1,2})天`)
	if err != nil {
		return
	}
	t := re.FindStringSubmatch(s)
	// log.Println(t, len(t))
	sm, err = strconv.Atoi(t[1])
	sd, err = strconv.Atoi(t[2])
	if t[3] == "" {
		em = sm
	} else {
		em, err = strconv.Atoi(t[4])
	}
	ed, err = strconv.Atoi(t[5])
	start = time.Date(year, time.Month(sm), sd, 0, 0, 0, 0, time.Now().Location())
	end = time.Date(year, time.Month(em), ed, 0, 0, 0, 0, time.Now().Location())
	if t[6] != "" {
		isDaysOff = true
	}
	count, err = strconv.Atoi(t[7])
	// log.Println(int(end.Sub(start).Hours()), (count-1) * 24)
	if int(end.Sub(start).Hours()) != (count-1)*24 {
		log.Printf("day not match!")
	}
	return
}

func findOvertime(s string, year int) (res []time.Time, count int) {
	// TODO 调休只能算两天的最多
	res = make([]time.Time, 0)
	re, err := regexp.Compile(`(([0-9]{1,2})月([0-9]{1,2})日（星期([\S])）、)?([0-9]{1,2})月([0-9]{1,2})日（星期([\S])）上班`)
	if err != nil {
		return
	}
	t := re.FindStringSubmatch(s)
	log.Println(t, len(t))
	if t[2] != "" {
		count += 1
		m, _ := strconv.Atoi(t[2])
		d, _ := strconv.Atoi(t[3])
		o := time.Date(year, time.Month(m), d, 0, 0, 0, 0, time.Now().Location())
		res = append(res, o)
	}
	if t[5] != "" {
		count += 1
		m, _ := strconv.Atoi(t[5])
		d, _ := strconv.Atoi(t[6])
		o := time.Date(year, time.Month(m), d, 0, 0, 0, 0, time.Now().Location())
		res = append(res, o)
	}
	return
}

func formatH(name string, i, t int) (res string) {
	res = fmt.Sprintf("%s 假期 第%d天/共%d天", name, i, t)
	return
}

func formatO(name string, i, t int) (res string) {
	res = fmt.Sprintf("%s 补班 第%d天/共%d天", name, i, t)
	return
}

func genIcs(res *Year) {
	for _,day := range res.Holiday {
		
	}
}