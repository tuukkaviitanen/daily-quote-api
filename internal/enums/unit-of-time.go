package enums

type UnitOfTime string

const (
	SECOND    UnitOfTime = "second"
	MINUTE    UnitOfTime = "minute"
	HOUR      UnitOfTime = "hour"
	DAY       UnitOfTime = "day"
	WEEK      UnitOfTime = "week"
	FORTNIGHT UnitOfTime = "fortnight"
	MONTH     UnitOfTime = "month"
	YEAR      UnitOfTime = "year"
)

func (unitOfTime UnitOfTime) IsValidUnitOfTime() bool {
	switch unitOfTime {
	case SECOND, MINUTE, HOUR, DAY, WEEK, FORTNIGHT, MONTH, YEAR:
		return true
	}
	return false
}
