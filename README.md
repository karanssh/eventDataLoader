# eventDataLoader

Load the events generated using <https://github.com/karanssh/generateCallerEventData> to a sql database

## How to use the data loader

All the binary settings can be controlled using config json file located in config.json file. Following parameters can be changed:

```json
{
    "DBHost" : "127.0.0.1",
    "DBPort" : "3306",
    "DBProtocol": "tcp",
    "DBUserName" : "root",
    "DBPassword" : "karanTest",
    "TableName": "ingestEvents",
    "IngestFolderLocation": "/home/karan/projects/netcracker/generateEventData/outputData",
    "PersistData": true,
    "ClearDataBeforeRun": true
}

```

persistData flag will ensure that DB remains same after insert, and ClearDataBeforeRun will soft-drop DB and recreate it before run to ensure no items from previous run exist

## Runtime results with 8000000 objects

```text
2021/05/29 12:57:43 start ImportCSVFile for folderPath /home/karan/projects/netcracker/generateEventData/outputData
2021/05/29 12:57:43 time started: 2021-05-29 12:57:43.426349914 +0530 IST m=+0.032756182
2021/05/29 12:57:43 importing file : /home/karan/projects/netcracker/generateEventData/outputData/output_1.csv
2021/05/29 12:57:43 start ImportCSVFile for filePath /home/karan/projects/netcracker/generateEventData/outputData/output_1.csv
2021/05/29 12:57:43 time started: 2021-05-29 12:57:43.426824112 +0530 IST m=+0.033230376
2021/05/29 12:58:57 LOAD DATA LOCAL INFILE took 1m13.67494859s time
2021/05/29 12:58:57 Rows affected 1000000
2021/05/29 12:58:57 finished ImportCSVFile for filePath /home/karan/projects/netcracker/generateEventData/outputData/output_1.csv
2021/05/29 12:58:57 ImportCSVFile() took 1m13.675129369s time
2021/05/29 12:58:57 importing file : /home/karan/projects/netcracker/generateEventData/outputData/output_2.csv
2021/05/29 12:58:57 start ImportCSVFile for filePath /home/karan/projects/netcracker/generateEventData/outputData/output_2.csv
2021/05/29 12:58:57 time started: 2021-05-29 12:58:57.10196275 +0530 IST m=+73.708368964
```

Since import of 1000000 entries takes 1m13.675129369s we can expect 8x-15x time for similar no of entries, so around ~10-30 mins for import. Index is dropped and updated after entries are added which would take some more time.

The application only uses ~250MiB at load, which is about right since the size of each csv is ~150 MB

## Index creation runtime

```text
2021/05/29 13:06:04 creating table indexes...
2021/05/29 13:06:04 start CreateIndexes()
2021/05/29 13:06:04 CreateIndexes() creating index: create index event_called_number_index on ingestEvents (called_number);
2021/05/29 13:06:04 CreateIndexes() creating index: create index event_calling_number_index on ingestEvents (calling_number);
2021/05/29 13:06:04 CreateIndexes() creating index: create index event_event_date_index on ingestEvents (event_date);
2021/05/29 13:06:04 CreateIndexes() creating index: create unique index event_event_ref_uindex on ingestEvents (event_ref);
2021/05/29 13:06:07 CreateIndexes() creating index: create index event_event_type_index on ingestEvents (event_type);
2021/05/29 13:06:09 CreateIndexes() creating index: create index event_location_index on ingestEvents (location);
2021/05/29 13:06:12 CreateIndexes() took 8.636941966s time
2021/05/29 13:06:12 finished CreateIndexes()
2021/05/29 13:06:12 created table indexes...
2021/05/29 13:06:12 counting total table events...
2021/05/29 13:06:12 start GetRowCount()
2021/05/29 13:06:13 Total count:  100001
2021/05/29 13:06:13 GetRowCount() took 289.629207ms time
2021/05/29 13:06:13 finished GetRowCount()
```

Index creation is slow, and takes time, especially when updating large no of objects. I create index after import of data, but if index already existed I would drop and re-created it if I had to do this to some DB with pre-existing data.
