package services

const (
	TableCreation = `create table event
	(
	
		event_source     text      not null, 
	
		event_ref        text      not null, 
	
		event_type       integer   not null,
	
		event_date       timestamp not null,
	
		calling_number   integer   not null,
	
		called_number    integer   not null,
	
		location         text      not null,
	
		duration_seconds integer   not null,
	
		attr_1           text,
	
		attr_2           text,
	
		attr_3           text,
	
		attr_4           text,
	
		attr_5           text,
	
		attr_6           text,
	
		attr_7           text,
	
		attr_8           text,
	
		PRIMARY KEY (event_source, event_ref)
	
	);`

	TableCreationV2 = `create table {{tableName}}
	(
	
		event_source     text      not null, 
	
		event_ref        text      not null, 
	
		event_type       integer   not null,
	
		event_date       timestamp not null,
	
		calling_number   integer   not null,
	
		called_number    integer   not null,
	
		location         text      not null,
	
		duration_seconds integer   not null,
	
		attr_1           text,
	
		attr_2           text,
	
		attr_3           text,
	
		attr_4           text,
	
		attr_5           text,
	
		attr_6           text,
	
		attr_7           text,
	
		attr_8           text,
	
		PRIMARY KEY (event_source(10), event_ref(50))
	
	);`
	CountTotalEntries = `SELECT COUNT(*) as count FROM {{tableName}};`
)

var (
	IndexStatements = []string{
		`create index event_called_number_index on {{tableName}} (called_number);`,
		`create index event_calling_number_index on ingestEvents (calling_number);`,
		`create index event_event_date_index on ingestEvents (event_date);`,
		`create unique index event_event_ref_uindex on ingestEvents (event_ref);`,
		`create index event_event_type_index on ingestEvents (event_type);`,
		`create index event_location_index on ingestEvents (location);`,
	}
	IndexDropStatements = []string{
		`DROP INDEX event_called_number_index ON {{tableName}};`,
		`DROP INDEX event_calling_number_index ON {{tableName}};`,
		`DROP INDEX event_event_date_index ON {{tableName}};`,
		`DROP INDEX event_event_ref_uindex ON {{tableName}};`,
		`DROP INDEX event_event_type_index ON {{tableName}};`,
		`DROP INDEX event_location_index ON {{tableName}};`,
	}
)
