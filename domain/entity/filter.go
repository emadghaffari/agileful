package entity

type Filter struct{
	Query string
	Order string
	Limit int
	Offset int
}

var QueryValidate = map[string]int{
	"SELECT" :1,
	"INSERT" :2,
	"UPDATE" :3,
	"DELETE" :4,
}

var OrderByValidate = map[string]int{
	"desc":1,
	"asc":2,
}