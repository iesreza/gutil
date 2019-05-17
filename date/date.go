package date

import (
	"strconv"
	"strings"
	"time"
	"github.com/araddon/dateparse"
)

type dt struct {
	t time.Time
}

var formats = map[string]string{
	"date":"2006-01-02",
	"sdate":"2006/01/02",
	"sdatetime":"2006/01/02 15:04:05",
	"datetime":"2006-01-02 15:04:05",
	"usdate":"01/02/2006",
	"usdatetime":"01/02/2006 15:04:05",
	"time":"15:04:05",
	"":"2006-01-02 15:04:05-0700",
}


func Now() *dt{
	date := dt{}
	date.t = time.Now()
	return &date
}

func Parse(input interface{}) (*dt,error) {
	date := dt{}
	var e error
	switch input.(type){
		case int64:
			date.t = time.Unix(input.(int64),0)
		case string:
			i,err := strconv.ParseInt(input.(string),10,64)
			if err == nil{
				date.t = time.Unix(i,0)
			}else{
				date.t, err = dateparse.ParseAny("3/1/2014")
				if e != nil{
					e = err
				}
			}
	}

	return &date,e
}

func Today() *dt{
	date := dt{}
	now := time.Now()
	date.t = time.Date( now.Year() , now.Month(),now.Day(),0,0,0,0,time.UTC )
	return &date
}

func (o *dt)Time() *time.Time  {
	return &o.t
}

func (o *dt)AddSeconds(seconds int) *dt  {
	o.t.Add(time.Second * time.Duration(seconds))
	return o
}

func (o *dt)AddDuration(duration string) *dt  {
	d,err := time.ParseDuration(duration)
	if err == nil{
		o.t.Add(d)
	}
	return o
}


func (o *dt)Format(format string) string  {
	if val, ok := formats[ strings.ToLower(format) ]; ok {
		return o.t.Format(val)
	}
	return o.t.Format(format)
}


func (o *dt)SetFormat(name,format string) *dt  {
	formats[ strings.ToLower(name) ] = format
	return o
}

func (o *dt)String() string {
	return o.Format("datetime")
}