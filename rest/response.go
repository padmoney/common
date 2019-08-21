package rest

type Response struct {
	Body  []byte
	Code  int
	Error error
}

func (r Response) Ok() bool {
	return r.Code == 200
}
