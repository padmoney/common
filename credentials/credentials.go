package credentials

import (
	"encoding/base64"
	"fmt"
)

type Credentials struct {
	url  string
	user string
	pass string
}

func New(url, user, pass string) Credentials {
	return Credentials{
		url:  url,
		user: user,
		pass: pass,
	}
}

func (c Credentials) Token() string {
	up := fmt.Sprintf("%s:%s", c.user, c.pass)
	return base64.StdEncoding.EncodeToString([]byte(up))
}

func (c Credentials) URL() string {
	return c.url
}
