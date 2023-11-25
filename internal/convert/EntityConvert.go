package convert

import (
	"sql-to-java/internal/model"
	"strings"
	"time"
)

func TableToEntity(table model.Table, config model.GenConfig) model.Entity {
	var entity = model.Entity{Fields: []model.EntityField{}}
	entity.TableName = table.Name

	var fixTableName = table.Name
	if len(config.TablePrefix) > 0 {
		fixTableName = strings.TrimPrefix(table.Name, config.TablePrefix)
	}

	var camelName = SnakeCaseToCamel(fixTableName)
	var domainName = strings.ToUpper(camelName[:1]) + camelName[1:]

	entity.EntityName = domainName
	entity.Author = config.SystemUser
	entity.Description = table.Description
	entity.ParentEntityName = config.ParentEntityName

	if strings.HasSuffix(entity.Description, "表") {
		// Chinese 3 chars
		entity.Description = entity.Description[:len(entity.Description)-3]
	}

	// field convert
	for i := range table.Fields {
		var tableField = table.Fields[i]
		var entityField = tableFieldToEntityField(tableField)
		entity.Fields = append(entity.Fields, entityField)
	}

	if len(config.ParentEntityFields) > 0 {
		var parentFields = strings.Split(config.ParentEntityFields, ",")
		entity.Fields = removeParentFields(entity.Fields, parentFields)
	}

	return entity
}

func ToTplData(entity model.Entity, config model.GenConfig) model.TplData {
	var tplData = model.TplData{}
	tplData.BasePackage = config.BasePackage
	tplData.TableName = entity.TableName
	tplData.EntityName = entity.EntityName
	tplData.ParentEntityName = config.ParentEntityName
	tplData.EntityFields = entity.Fields
	tplData.Author = entity.Author
	tplData.SystemUser = config.SystemUser
	tplData.EntityDescription = entity.Description
	tplData.DateTime = time.Now().Format("2006-01-02 15:04")
	tplData.BasePath = config.BasePath

	return tplData
}

func removeParentFields(fields []model.EntityField, toRemove []string) []model.EntityField {

	for _, item := range toRemove {
		for i, field := range fields {
			if field.Name == item {
				fields = append(fields[:i], fields[i+1:]...)
			}
		}
	}
	return fields
}

func tableFieldToEntityField(tableField model.TableField) model.EntityField {
	var entityField = model.EntityField{}
	entityField.Name = SnakeCaseToCamel(tableField.Name)
	entityField.Description = tableField.Description
	entityField.NotNull = tableField.NotNull
	var dbType = tableField.Type
	entityField.Type = Db2JavaType(dbType)
	return entityField
}

func SnakeCaseToCamel(str string) string {
	builder := strings.Builder{}
	index := 0
	if str[0] >= 'a' && str[0] <= 'z' {
		builder.WriteByte(str[0] - ('a' - 'A'))
		index = 1
	}
	for i := index; i < len(str); i++ {
		if str[i] == '_' && i+1 < len(str) {
			if str[i+1] >= 'a' && str[i+1] <= 'z' {
				builder.WriteByte(str[i+1] - ('a' - 'A'))
				i++
				continue
			}
		}
		builder.WriteByte(str[i])
	}

	var result = builder.String()
	result = strings.ToLower(result[:1]) + result[1:]
	return result
}
