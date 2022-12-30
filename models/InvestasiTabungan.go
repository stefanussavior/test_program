package models

import "gorm.io/gorm"

type Investasi struct {
	gorm.Model
	Nama              string
	JenisKelamin      string
	Usia              int16
	Email             string
	Perokok           string
	Nominal           int32
	PeriodePembayaran string
	MetodeBayar       string
	LamaInvestasi     int16
}
