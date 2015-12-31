package config

import (
	. "controllers"

	"github.com/singpenguin/mango"
)

var Urls = map[string]map[string]mango.Handler{
	"/":       map[string]mango.Handler{"GET": IndexGET},
	"/(\\d+)": map[string]mango.Handler{"GET": IndexGET},
}
