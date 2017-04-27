package backend

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/kataras/iris.v6"
)

// StatusHandler is
func StatusHandler(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, dbStatus())
}

// GetValHandler is
func GetValHandler(ctx *iris.Context) {
	parm := ctx.URLParams()

	log.Debug("start to querying")
	val, _ := getValFromBucket(parm["key"], parm["bucket"])

	ctx.JSON(iris.StatusOK, string(val))
}
