package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type Item struct {
	Text     string
	Prioirty int
	Position int
	Done     bool
	Date     string
	Deadline string
}

func SaveItems(filename string, items []Item) error {
	str, err := json.Marshal(items)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, str, 0644)
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}
	return nil
}
func ReadItems(filename string) ([]Item, error) {
	str, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	err = json.Unmarshal(str, &items)
	if err != nil {
		return []Item{}, err
	}
	for i := 0; i < len(items); i++ {
		items[i].Position = i + 1
	}
	return items, nil

}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Prioirty = 1
	case 3:
		i.Prioirty = 3
	default:
		i.Prioirty = 2
	}

}
func (i *Item) PrettyP() string {
	if i.Prioirty == 1 {
		return "(L)"
	}
	if i.Prioirty == 3 {
		return "(H)"
	}
	return ""
}

func (i *Item) Label() string {
	return strconv.Itoa(i.Position) + "."
}

func (i *Item) SetDate() {
	t := time.Now()
	i.Date = t.Format(time.UnixDate)
}
func (i *Item) GetDate() string {
	t, _ := time.Parse(time.UnixDate, i.Date)
	year, month, day := t.Date()
	hours, minutes := t.Hour(), t.Minute()
	return strconv.Itoa(day) + "/" + month.String() + "/" + strconv.Itoa(year) + " " + strconv.Itoa(hours) + ":" + strconv.Itoa(minutes)

}
func validateMonth(month string) time.Month {
	intMonth, err := strconv.Atoi(month)
	if err != nil || intMonth < 1 || intMonth > 12 {
		log.Fatal("Invalid Month ")
	}
	return time.Month(intMonth)
}
func validateYear(year string) int {
	intYear, err := strconv.Atoi(year)
	if err != nil || intYear < 0 {
		log.Fatal("Invalid Year ")
	}
	return intYear
}

func validateDay(day string) int {
	intDay, err := strconv.Atoi(day)
	if err != nil || intDay < 1 || intDay > 31 {
		log.Fatal("Invalid Day ")
	}
	return intDay
}

func validDateHour(hour string) int {
	intHour, err := strconv.Atoi(hour)
	if err != nil || intHour < 0 || intHour > 23 {
		log.Fatal("Invalid hour")
	}
	return intHour
}
func validateMinute(minute string) int {
	intMinute, err := strconv.Atoi(minute)
	if err != nil || intMinute < 0 || intMinute > 60 {
		log.Fatal("Invalid minute")
	}
	return intMinute
}

func (i *Item) SetDeadline(deadline string) {
	if strings.Compare(deadline, "none") == 0 {
		i.Deadline = "none"
		return
	}
	dateStr := strings.Split(deadline, " ")

	if len(dateStr) != 2 {
		log.Fatal("Invalid date time")
	}
	rawDateStrArr := strings.Split(dateStr[0], "/")

	if len(rawDateStrArr) != 3 {
		log.Fatal("Invalid Date")
	}

	rawTimeStrArr := strings.Split(dateStr[1], ":")

	if len(rawTimeStrArr) != 2 {
		log.Fatal("Invalid time")
	}

	month := validateMonth(rawDateStrArr[1])
	year := validateYear(rawDateStrArr[2])
	day := validateDay(rawDateStrArr[0])
	hour := validDateHour(rawTimeStrArr[0])
	minutes := validateMinute(rawTimeStrArr[1])
	date := time.Date(year, month, day, hour, minutes, 0, 0, time.UTC)
	i.Deadline = date.Format(time.UnixDate)

}
func convertToDoubleDigit(str string) string {
	if len(str) == 1 {
		str = "0" + str
	}
	return str
}
func (i *Item) GetDeadline() string {
	if i.Deadline == "none" {
		return ""
	}
	t, _ := time.Parse(time.UnixDate, i.Deadline)
	year, month, day := t.Date()
	monthStr := convertToDoubleDigit(month.String())
	dayStr := convertToDoubleDigit(strconv.Itoa(day))
	yearStr := strconv.Itoa(year)
	hours, minutes := convertToDoubleDigit(strconv.Itoa(t.Hour())), convertToDoubleDigit(strconv.Itoa(t.Minute()))

	return dayStr + "/" + monthStr + "/" + yearStr + " " + hours + ":" + minutes

}
func Reverse(items []Item) []Item {
	len := len(items)
	for i := 0; i < len/2; i++ {
		items[i], items[len-i-1] = items[len-i-1], items[i]
	}
	return items
}

type ByPri []Item

func (s ByPri) Len() int      { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return s[i].Done
	}
	if s[i].Prioirty == s[j].Prioirty {
		return s[i].Position > s[j].Position
	}
	return s[i].Prioirty > s[j].Prioirty
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	}
	return ""
}

type ByDeadline []Item

func (s ByDeadline) Len() int      { return len(s) }
func (s ByDeadline) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByDeadline) Less(i, j int) bool {
	if s[i].Deadline == "none" && s[j].Deadline == "none" {
		return s[i].Position < s[j].Position
	} else if s[i].Deadline == "none" {
		return false
	} else if s[j].Deadline == "none" {
		return true
	} else {

		time1, _ := time.Parse(time.UnixDate, s[i].Deadline)
		time2, _ := time.Parse(time.UnixDate, s[j].Deadline)

		if time1.Equal(time2) {
			if s[i].Prioirty == s[j].Prioirty {

				return s[i].Position < s[j].Position
			}
			return s[i].Prioirty > s[j].Prioirty
		}
		return time1.Before(time2)
	}
}
