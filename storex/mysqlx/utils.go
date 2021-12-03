package mysqlx

import "strings"

func columnStandard(column string) string {
	return "`" + strings.Trim(column, "`") + "`"
}
