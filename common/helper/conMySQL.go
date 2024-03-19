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

	return db, err
}

func AutoMigrate(db *gorm.DB) (err error) {

	err = db.AutoMigrate(
	//&models.User{},
	//&models.Upload{},
	//&models.Exhibition{},
	//&models.Blog{},
	//&models.Comment{},
	//&models.Likes{},
	//&models.Follow{},
	//&models.Tag{},
	//&models.Star{},
	//&models.Article{},
	//&models.Notice{},
	)

	return err
}
