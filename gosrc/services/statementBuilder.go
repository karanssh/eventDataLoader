package services

import (
	"eventDataLoader/config"
	"strings"
)

//BuildStatementReplaceAllWithTableName will replace all {{tableName}} with table name from config file
func BuildStatementReplaceAllWithTableName(inputString string) string {

	return strings.ReplaceAll(inputString, "{{tableName}}", config.Config.TableName)
}
