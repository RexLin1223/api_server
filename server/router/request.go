package router

// IRequestHandler is interface of request handler
type IRequestHandler interface {
	Parse(string, string) (string, error)
	Insert(string) error
	Query(string) (string, error)
	Update(string) error
	Delete(string) error
}
