package request

type GroupBy string

const (
	Day GroupBy = "day"
	Hour = "hour"
	Minute = "minute"
)

type StatConversionGraphRequest struct {
	GroupBy GroupBy `form:"group_by"`
}
