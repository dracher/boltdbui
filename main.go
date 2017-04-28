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
	flag.Parse()

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
		view.HTML("./frontend", ".html").Binary(Asset, AssetNames),
		cors.New(cors.Options{AllowedOrigins: []string{"*"}}))
	app.StaticEmbedded("/static", "./frontend/static", Asset, AssetNames)

	backend.RegisterRouter(app)
	app.Listen(":6300")
}
