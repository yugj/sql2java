package model

type Table struct {
	Name        string
	Description string
	Fields      []TableField
}

type TableField struct {
	Name        string
	Type        string
	NotNull     bool
	Description string
}

type Entity struct {
	Package          string
	TableName        string
	EntityName       string
	ParentEntityName string
	Fields           []EntityField
	Author           string
	Description      string
}

type EntityField struct {
	Name        string
	Type        string
	Description string
	NotNull     bool
}

type GenConfig struct {
	BasePath           string
	BasePackage        string
	ParentEntityName   string
	SystemUser         string // auto get from sys api
	Author             string
	TablePrefix        string // table prefix like t_
	DateType           string // support Data&LocalDateTime, not support yet
	ParentEntityFields string
	OutputFormat       bool
}

// TplData built-in Params
type TplData struct {
	BasePackage       string
	TableName         string
	EntityName        string
	EntityFields      []EntityField
	EntityDescription string
	ParentEntityName  string
	Author            string
	SystemUser        string
	DateTime          string
	BasePath          string
	OutputFormat      bool
}
