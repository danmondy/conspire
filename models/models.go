package models

import (
	"time"
)

type Project struct {
	Id            string    "db: id, form: text"
	Name          string    "db: name, form: text"
	CustomerEmail string    "db: customer_email, form: email"
	Description   string    "db: desc, form: rtb"
	Lat           string    "db: lat, form: dropdown"
	Lon           string    "db:lon, form: text"
	StartDate     time.Time "db:time, form:date"
}

type Component struct {
}
