package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/cors"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"

	"flag"

	"github.com/boltdb/bolt"
	"github.com/dracher/boltdbui/backend"
)

var dbPath = flag.String("db",
	"/home/dracher/GoProjects/src/github.com/dracher/boltdbui/aqidata.boltdb",
	"database absolute path")

func main() {
	app := iris.New()

	db, err := bolt.Open(*dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	backend.DB = db
	defer db.Close()

	app.Adapt(
		iris.DevLogger(),
		httprouter.New(),
		view.HTML("./views", ".html"),
		cors.New(cors.Options{AllowedOrigins: []string{"*"}}))

	app.Get("/", func(ctx *iris.Context) {
		ctx.Render("index.html", iris.Map{"Title": "Page Title"}, iris.RenderOptions{"gzip": true})
	})

	api := app.Party("/api/v1")
	{
		api.Get("/r", backend.GetValHandler)

		bucket := api.Party("/bucket")
		{
			bucket.Get("/", backend.ListBucketsHandler)

			bucket.Get("/:name", backend.ListBucketByNameHandler)
			bucket.Post("/:name", backend.CreateBucketHandler)
			bucket.Delete("/:name", backend.DeleteBucketHandler)
		}

		db := api.Party("/db")
		{
			db.Get("/status", backend.StatusHandler)
		}
	}

	app.Listen(":6300")
}
