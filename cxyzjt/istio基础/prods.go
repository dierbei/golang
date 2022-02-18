package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"mypro/src/controllers"
)


// 商品API
func main() {
   goft.Ignite().Mount("",controllers.NewProdCtl(), controllers.NewAdminCtl()).Launch()
}
