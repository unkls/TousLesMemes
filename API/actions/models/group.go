package models

import (
  "github.com/jinzhu/gorm"
)

type Group struct {
  ID    string    `jsonapi:"primary,groups,omitempty"`
  Name  string    `jsonapi:"attr,name,omitempty"`
  Users Users     `jsonapi:"relation,users,omitempty"`
}

type Groups []*Groups

func (g Group) Delete(db *gorm.DB) error {
  return db.Delete(g).Error
}

func (g Group) Create(db *gorm.DB) error {
  return db.Create(g).Error
}

func (g Group) Update(db *gorm.DB, update Group) error {
  return db.Model(g).Update(update).Error
}

func (g Group) Show(db *gorm.DB) error {
  return db.Where("id LIKE ?", g.ID).Scan(g).Error
}

func (g Groups) Show(db *gorm.DB) error {
  return db.Table("groups").Scan(g).Error
}
