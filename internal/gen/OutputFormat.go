package gen

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func outputFormat(baseDir string) {

	var outputDir = baseDir + "/output"
	var outputFormatDir = baseDir + "/output_format"

	filePackageMap, err := getJavaPackagePath(outputDir)
	if err != nil {
		return
	}

	for javaFile, packagePath := range filePackageMap {
		println(javaFile + "->" + packagePath)
		dest := outputFormatDir + "/" + packagePath
		fileCopy(javaFile, dest)
	}
}

// java file path -> packagePath
func getJavaPackagePath(outputDir string) (map[string]string, error) {

	var result = make(map[string]string)
	outputDirFiles, _ := ioutil.ReadDir(outputDir)

	for _, subDir := range outputDirFiles {
		abstSubDir := outputDir + "/" + subDir.Name()

		if !subDir.IsDir() {
			continue
		}

		subDirFiles, _ := ioutil.ReadDir(abstSubDir)
		for _, file := range subDirFiles {
			if strings.HasSuffix(file.Name(), ".java") {
				javaPath := abstSubDir + "/" + file.Name()

				packageName := getPackageName(javaPath)

				if len(packageName) == 0 {
					continue
				}

				result[javaPath] = packageName2Path(packageName)
			}
		}
	}

	return result, nil
}

func packageName2Path(packageName string) string {
	return strings.ReplaceAll(packageName, ".", "/")
}

func getPackageName(javaFile string) string {

	packageLine := readFirstLine(javaFile)
	if len(packageLine) == 0 {
		return ""
	}

	packageLine = strings.TrimPrefix(packageLine, "package")
	packageLine = strings.TrimSuffix(packageLine, ";")
	packageLine = strings.TrimSpace(packageLine)

	return packageLine
}

func readFirstLine(javaFile string) string {

	file, err := os.Open(javaFile)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return ""
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		firstLine := scanner.Text()
		return firstLine
	} else {
		return ""
	}
}

func fileCopy(sourceFile string, dir string) {

	if !existsPath(dir) {
		autoCreateDir(dir)
	}

	_, fileName := filepath.Split(sourceFile)

	input, _ := ioutil.ReadFile(sourceFile)

	destFile := dir + "/" + fileName
	_ = ioutil.WriteFile(destFile, input, 0644)

}
