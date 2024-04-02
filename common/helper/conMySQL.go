package helper

import (
	"blog_backend/internal/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConMySQL(mySQLConf config.MySQLConf) (db *gorm.DB, err error) {

	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&writeTimeout=%s",
		mySQLConf.Username,
		mySQLConf.Password,
		mySQLConf.Host,
		mySQLConf.Port,
		mySQLConf.Database,
		mySQLConf.Charset,
		mySQLConf.Timeout,
	)

	db, err = gorm.Open(mysql.Open(args), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err = AutoMigrate(db); err != nil {
		return db, err
	}

	return db, err
}

func AutoMigrate(db *gorm.DB) (err error) {

	err = db.AutoMigrate(
	//&models.BaseComment{},
	//&models.User{},
	//&models.Article{},
	//&models.Tag{},
	//&models.Image{},
	//&models.Upload{},
	)

	return err
}
