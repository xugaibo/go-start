package enums

type YesNoEnum int

const (
	Yes YesNoEnum = 1
	No  YesNoEnum = 0
)

func (b YesNoEnum) Code() int {
	return int(b)
}
