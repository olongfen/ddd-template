package initialization

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(InitLog, InitConf)
