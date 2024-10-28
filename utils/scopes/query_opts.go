package scopes

type Direction string

const (
	Ascending  Direction = "ASC"
	Descending Direction = "DESC"
)

type Query_opts struct {
	Page           int
	PageSize       int
	Order          string
	OrderDirection Direction
}
