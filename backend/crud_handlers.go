package backend

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/kataras/iris.v6"
)

// PutValData is
type PutValData struct {
	Bucket string
	Key    string
	Val    string
}

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

// PutValHander is
func PutValHander(ctx *iris.Context) {
	data := PutValData{}
	err := ctx.ReadJSON(&data)
	if err != nil {
		ctx.JSON(iris.StatusInternalServerError, err.Error())
	}
	err = putValToBucket(data.Key, data.Val, data.Bucket)
	if err != nil {
		ctx.JSON(iris.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(iris.StatusOK, "OK")
	}
}

// PrefixScanHandler is
func PrefixScanHandler(ctx *iris.Context) {
	val, _ := prefixScan(ctx.Param("bucket"), ctx.Param("prefix"))

	ctx.JSON(iris.StatusOK, val)
}

// DeleteKeyHandler is
func DeleteKeyHandler(ctx *iris.Context) {
	err := deleteKeyFromBucket(ctx.Param("key"), ctx.Param("bucket"))
	ctx.JSON(iris.StatusOK, err.Error())
}
