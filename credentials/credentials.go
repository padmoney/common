package credentials

import "github.com/padmoney/common/token"

type Credentials struct {
	url  string
	key  string
	cost int
}

func New(url, key string, cost int) Credentials {
	return Credentials{
		url:  url,
		key:  key,
		cost: cost,
	}
}

func (c Credentials) NewToken() string {
	return token.Hash(c.key, c.cost)
}

func (c Credentials) URL() string {
	return c.url
}
