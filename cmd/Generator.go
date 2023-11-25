package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"sql-to-java/internal/convert"
	"sql-to-java/internal/gen"
	"sql-to-java/internal/model"
	"sql-to-java/internal/parser"
	"strings"
)

// Execute cmd
// arg0 sql2java
// arg1 table.sql
// arg2 config.json
func Execute() {

	// system config
	username, _ := user.Current()
	var genConfig = model.GenConfig{}
	genConfig.SystemUser = username.Name

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	genConfig.BasePath = dir
	log.Println("base path : " + dir)

	// table.sql
	var dataFile = os.Args[1]
	sql := readFile(dataFile)

	var argsLen = len(os.Args)

	// read local config
	if argsLen > 2 {
		var configFile = os.Args[2]
		configData := readFile(configFile)
		log.Println("config file" + configFile)
		var localConfig model.GenConfig
		json.Unmarshal([]byte(configData), &localConfig)

		// copy config from json file
		genConfig.BasePackage = localConfig.BasePackage
		genConfig.ParentEntityFields = localConfig.ParentEntityFields
		genConfig.ParentEntityName = localConfig.ParentEntityName
		genConfig.Author = localConfig.Author
		genConfig.TablePrefix = localConfig.TablePrefix
		genConfig.DateType = localConfig.DateType

		if len(localConfig.BasePath) > 0 {
			genConfig.BasePath = localConfig.BasePath
		}
	}

	sqlArray := strings.Split(sql, ";")

	for index := range sqlArray {
		var sqlItem = sqlArray[index]
		var trimSql = strings.TrimSpace(sqlItem)
		if len(trimSql) > 0 {
			var table, _ = parser.ParseTable(trimSql)
			showTableProperty(table)
			log.Println("-----------------------")

			entity := convert.TableToEntity(table, genConfig)
			showEntityProperty(entity)

			gen.Generate(entity, genConfig)
		}
	}

}

func readFile(filepath string) string {

	log.Println("read file: " + filepath)
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Println("read input file error, >> "+filepath, err)
	}

	var text = string(data)
	log.Println(filepath + " file text :\n" + text)
	return text
}

func showTableProperty(table model.Table) {
	log.Println(table.Name)
	for i := range table.Fields {
		var field = table.Fields[i]
		log.Println(field.Name + "," + field.Type + "," + field.Description)
	}
}

func showEntityProperty(entity model.Entity) {
	var lines = ""
	for i := range entity.Fields {
		var field = entity.Fields[i]
		lines += field.Name + "," + field.Type + "," + field.Description + "; "
	}

	log.Println(entity.EntityName + " fields : -->> " + lines)
}
