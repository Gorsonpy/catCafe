// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	cat "github.com/Gorsonpy/catCafe/biz/router/cat"
	membership "github.com/Gorsonpy/catCafe/biz/router/membership"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	cat.Register(r)

	membership.Register(r)

}
