package dialect

import "reflect"

var dialectMap = map[string]Dialect{}

// Dialect is an interface contains methods that a dialect has to implement
type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

// RegisterDialect register a dialect to the global variable
func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

// Get the dialect from global variable if it exists
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return
}
