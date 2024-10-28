package types

type Query_opts struct {
	Page           int
	PageSize       int
	Order          string
	OrderDirection string
}

func NewQueryOpts(page, pageSize int) Query_opts {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return Query_opts{
		Page:           page,
		PageSize:       pageSize,
		Order:          "created_at",
		OrderDirection: "DESC",
	}
}
