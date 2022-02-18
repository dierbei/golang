package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"mypro/src/controllers"
)

func main() {
	goft.Ignite().Mount("", controllers.NewProdV2Ctl(), controllers.NewAdminCtl()).Launch()
}
