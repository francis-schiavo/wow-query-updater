package connections

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"log"
)

var dbConn *pg.DB
var ReportingMode bool

type dbLogger struct { }

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

// Connect is used to create the Postgres connection pool
func Connect(username string, password string, db string, poolSize int, schema string, hostname string, port int) {
	createSchemaStatement := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS \"%s\";", schema)
	useSchemaStatement := fmt.Sprintf("SET SEARCH_PATH = \"%s\";", schema)
	dbConn = pg.Connect(&pg.Options{
		Addr: fmt.Sprintf("%s:%d", hostname, port),
		User:     username,
		Password: password,
		Database: db,
		PoolSize: poolSize,
		OnConnect: func(ctx context.Context, conn *pg.Conn) error {
			_, err := conn.Exec(createSchemaStatement)
			if err != nil {
				log.Fatal(err)
			}
			_, err = conn.Exec(useSchemaStatement)
			if err != nil {
				log.Fatal(err)
			}
			return nil
		},
	})
}

// Disconnect is used to disconnect all Postgres connections
func Disconnect() {
	err := dbConn.Close()
	if err != nil {
		log.Panic(err)
	}
}

// GetConn is a safer way to access the connection pool
func GetDBConn() *pg.DB {
	if dbConn != nil {
		return dbConn
	}
	log.Println("database connection not active yet")
	return nil
}
