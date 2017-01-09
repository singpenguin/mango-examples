package controllers

import (
	"fmt"
	"github.com/singpenguin/mango"
)

type Person struct {
	Name string
	Age  int
}

func IndexGET(ctx *mango.HTTPRequest) {
	//ctx.Write([]byte("Hello World"))
	//return
	p := Person{"singpenguin", 27}
	ctx.SetHeader("x-test", "test")
	//ctx.SetCookie("uname", "singpenguin", 30)
	ctx.SetSecureCookie("session", "is_login", 600)
	ctx.Render("index.html", p)
}

func PostHandler(ctx *mango.HTTPRequest) {
	id := ctx.Params.Get("id")
	fmt.Println(id)
	if id == "" {
		ctx.Write([]byte("id cannot be empty"))
		return
	}
	ctx.Write([]byte("insert ok: " + id))
}

func PutHandler(ctx *mango.HTTPRequest) {
	id := ctx.Params.Get("id")
	if id == "" {
		ctx.Write([]byte("id cannot be empty"))
		return
	}
	ctx.Write([]byte("update ok: " + id))
}

func DeleteHandler(ctx *mango.HTTPRequest) {
	id := ctx.Params.Get("id")
	if id == "" {
		ctx.Write([]byte("id cannot be empty"))
		return
	}
	ctx.Write([]byte("delete ok: " + id))
}
