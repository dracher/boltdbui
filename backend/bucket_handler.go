package backend

import (
	"fmt"

	iris "gopkg.in/kataras/iris.v6"
)

// ListBucketByNameHandler is
func ListBucketByNameHandler(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, listBucketByName(ctx.Param("name")))
}

// ListBucketsHandler is
func ListBucketsHandler(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, listAllBuckets())
}

// CreateBucketHandler is
func CreateBucketHandler(ctx *iris.Context) {
	err := createBucket(ctx.Param("name"))
	if err != nil {
		ctx.Redirect("/api/v1/bucket", iris.StatusInternalServerError)
	} else {
		ctx.Redirect("/api/v1/bucket", iris.StatusOK)
	}
}

// DeleteBucketHandler is
func DeleteBucketHandler(ctx *iris.Context) {
	err := deleteBucket(ctx.Param("name"))
	if err != nil {
		ctx.JSON(iris.StatusInternalServerError,
			fmt.Sprintf("bucket %s can't delete", ctx.Param("name")))
	} else {
		ctx.JSON(iris.StatusOK,
			fmt.Sprintf("bucket %s deleted", ctx.Param("name")))
	}
}
