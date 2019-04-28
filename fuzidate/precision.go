package fuzidate

type Precision uint8

const (
	None Precision = 0 + iota
	AtYear
	AtMonth
	AtDay
)

var precisionStrings = [...]string {
	"None",
	"AtYear",
	"AtMonth",
	"AtDay",
}

func (p Precision) String() string {
	return precisionStrings[p]
}
