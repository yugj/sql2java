package parser

import (
	"github.com/xwb1989/sqlparser"
	"regexp"
	"sql-to-java/internal/model"
	"strings"
)

// ParseTable sql -> Table struct
func ParseTable(sql string) (model.Table, error) {

	statement, _ := sqlparser.ParseStrictDDL(sql)
	stmt, _ := statement.(*sqlparser.DDL)
	tableName := stmt.NewName.Name.String()

	//  ENGINE=InnoDB default charset=utf8mb4 comment='设计区域表'
	var options = stmt.TableSpec.Options
	var desc = extractTableDescription(options)

	// table name
	var table = model.Table{Name: tableName, Description: desc, Fields: []model.TableField{}}

	for _, col := range stmt.TableSpec.Columns {
		var field = model.TableField{}

		columnType := col.Type.Type
		field.Type = columnType
		field.Name = col.Name.String()
		field.NotNull = col.Type.NotNull == true

		comment := col.Type.Comment
		if comment != nil {
			field.Description = string(comment.Val)
		}

		// field.EntityDescription = string(col.Type.Comment.Val)
		table.Fields = append(table.Fields, field)
	}

	return table, nil
}

func extractTableDescription(options string) string {

	if len(options) == 0 {
		return ""
	}

	descRegexp := regexp.MustCompile(`\'(.*?)\'`)
	params := descRegexp.FindStringSubmatch(options)
	if len(params) != 0 {
		return strings.ReplaceAll(params[0], "'", "")
	}
	return ""
}
