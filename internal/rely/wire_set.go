package rely

import (
	"github.com/google/wire"
	"github.com/olongfen/toolkit/db_data"
)

var Set = wire.NewSet(NewLogger, db_data.NewTransaction, db_data.NewData, InitDBConnect, InitConfigs)
