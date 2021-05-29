package main

import (
	"eventDataLoader/config"
	"eventDataLoader/ingestion"
	"eventDataLoader/services"
	"log"
)

func main() {
	performDataIngestion()
}

func performDataIngestion() {
	log.Printf("start DB connection attempt")
	var db services.Database
	db.CreateConnectionToDatabase("")
	if config.Config.ClearDataBeforeRun {
		db.DropDatabaseDontCareErr(config.Config.DBName)
	}

	db.CreateDatabase(config.Config.DBName)
	db.UseDatabase(config.Config.DBName)
	log.Printf("creating table events...")
	_, err := db.DatabaseConnection.Exec(services.BuildStatementReplaceAllWithTableName(services.TableCreationV2))
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("created table events...")

	log.Printf("start ingestion of data")
	ingestion.ImportCSVFiles(db, config.Config.IngestFolderLocation, config.Config.TableName)
	log.Printf("creating table indexes...")
	db.CreateIndexes() //TODO: check if indexes exist, drop before insert and recreate
	log.Printf("created table indexes...")
	if !config.Config.PersistData {
		db.DropIndexes()
		db.DropDatabase(config.Config.DBName)
	}
	log.Printf("counting total table events...")
	db.GetRowCount()
	db.CloseConnection()

	log.Printf("done ingestion of data")

}
