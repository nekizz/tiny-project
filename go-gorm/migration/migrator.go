package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/nekizz/tiny-project/go-gorm/config"
	"github.com/nekizz/tiny-project/go-gorm/migration/versions"
	"gorm.io/driver/mysql"
)

func Run() (respErr error) {
	db, err := config.New(mysql.Open("gorm:secret@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=true&loc=UTC&autocommit=true"))
	if err != nil {
		return err
	}

	defer func() {
		sqlDb, err := db.DB()
		if sqlDb != nil && err == nil {
			sqlDb.Close()
		}
	}()

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID:       "20230808080810",
			Migrate:  versions.Version20230808080810,
			Rollback: versions.Rollback20230808080810,
		},
	})

	return m.Migrate()
}
