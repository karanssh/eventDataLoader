package ingestion

import (
	"encoding/csv"
	"eventDataLoader/config"
	"eventDataLoader/services"
	"io"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

// encoding/csv is very slow on big files. never do readAll unless you dont care about OOM lol
// https://stackoverflow.com/questions/32027590/efficient-read-and-write-csv-in-go
// followed this to read the csv files
// not in use

func processCSV(rc io.Reader) (ch chan []string) {
	ch = make(chan []string, 10)
	go func() {
		r := csv.NewReader(rc)
		if _, err := r.Read(); err != nil { //read header
			log.Fatal(err)
		}
		defer close(ch)
		for {
			rec, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)

			}
			ch <- rec
		}
	}()
	return
}

// read a csv file
// not in use
func readCSVFile() {
	file, err := os.Open(config.Config.IngestFolderLocation + "/output_1.csv")
	if err != nil {
		log.Fatalf("os.Open() Error: %v", err)
	}

	for records := range processCSV(file) {
		log.Print(records)
	}

}

//ImportCSVFile imports a single csv file to database
func ImportCSVFile(DB services.Database, filePath, tableName string) {
	log.Printf("start ImportCSVFile for filePath %s", filePath)
	timeStampCurrent := time.Now()
	log.Printf("time started: %v", timeStampCurrent)
	mysql.RegisterLocalFile(filePath)
	res, err := DB.DatabaseConnection.Exec("LOAD DATA LOCAL INFILE '" + filePath + "' INTO TABLE " + tableName + " FIELDS TERMINATED BY ',' LINES TERMINATED BY '\\n'")
	if err != nil {
		log.Fatalf("importCSVFile load err %v", err)
	}
	val, err := res.RowsAffected()
	if err != nil {
		log.Printf("res.RowsAffected err %v", err)
	}
	mysql.DeregisterLocalFile(filePath)
	log.Printf("LOAD DATA LOCAL INFILE took %s time", time.Since(timeStampCurrent).String())
	log.Printf("Rows affected %d", val)
	log.Printf("finished ImportCSVFile for filePath %s", filePath)

}

//ImportCSVFiles imports a bunch of files from folder to database
func ImportCSVFiles(DB services.Database, folderPath string, tableName string) {
	log.Printf("start ImportCSVFile for folderPath %s", folderPath)
	timeStampCurrent := time.Now()
	log.Printf("time started: %v", timeStampCurrent)
	fileList := fetchFilesFromFolder(folderPath)
	for _, fs := range fileList {
		absolutepath := config.Config.IngestFolderLocation + "/" + fs.Name()
		log.Printf("importing file : %s", absolutepath)
		timeStampStartImportForSingleFile := time.Now()
		ImportCSVFile(DB, absolutepath, tableName)
		log.Printf("ImportCSVFile() took %s time", time.Since(timeStampStartImportForSingleFile).String())

	}
	log.Printf("finished ImportCSVFile() took %s time", time.Since(timeStampCurrent).String())
}
