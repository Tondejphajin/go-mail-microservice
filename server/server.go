package main

import (
	"sync"
	"database/sql"
	"github.com/Tondejphajin/go-mail-microservice/jsonapi"
	"github.com/Tondejphajin/go-mail-microservice/mdb"
	"github.com/alexflint/go-arg"
)

var args strct {
	DbPath string `arg:"env:MAILINGLIST_DB"`
	BindJson string `arg:env:MAILINGLIST_BIND_JSON"`
}

func main() {
	arg.MustParse(&args)

	if args.DbPath == "" {
		args.DbPath = "list.db"
	}
	if args.BindJson == "" {
		args.BindJson = ":8080"
	}

	log.Printf("using database '%v\n'", args.DbPath)
	db, err := sql.Open("sqlite3", args.DbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mdb.TryCreate(db)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		log.Printf("starting JSON API server ...\n")
		jsonapi.Serve(db,args.BindJson)
		wg.Done()
	}()

	wg.Wait()

}