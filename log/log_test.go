package log_test

import (
	"github.com/iesreza/gutil/log"
	"testing"
)

func TestLog(t *testing.T) {

	data := map[string]int{
		"Age":    25,
		"Weight": 58,
		"Birth":  1870,
	}

	log.Error("This is error")
	log.Notice("This is notice")
	log.Debug("This is debug")
	log.Critical("This is critical")
	log.Info("This is info")
	log.WarningF("This is parametrized warning %d", 100)
	log.Error("Parametrized error %+v", data)

}
