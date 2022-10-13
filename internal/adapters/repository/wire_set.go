package repository

import "github.com/google/wire"

var Set = wire.NewSet(NewStudentRepository, NewClassRepository, NewData, InitDBConnect)
