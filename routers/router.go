package routers

import (
	"chain/controllers/basis"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/basis",

		beego.NSNamespace("/blockchain",
			beego.NSInclude(
				&basis.BlockContrillers{},
			),
		),
	)
	beego.AddNamespace(ns)
}
