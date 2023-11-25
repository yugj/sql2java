package convert

// Db2JavaType db type 2 java type
func Db2JavaType(dbType string) string {

	switch dbType {
	case "int":
		return "Integer"
	case "tinyint":
		return "Integer"
	case "bigint":
		return "Long"
	case "decimal":
		return "BigDecimal"
	case "varchar":
		return "String"
	case "text":
		return "String"
	case "datetime":
		return "Date"
	default:
		return "Object"
	}
}
