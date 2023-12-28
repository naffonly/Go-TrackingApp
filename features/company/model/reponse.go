package model

type CompanyResponse struct {
	ID      string `gorm:"primaryKey; type:varchar(255)"`
	Address string `json:"address" gorm:"varchar(100);Not null"`
	Phone   string `json:"phone" gorm:"varchar(100);Not null"`
	Name    string `json:"name" gorm:"varchar(100);Not null"`
}
