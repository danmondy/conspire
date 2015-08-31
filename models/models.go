package models

import (
	"time"
)

type Project struct {
	Id            int       "db: id, form: hidden"
	Name          string    "db: name, form: text"
	CustomerEmail string    "db: customer_email, form: email"
	Description   string    "db: desc, form: rtb"
	Lat           string    "db: lat, form: dropdown"
	Lon           string    "db:lon, form: text"
	NotifyList    string    "db:notify_list, form: select_multiple"
	StartDate     time.Time "db:time, form:date"
}

type Component struct {
	Id int "db: id, form: hidden"
	Name string "db: name"	
}

type Task struct {
	Id int "db: id, form: hidden"
}


