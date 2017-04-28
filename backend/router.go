package backend

import (
	"gopkg.in/kataras/iris.v6"
)

// RegisterRouter register all project routes
func RegisterRouter(app *iris.Framework) {
	app.Get("/", func(ctx *iris.Context) {
		ctx.Render("index.html", iris.Map{})
	})

	api := app.Party("/api/v1")
	{
		api.Get("/r/:bucket/:key", GetValHandler)
		api.Get("/prefix/:bucket/:prefix", PrefixScanHandler)
		api.Put("/w", PutValHander)
		api.Get("/d/:bucket/:key", DeleteKeyHandler)

		bucket := api.Party("/bucket")
		{
			bucket.Get("/", ListBucketsHandler)

			bucket.Get("/:name", ListBucketByNameHandler)
			bucket.Post("/:name", CreateBucketHandler)
			bucket.Delete("/:name", DeleteBucketHandler)
		}

		db := api.Party("/db")
		{
			db.Get("/status", StatusHandler)
		}
	}
}
