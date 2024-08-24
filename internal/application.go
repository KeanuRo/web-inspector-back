package internal

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Application struct {
	config Config
	db     *pgx.Conn
}

func NewApplication(config Config) *Application {
	app := new(Application)
	app.config = config

	return app
}

func (app *Application) Connect() error {
	dbConf := app.config.Database
	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Dbname)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return err
	}

	app.db = conn
	return nil
}

func (app *Application) CloseConnection() {
	if app.db == nil {
		return
	}

	app.db.Close(context.Background())
}
