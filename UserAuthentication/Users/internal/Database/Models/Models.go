package Models 

import(
	"gorm.io/gorm"
)

type User struct{
	gorm.Model

	IdUser uint
	ProfileID uint
	username string `gorm:"unique;not null"`
	Email string `gorm:"unique;not null"`
	Password string `gorm:"type:char(60);not null`
	Profile  Profile `gorm:"foreignKey:ProfileID"`
	Roles    []Role  `gorm:"many2many:user_roles"`

}


type Profile struct{
	gorm.Model

	IdProfile uint
	FName string `gorm:"not null"`
	LName string `gorm:"not null"`
	Phone string `gorm:"not null"`
	Address string `gorm:"not null"`
	
}


type Role struct{
	gorm.Model

	IdRole uint
	Role string `gorm:"not null"`
}

type Permission struct {
	gorm.Model
	
    ID   uint
    Name string
}