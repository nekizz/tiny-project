package versions

import (
	"github.com/nekizz/tiny-project/go-gorm/model"
	"gorm.io/gorm"
)

func Version20230808080810(tx *gorm.DB) error {

	type Company struct {
		gorm.Model
		Name          string `gorm:"type:varchar(128)"`
		ContactNumber string `gorm:"type:varchar(128)"`
		Description   string `gorm:"type:TEXT"`
	}

	type Task struct {
		gorm.Model
		Name        string `gorm:"type:varchar(128)"`
		Description string `gorm:"type:TEXT"`
	}

	type User struct {
		gorm.Model
		Name      string  `gorm:"type:varchar(32)"`
		Email     string  `gorm:"type:varchar(32)"`
		Gender    string  `gorm:"type:varchar(32)"`
		Age       uint    `gorm:"type:int"`
		Note      string  `gorm:"type:TEXT"`
		CompanyID uint    `gorm:"type:int;default:null"`
		Company   Company `gorm:"foreignKey:CompanyID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
		Task      []Task  `gorm:"many2many:user_tasks;"`
	}

	return tx.AutoMigrate(&User{}, &Company{}, &Task{})
}

func Rollback20230808080810(tx *gorm.DB) error {
	if err := tx.Migrator().DropTable(&model.User{}); err != nil {
		return err
	}

	if err := tx.Migrator().DropTable(&model.Company{}); err != nil {
		return err
	}

	if err := tx.Migrator().DropTable(&model.Task{}); err != nil {
		return err
	}

	return nil
}
