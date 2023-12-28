package model

import (
	"gorm.io/gorm"
	"time"
	companyModel "trackingApp/features/company/model"
)

type User struct {
	ID        string         `json:"id" gorm:"primaryKey; type:varchar(255)"`
	Username  string         `json:"username"  gorm:"type:varchar(100)"`
	Name      string         `json:"name" gorm:"type:varchar(100)"`
	CompanyID string         `json:"company_id" gorm:"default:null"`
	Email     string         `json:"email" gorm:"ype:varchar(100);uniqueIndex"`
	Password  string         `json:"password" gorm:"type:varchar(100)"`
	Role      uint           `json:"role" gorm:"type:tinyint(10)"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`

	Company companyModel.Company ` json:"company" gorm:"foreignKey:CompanyID"`
}

func (u *User) PrepareGive() {
	u.Password = ""
}
