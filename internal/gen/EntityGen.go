package gen

import (
	"io/ioutil"
	"log"
	"os"
	"sql-to-java/internal/convert"
	"sql-to-java/internal/model"
	"strings"
	"text/template"
)

func Generate(entity model.Entity, config model.GenConfig) {

	// tpl abst path
	var tplBase = config.BasePath + "/templates/"

	// register custom functions
	funcMap := template.FuncMap{
		"toLower": toLower,
		"toUpper": toUpper,
	}

	tpl, err := template.New("sql2java").Funcs(funcMap).ParseGlob(tplBase + "/*")
	if err != nil {
		log.Println("create templates failed, err:", err)
		return
	}

	files, _ := ioutil.ReadDir(tplBase)
	for _, f := range files {

		tplGenerate(convert.ToTplData(entity, config), tpl, f.Name())
	}

	// format output file in need
	if config.OutputFormat {
		outputFormat(config.BasePath)
	}

	log.Println("end of generate")
}

func tplGenerate(tplData model.TplData, tpl *template.Template, filename string) {

	// name := strings.TrimRight(filename, ".tpl") ？？ request.tpl - > request
	name := filename[:len(filename)-4]

	// auto create dir -> /xxx/output/entity
	createDirIfNeed(tplData.BasePath, "output/"+name)

	// camel name service_impl -> ServiceImpl
	var camelName = convert.SnakeCaseToCamel(name)
	camelName = firstUpper(camelName)

	var destFile = tplData.BasePath + "/output/" + name + "/" + tplData.EntityName + camelName + ".java"

	log.Println("generate: " + destFile)

	file := openFile(destFile)

	// console output
	_ = tpl.ExecuteTemplate(os.Stdout, filename, tplData)
	tplErr := tpl.ExecuteTemplate(file, filename, tplData)

	if tplErr != nil {
		log.Println("generate failed", tplErr)
		return
	}
}

func firstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func toLower(s string) string {
	return strings.ToLower(s)
}

func toUpper(s string) string {
	return strings.ToUpper(s)
}
