package parser

import (
	"log"
	"regexp"
	"testing"
)

func TestParseTable(t *testing.T) {
	var options = " ENGINE=InnoDB default charset=utf8mb4 comment='空白产品设计区域关联表'"
	descRegexp := regexp.MustCompile(`\'(.*?)\'`)
	params := descRegexp.FindStringSubmatch(options)
	if len(params) != 0 {
		log.Println(params[0])
		return
	}
	log.Println("mismatch")
}
