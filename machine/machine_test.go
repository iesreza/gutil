package machine

import (
	"github.com/iesreza/gutil/log"
	"gutil/machine"
	"testing"
)

func TestMachine(t *testing.T) {

	log.Info("%+v", machine.NetworkConfig())
	log.Info("Unique Machine ID:%s", machine.UniqueID())

}
