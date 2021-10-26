package repository

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateUuid() string {
	return fmt.Sprintf("0x%x", [16]byte(uuid.New()))
}

func GenerateFindManyQuery(tableName string, column string) string {
	return fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", tableName, column)
}

func GenerateFindOneQuery(tableName string, column string) string {
	return fmt.Sprintf("SELECT * FROM %s WHERE %s = ? LIMIT 1", tableName, column)
}
