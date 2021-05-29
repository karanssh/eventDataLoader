package services

import (
	"database/sql"
	"eventDataLoader/config"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	connectionString = fmt.Sprint(config.Config.DBUserName, ":", config.Config.DBPassword, "@",
		config.Config.DBProtocol, "(", config.Config.DBHost, ":", config.Config.DBPort, ")/",
	)
)

type Database struct {
	DatabaseConnection *sql.DB
}

func ConnectToDatabase() {
	log.Printf("connecting to database!")
	conn, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("sql.Open(%s) err: %v", connectionString, err)
	}
	log.Printf("connected to database!")
	conn.Close()
}

func EstablishConnection() {
	log.Printf("connecting to database!")
	var db Database
	db.CreateConnectionToDatabase("")
	db.CreateDatabase(config.Config.DBName)
	db.UseDatabase(config.Config.DBName)
	db.DropDatabase(config.Config.DBName)
	db.CloseConnection()
	log.Printf("connected to database!")

}

//CreateConnectionToDatabase creates connection to a database
func (D *Database) CreateConnectionToDatabase(databaseName string) {
	log.Printf("creating connection to database %s", databaseName)
	connString := connectionString + databaseName

	conn, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatalf("sql.Open(%s) err: %v", connString, err)
	}
	D.DatabaseConnection = conn
	log.Printf("created connection to database %s", databaseName)
}

//close connection to a database
func (D *Database) CloseConnection() {
	log.Printf("closing connection to database")
	err := D.DatabaseConnection.Close()
	if err != nil {
		log.Fatalf("D.DatabaseConnection.Close()(%s) err: %v", connectionString, err)
	}
	log.Printf("closed connection to database")
}

// CreateDatabase creates a database. Remember to Choose it!
func (D *Database) CreateDatabase(databaseName string) {
	log.Printf("creating database : %s ", databaseName)
	_, err := D.DatabaseConnection.Exec(fmt.Sprintf("CREATE DATABASE %s", databaseName))
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("created database : %s ", databaseName)
}

// UseDatabase uses a database.
func (D *Database) UseDatabase(databaseName string) {
	log.Printf("using database : %s ", databaseName)
	_, err := D.DatabaseConnection.Exec(fmt.Sprintf("USE %s", databaseName))
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("set in use database : %s ", databaseName)
}

// CreateDatabase creates a database. Remember to Choose it!
func (D *Database) DropDatabase(databaseName string) {
	log.Printf("dropping database : %s ", databaseName)
	_, err := D.DatabaseConnection.Exec(fmt.Sprintf("DROP DATABASE %s", databaseName))
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("dropped database : %s ", databaseName)
}

// DropDatabaseDontCareErr drops a database but soft logs err
func (D *Database) DropDatabaseDontCareErr(databaseName string) {
	log.Printf("dropping database : %s ", databaseName)
	_, err := D.DatabaseConnection.Exec(fmt.Sprintf("DROP DATABASE %s", databaseName))
	if err != nil {
		log.Print(err.Error())
	}
	log.Printf("dropped database : %s ", databaseName)
}

func (D *Database) CreateIndexes() {
	log.Print("start CreateIndexes()")
	timeStampCurrent := time.Now()
	for _, v := range IndexStatements {
		statement := strings.ReplaceAll(v, "{{tableName}}", config.Config.TableName)
		log.Printf("CreateIndexes() creating index: %s", statement)
		_, err := D.DatabaseConnection.Exec(statement)
		if err != nil {
			log.Print(err.Error())
		}
	}
	log.Printf("CreateIndexes() took %s time", time.Since(timeStampCurrent).String())
	log.Print("finished CreateIndexes()")
}

func (D *Database) DropIndexes() {
	log.Print("start DropIndexes()")
	timeStampCurrent := time.Now()
	for _, v := range IndexDropStatements {
		statement := strings.ReplaceAll(v, "{{tableName}}", config.Config.TableName)
		log.Printf("DropIndexes() creating index: %s", statement)
		_, err := D.DatabaseConnection.Exec(statement)
		if err != nil {
			log.Print(err.Error())
		}
	}
	log.Printf("DropIndexes() took %s time", time.Since(timeStampCurrent).String())
	log.Print("finished DropIndexes()")
}

func (D *Database) GetRowCount() {
	log.Print("start GetRowCount()")
	timeStampCurrent := time.Now()
	rows, err := D.DatabaseConnection.Query(BuildStatementReplaceAllWithTableName(CountTotalEntries))
	if err != nil {
		log.Print(err.Error())
	}
	log.Println("Total count: ", checkCount(rows))
	log.Printf("GetRowCount() took %s time", time.Since(timeStampCurrent).String())
	log.Print("finished GetRowCount()")
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Print(err.Error())
		}
	}
	return count
}
