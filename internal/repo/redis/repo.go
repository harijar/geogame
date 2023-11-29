package redis

type Tokens interface {
	Get(token string) (int, error)
	Set(token string, id int) error
	Delete(token string) error
}
