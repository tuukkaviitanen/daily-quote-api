package enums

type TimeInterval string

const (
	SECONDLY    TimeInterval = "secondly"
	MINUTELY    TimeInterval = "minutely"
	HOURLY      TimeInterval = "hourly"
	DAILY       TimeInterval = "daily"
	WEEKLY      TimeInterval = "weekly"
	FORTNIGHTLY TimeInterval = "fortnightly"
	MONTHLY     TimeInterval = "monthly"
	YEARLY      TimeInterval = "yearly"
)

func (timeInterval TimeInterval) IsValidTimeInterval() bool {
	switch timeInterval {
	case SECONDLY, MINUTELY, HOURLY, DAILY, WEEKLY, FORTNIGHTLY, MONTHLY, YEARLY:
		return true
	}
	return false
}
