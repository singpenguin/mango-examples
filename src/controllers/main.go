package controllers

import (
	"github.com/singpenguin/mango"
)

type Person struct {
	Name string
	Age  int
}

func IndexGET(ctx *mango.HTTPRequest) {
	//ctx.Write("Hello World")
	//return
	p := Person{"singpenguin", 27}
	ctx.SetHeader("x-test", "test")
	//ctx.SetCookie("uname", "singpenguin", 30)
	ctx.SetSecureCookie("session", "is_login", 600)
	ctx.Render("index.html", p)
}
