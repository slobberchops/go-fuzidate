package fuzidate

import (
	"time"
)

func Ymdo(
	year Year,
	month Month,
	day Day,
	offset Offset) Fuzidate {
	return Fuzidate {composeNumber(year, month, day), offset}
}

func Ymo(
	year Year,
	month Month,
	offset Offset) Fuzidate {
	return Ymdo(year, month, 0, offset)
}

func Yo(
	year Year,
	offset Offset) Fuzidate {
	return Ymdo(year, 0, 0, offset)
}

func FromNumberO(number Number, offset Offset) Fuzidate {
	return Fuzidate {number, offset}
}

func Ymd(
	year Year,
	month Month,
	day Day) Fuzidate {
	return Fuzidate{composeNumber(year, month, day), 0}
}

func Ym(
	year Year,
	month Month) Fuzidate {
	return Ymdo(year, month, 0, 0)
}

func Y(
	year Year) Fuzidate {
	return Ymdo(year, 0, 0, 0)
}

func FromNumber(number Number) Fuzidate {
	return FromNumberO(number, 0)
}

func FromTime(t time.Time) Fuzidate {
	return Ymd(Year(t.Year()), Month(t.Month()), Day(t.Day()))
}
