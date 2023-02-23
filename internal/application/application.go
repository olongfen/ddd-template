package app

import (
	"ddd-template/internal/adapters/repository/db_iface"
	"ddd-template/internal/adapters/store"
	"log"
)

type Application struct {
	Mutations Mutations
	Queries   Queries
	Cleanup   func()
}

type Close struct {
	data  db_iface.DBData
	store store.Store
}

// Mutations 操作变动的数据
type Mutations struct {
}

// Queries 查询的数据
type Queries struct {
}

func SetQueries() Queries {
	return Queries{}
}

func SetMutations() Mutations {
	return Mutations{}
}

func SetClose(data db_iface.DBData, store store.Store) Close {
	return Close{data: data, store: store}
}

func NewApplication(mut Mutations, que Queries, close Close) (*Application, func()) {
	app := &Application{}
	app.Queries = que
	app.Mutations = mut
	app.Cleanup = func() {

		err := close.data.Close()
		log.Println("db data close", err)
		err = close.store.Close()
		log.Println("cache store close", err)
	}
	return app, app.Cleanup
}
