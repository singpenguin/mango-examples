package controllers

import (
	"github.com/singpenguin/mango"
)

func IndexGET(ctx *mango.HTTPRequest) {
	ctx.Write("Hello World")
}
