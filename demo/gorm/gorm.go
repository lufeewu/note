package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
	Test  uint
}

type User struct {
	gorm.Model
	UUID string
}

type Order struct {
	ID   uint `gorm:"primary_key"`
	Name string
}

// Animal 列名
type Animal struct {
	AnimalID int64     `gorm:"column:beast_id"`         // 设置列名为`beast_id`
	Birthday time.Time `gorm:"column:day_of_the_beast"` // 设置列名为`day_of_the_beast`
	Age      int64     `gorm:"column:age_of_the_beast"` // 设置列名为`age_of_the_beast`
}

func testGorm(db *gorm.DB) {
	// 1 .create
	// db.AutoMigrate(&Product{}, &User{})
	db.CreateTable(&Animal{})
	// db.DropTable("orders")
	// db.AutoMigrate(&Order{})
	animal := Animal{Birthday: time.Now(), Age: 13}
	r := db.NewRecord(animal)
	logrus.Infoln(r)
	db.Create(&animal)
	r = db.NewRecord(animal)
	logrus.Infoln(r)
	return

	db.Create(&Product{Code: "12121", Price: 1000})
	return

	// 2. query
	var p Product
	db.First(&p, 1)
	db.First(&p, "code  = ? ", "12121")

	// 3. update
	db.Model(&p).Update("Price", 2000)

	// 4. Delete
	db.Delete(&p)
}

func main() {

	cfg := struct {
		Address string
		Port    int
		DbUser  string
		DbPass  string
		DbName  string
	}{}

	app := kingpin.New("gormtest", "gorm test app")
	app.Flag("address", "remote server address").Default("127.0.0.1").StringVar(&cfg.Address)
	app.Flag("port", "remote server port").Default("3306").IntVar(&cfg.Port)
	app.Flag("dbuser", "remote server database user").Default("root").StringVar(&cfg.DbUser)
	app.Flag("dbpass", "remote server database password").Default("password").StringVar(&cfg.DbPass)
	app.Flag("dbname", "remote server database name").Default("dbname").StringVar(&cfg.DbName)

	kingpin.MustParse(app.Parse(os.Args[1:]))

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DbUser, cfg.DbPass, cfg.Address, cfg.Port, cfg.DbName)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		logrus.Errorf("failed to connect remote server, error: %v", err)
	}
	defer db.Close()

	testGorm(db)
}
