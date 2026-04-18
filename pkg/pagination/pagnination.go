package pagination

type Order string

const (
	DefaultLimit = 50
)

type PaginationQuery struct {
	Limit  int
	Offset int
}

type PaginationResult[T any] struct {
	Items   []T
	Limit   int
	Offset  int
	Total   int64
	HasNext bool
	HasPrev bool
}

func (q *PaginationQuery) Normalize() {
	if q.Limit == 0 {
		q.Limit = DefaultLimit
	}
	if q.Offset < 0 {
		q.Offset = 0
	}
}
