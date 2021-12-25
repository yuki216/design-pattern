package service


type Repository interface {
	InsertArea(param1 int64, param2 int64, _type string)	(err error)
}

type Service interface {
	InsertArea(param1 int64, param2 int64, _type string)	(err error)
}
