package fuzidate

import (
	"fmt"
	"strconv"
)

type Number uint32
type Offset uint
type Year uint16
type Month uint8
type Day uint8
type Int uint64

type Fuzidate struct {
	number Number
	offset Offset
}

type Sequence []Fuzidate

var Unknown Fuzidate = Y(0)

func composeNumber(year Year, month Month, day Day) Number {
	return Number(year) * 10000 + Number(month) * 100 + Number(day)
}

func (f Fuzidate) Number() Number {
	return f.number
}

func (f Fuzidate) Offset() Offset {
	return f.offset
}

func (f Fuzidate) Year() Year {
	return Year(f.number / 10000)
}

func (f Fuzidate) Month() Month {
	return Month(f.number / 100 % 100)
}

func (f Fuzidate) Day() Day {
	return Day(f.number % 100)
}

func (f Fuzidate) Precision() Precision {
	if f.Year() == 0 {
		return None
	}
	if f.Month() == 0 {
		return AtYear
	}
	if f.Day() == 0 {
		return AtMonth
	}
	return AtDay
}

func offsetAsString(o Offset) string {
	if o == 0 {
		return ""
	} else {
		return "+" + strconv.FormatUint(uint64(o), 10)
	}
}

func (f Fuzidate) String() string {
	offsetString := offsetAsString(f.Offset())

	if f.Year() == 0 && f.Month() == 0 && f.Day() == 0 {
		return "0" + offsetString
	}

	if f.Month() == 0 && f.Day() == 0 {
		return fmt.Sprintf("%d%s", f.Year(), offsetString)
	}

	if f.Day() == 0 && f.Year() != 0 {
		return fmt.Sprintf("%d-%02d%s", f.Year(), f.Month(), offsetString)
	}

	return fmt.Sprintf(
		"%d-%02d-%02d%s", f.Year(), f.Month(), f.Day(), offsetString)
}

func (f Fuzidate) Uint() uint64 {
	return uint64(f.Number()) << 32 + uint64(f.Offset())
}
