package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/gchaincl/dotsql"
	"os"
	"flag"
	"log"
	"strconv"
	"strings"
)

func main() {

	var driver = flag.String("driver", os.Getenv("DSQL_DRIVER"), "Specify the sql driver for the db connection. Currently supported: (mysql | postgres).")
	var dsn = flag.String("dsn", os.Getenv("DSQL_DSN"), "Specify the data source as string. [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]")
	var filepath = flag.String("filepath", os.Getenv("DSQL_FILEPATH"), "Specify the path to the file that needs to be run on the database.")
	var statement = flag.String("statement", os.Getenv("DSQL_STATEMENT"), "Specify the statement tag within the sql file to run only a subset of statements.")

	flag.Parse()

	if *dsn == "" || *filepath == "" || *statement == "" {
		flag.Usage()
		os.Exit(1)
	}

	db, err := sql.Open(*driver, *dsn)
	if err != nil {
		log.Fatalf("Error while connecting to database: \n %s", err)
		os.Exit(1)
	}

	dsql, err := dotsql.LoadFromFile(*filepath)
	if err != nil {
		log.Fatalf("Error while opening sql file: \n %s", err)
		os.Exit(1)
	}

	statements := strings.Split(*statement, ",")

	for _, stmt := range statements {
		result, err := dsql.Exec(db, stmt)

		if err != nil {
			log.Fatalf("Error while executing statement: \n %s", err)
			os.Exit(1)
		}

		rowsAffected, err := result.RowsAffected()

		if err != nil {
			rowsAffected = 0
		}

		log.Print("Successfull run! Rows affected: " + strconv.Itoa(int(rowsAffected)))
	}

	defer db.Close()
}



