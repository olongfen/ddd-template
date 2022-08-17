package repository

import (
	"context"
	"ddd-template/internal/common/utils"
	"ddd-template/internal/domain"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"regexp"
	"testing"
	"time"
)

func TestDemoRepo_SayHello(t *testing.T) {
	logger, _ := zap.NewProduction()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	dialector := postgres.New(postgres.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "postgres",
		Conn:       db,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: tableNamePrefix,
		},
		//Logger: utils.New(lg,utils.Config{Colorful: true}),
	})

	if err != nil {
		t.Fatal(err)
	}
	data := NewData(gdb, logger)
	defer db.Close()
	repo := NewDemoDependency(data, logger)
	demo := &domain.Demo{
		Model: utils.Model{
			ID:        1,
			CreatedAt: utils.JSONTime{Time: time.Now()},
		},
		Message: "wwwaaa",
	}
	rows := sqlmock.NewRows([]string{"id", "created_at", "message"}).AddRow(demo.ID, demo.CreatedAt, demo.Message)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "demo_demos" WHERE "demo_demos"."deleted_at" IS NULL AND
"demo_demos"."id" = $1 ORDER BY "demo_demos"."id" LIMIT 1
`)).WithArgs(1).WillReturnRows(rows)
	_demo := new(domain.Demo)
	_demo.ID = 1
	err = repo.Get(context.Background(), _demo)
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, *_demo, *demo)
}
