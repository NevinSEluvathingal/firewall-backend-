package models

import (
	"errors"
	"hello/models/db"
)

type Bar struct {
	Jan   int
	Feb   int
	Mar   int
	April int
	May   int
	June  int
	Jul   int
	Aug   int
	Sept  int
	Oct   int
	Nov   int
	Dec   int
}

func Bardetails(email string, ip string) (Bar, error) {
	query := "SELECT jan, feb, march, april, may, june, july, aug, sept, oct, nov, dec FROM users WHERE mail = ? and ip=?"
	row := db.DB.QueryRow(query, email, ip)
	var value Bar
	err := row.Scan(
		&value.Jan, &value.Feb, &value.Mar, &value.April, &value.May, &value.June,
		&value.Jul, &value.Aug, &value.Sept, &value.Oct, &value.Nov, &value.Dec,
	)
	if err != nil {
		return Bar{}, errors.New("value retrieving error")
	}
	return value, nil
}
