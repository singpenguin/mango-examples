package config

import (
	. "controllers"

	"github.com/singpenguin/mango"
)

var Urls = map[string]map[string]mango.Handler{
	"/":       map[string]mango.Handler{"GET": IndexGET},
	"/(\\d+)": map[string]mango.Handler{"GET": IndexGET},

	"/insert": map[string]mango.Handler{"POST": PostHandler},
	"/update": map[string]mango.Handler{"PUT": PutHandler},
	"/delete": map[string]mango.Handler{"DELETE": DeleteHandler},
}
