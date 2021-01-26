package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/ini.v1"
)

type Hangye struct {
	ID            int    `xorm:"not null pk autoincr comment('') INT(11)" json:"id"`
	DeveloperName string `xorm:"comment('')" json:"code"`
	AppName       string `xorm:"comment('')" json:"name"`
}

var db *gorm.DB

func init() {
	cfg, err := ini.Load("./app.conf")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	username := cfg.Section("mysql").Key("username").String()
	password := cfg.Section("mysql").Key("password").String()
	database := cfg.Section("mysql").Key("database").String()
	host := cfg.Section("mysql").Key("host").String()
	port := cfg.Section("mysql").Key("port").String()
	strConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)

	db, err = gorm.Open("mysql", strConn)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
}

func (this *Hangye) TableName() string {
	return "hangye"
}

/*
func (this *Hangye) Save(row Hangye) {

	result := db.Create(&row)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}
	fmt.Println(row.ID)
}

func (this *Hangye) Update(row Hangye) {

		result := db.Model(&row).Where("code = ?", row.Code).Update("MarketCap", row.MarketCap)

		// result := db.Create(&row)
		if result.Error != nil {
			fmt.Println(result.Error)
			return
		}
		fmt.Println(row)

}

func (this *Hangye) UpdateAll() {
	db.Save(&this)
}

func (this *Hangye) UpdateZhuying(row Hangye) {
	fmt.Println(row)
	result := db.Model(&row).Where("code = ?", row.Code).Update("Zhuying", row.Zhuying)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

}

func (this *Hangye) UpdateNetProfit(row Hangye) {
	fmt.Println(row)
	result := db.Model(&row).Where("code = ?", row.Code).Update("NetProfit", row.NetProfit)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

}

func (this *Hangye) GetAll() (rows []Hangye) {
	db.Find(&rows)
	return rows

}
*/
