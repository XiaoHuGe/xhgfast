package main

import (
	"xhgfast/models"
	"xhgfast/routers"
	"xhgfast/utils/setting"
)

func main() {
	// Init
	setting.InitConfig()
	db := models.InitDB()
	defer db.Close()
	// router
	r := routers.InitRouter()
	r.Run("0.0.0.0:8000")
}
