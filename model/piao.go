package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/ini.v1"
)

type Piao struct {
	ID                  int     `xorm:"not null pk autoincr comment('') INT(11)" json:"id"`
	Code                string  `xorm:"comment('')" json:"code"`
	Name                string  `xorm:"comment('')" json:"name"`
	Price               float64 `xorm:"comment('')" json:"price"`
	MarketCap           float64 `xorm:"comment('')" json:"market_cap"`
	NetProfit           string  `xorm:"comment() json:"net_profit"`
	Industry            string  `xorm:"comment('') json:"industry"`
	Zhuying             string  `xorm:"comment('') json:"zhuying"`
	Description         string  `xorm:"comment('') json:"description"`
	PE                  string  `xorm:"comment() json:"pe"`
	NetprofitJson       string  `xorm:"comment() json:"netprofit_josn"`
	DetailJson          string  `xorm:"comment() json:"detail_json"`
	ProductJson         string  `xorm:"comment() json:"product_json"`
	FinancialReportJson string  `xorm:"comment() json:"financial_report_json"`
	Employees           string  `xorm:"comment() json:"employees"`
	CompanyInfoJson     string  `xorm:"comment() json:"company_info_json"`
}

var DB *gorm.DB

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

	DB, err = gorm.Open("mysql", strConn)
	if err != nil {
		panic(err)
	}
	//defer DB.Close()
}

func (this *Piao) TableName() string {
	return "gupiao"
}

func (row *Piao) Save() {

	//fmt.Print(article)
	//a := Piao{}
	// rQuery := DB.Where("title = ? and 1", article.Title).Find(&a)
	// if rQuery.Error == nil || a.Id > 0 {
	// 	return
	// }
	result := DB.Create(&row)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}
	fmt.Println(row.ID)
}

func (row *Piao) Update() {
	//DB.First(&user)

	/*
		DB.Where("code=?", row.Code).Take(&row)
		row.MarketCap = 4444
		result := DB.Save(row)
	*/

	result := DB.Model(&row).Where("code = ?", row.Code).Save(row)

	// result := DB.Create(&row)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}
	fmt.Println(row)
}

func (this *Piao) UpdateAll() {
	DB.Save(&this)
}

func (this *Piao) UpdateZhuying(row Piao) {
	fmt.Println(row)
	result := DB.Model(&row).Where("code = ?", row.Code).Update("Zhuying", row.Zhuying)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

}

func (this *Piao) UpdateNetProfit(row Piao) {
	fmt.Println(row)
	result := DB.Model(&row).Where("code = ?", row.Code).Update("NetProfit", row.NetProfit)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

}

func (this *Piao) GetAll() (rows []Piao) {
	//rows := []Piao{}
	DB.Find(&rows)
	return rows
	//fmt.Println(rows)

}

func (this *Piao) GetError() (rows []Piao) {
	//DB.Where("brief=?","").Find(&rows)
	DB.Where("code=?","300894").Find(&rows)
	
	return rows
}
