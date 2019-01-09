package main

import (
	"strconv"
	"fmt"
	"time"
	"strings"
	"unicode"
	"errors"
)

const (
	INVALID = iota
	SECOND
	MINUTE
	HOUR
	DAY
	MONTH
	YEAR

	TIME_FORMAT = "20060102150405"
)


func shouldIFire(str string, t time.Time, last time.Time) (bool, error){
	var err error
	var cronTime time.Time
	var mm int
	var sec,min,hour,day,year int
	var month time.Month

	f:=func(c rune) bool{
		return !unicode.IsNumber(c)
	}

	fields := strings.FieldsFunc(str, f)

	// reverse
	for i,j := 0, len(fields)-1; i<j; i,j = i+1,j-1{
		fields[i],fields[j] = fields[j],fields[i]
	}

	l:=len(fields)

	if l>INVALID{
		sec,err = strconv.Atoi(fields[0])
	}else{
		return false, errors.New("Bad cron value")
	}
	if err !=nil{
		return false, errors.New("Error")
	}

	if l>SECOND{
		min,err = strconv.Atoi(fields[1])
	}else{
		min = t.Minute()
	}
	if err !=nil{
		return false, errors.New("Error")
	}	

	if l>MINUTE{
		hour, err = strconv.Atoi(fields[2])
	}else{
		hour =t.Hour()
	}
	if err !=nil{
		return false, errors.New("Error")
	}	

	if l>HOUR{
		day, err = strconv.Atoi(fields[3])
	}else{
		day = t.Day()
	}
	if err !=nil{
		return false, errors.New("Error")
	}	

	if l>DAY{
		mm, err = strconv.Atoi(fields[4])
		switch(mm){
		case 1:
			month = time.January
		case 2:
			month = time.February
		case 3:
			month = time.March
		case 4:
			month = time.April
		case 5:
			month = time.May
		case 6:
			month = time.June
		case 7: 
			month = time.July
		case 8:
			month = time.August
		case 9:
			month = time.September
		case 10:
			month = time.October
		case 11:
			month = time.November
		case 12:
			month = time.December
		}
	}else{
		month = t.Month()
	}
	if err !=nil{
		return false, errors.New("Error")
	}	

	if l>MONTH{
		year, err = strconv.Atoi(fields[5])
	}else{
		year = t.Year()
	}
	if err !=nil{
		return false, errors.New("Error")
	}	

	cronTime = time.Date(year, month, day, hour, min, sec, 0, t.Location())

	if t.After(cronTime) && cronTime.After(last) {
		return true, nil
	}

	return false, err
}

func main(){
	var fire bool
	
	location, err := time.LoadLocation("Europe/Vienna")
	if err != nil{
		location,_ = time.LoadLocation("UTC")
	}
	last, err := time.ParseInLocation(TIME_FORMAT, "00010101000000", location)

	for{
		cron := "45"
		now := time.Now()

		fmt.Printf("cron\t:[%s]\n", cron)
		fmt.Printf("now\t:[%s]\n", now.Format(TIME_FORMAT))
		fmt.Printf("last\t:[%s]\n", last.Format(TIME_FORMAT))
		fire, err = shouldIFire(cron, now, last)
		if err != nil{
			return
		}
		fmt.Println(fire)
		if fire {
			last = now
		}
		fmt.Println()
		
		time.Sleep(time.Second*2)
	}
}