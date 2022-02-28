package main

import (
	"carbon/common"
	"carbon/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	db := common.InitDB()
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	r := gin.Default()
	r = routers.InitRouter(r)
	panic(r.Run())
}
