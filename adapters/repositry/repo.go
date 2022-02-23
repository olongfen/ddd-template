package repositry

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewDemoDependencyImpl)
