package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `grom:"type:varchar(30)"`
	Password string `grom:"type:varchar(32)"`
}

func (u *User) FindUser(db *gorm.DB, name string) error {
	return db.Debug().Where("name = ?", name).Find(&u).Error
}

func (u *User) RedisterAdd(db *gorm.DB) error {
	return db.Debug().Create(&u).Error
}

type Order struct {
	gorm.Model
	Name    string `grom:"type:varchar(30)"`
	Address string `grom:"type:varchar(50)"`
	Num     int    `gorm:"type:int"`
}

func (o *Order) FindOrder(db *gorm.DB, name string) error {
	return db.Debug().Where("id = ?", name).Find(&o).Error
}

func (o *Order) OrderAdd(db *gorm.DB) error {
	return db.Debug().Create(&o).Error
}
func (o *Order) OrderList(db *gorm.DB, id int32) error {
	return db.Debug().Where("id = ?", id).Find(&o).Error
}

type Gwc struct {
	gorm.Model
	Name string `grom:"type:varchar(30)"`
	Num  int    `gorm:"type:int"`
}

func (g *Gwc) GwcAdd(db *gorm.DB) error {
	return db.Debug().Create(&g).Error
}

func (g *Gwc) Findgwc(db *gorm.DB, name string) interface{} {
	return db.Debug().Where("name = ?", name).Find(&g).Error
}
