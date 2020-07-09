package request

type TopRequest struct {
	Limit int `form:"limit,default=20"`
}
