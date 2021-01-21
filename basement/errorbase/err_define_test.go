package errorbase

type ServerError interface {
	error
	RequestId() string
}

type UserError struct {
	Err string
	Id  string
}

func (ue *UserError) Error() string {
	return ue.Err
}

// 方法不能和属性同名
func (ue *UserError) RequestId() string {
	return ue.Id
}
