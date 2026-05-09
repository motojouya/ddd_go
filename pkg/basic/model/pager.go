package model

type Pager struct {
	Cursor uint `query:"cursor"`
	Limit  uint `query:"limit"`
}

func NewPager(cursor uint, limit uint) Pager {
	return Pager{
		Cursor: cursor,
		Limit:  limit,
	}
}
