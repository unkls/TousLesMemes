package models

import (
  "github.com/jinzhu/gorm"
)

type User struct {
  ID        string    `jsonapi:"primary,users,omitempty" gorm:"primary_key"`
  Login     string    `jsonapi:"attr,login,omitempty"`
  Password  string    `jsonapi:"attr,password,omitempty"`
  Nickname  string    `jsonapi:"attr,nickname,omitempty"`
  Score     int       `jsonapi:"attr,score,omitempty"`
  Groups    Groups    `jsonapi:"relation,groups,omitempty"`
}

type Users []*User

func (u *User) Delete(db *gorm.DB) error {
  return db.Table("users").Delete(&u).Error
}

func (u *User) Create(db *gorm.DB) error {
  return db.Create(u).Error
}

func (u *User) Update(db *gorm.DB, update *User) error {
  return db.Model(&u).Updates(update).Error
}

func (u *User) Show(db *gorm.DB) error {
  return db.Table("users").Where("id LIKE ?", u.ID).Scan(&u).Error
}

func (u *Users) Show(db *gorm.DB) error {
  return db.Table("users").Limit(-1).Offset(0).Scan(&u).Error
}

func (u *User) HidePassword() {
  u.Password = ""
}

func (u Users) HidePassword() {

  for key, _ := range u {
    u[key].Password = ""
  }

}
