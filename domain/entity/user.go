package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id    uuid.UUID `gorm:"column:Id;type:uuid;primaryKey"`
	Name  string    `gorm:"column:Name;type:string;primaryKey"`
	Email string    `gorm:"column:Email;type:string;uniqueIndex"`
}

// fungsi digunakan untuk membuat logic tambahan sebelum kita menyimpan
// data ke database. Dicontoh ini saya hanya mencontohkan jika developer
// lupa menambahkan data Id, maka disini Id akan diisi dengan uuid.New()
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Id == uuid.Nil {
		u.Id = uuid.New()
	}
	return
}
