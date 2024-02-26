package db

import (
	"fmt"
	"gotest/configs"
	"log"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalDb *gorm.DB

func init() {
	fmt.Println("### init db ###")

	var err error
	GlobalDb, err = initDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
}

// 初始化数据库连接
func initDB() (*gorm.DB, error) {
	sysConfig := configs.GetSysConfig()
	// dsn := "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local" // 替换为实际的数据库连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		sysConfig.Database.Username,
		sysConfig.Database.Password,
		sysConfig.Database.Host,
		sysConfig.Database.Port,
		sysConfig.Database.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DSN: ", dsn)
		fmt.Println("Failed to connect to the database:", err)
		return nil, err
	} else {
		fmt.Println("Connected to the database successfully")
	}

	// 自动迁移模式（根据模型结构同步数据库表结构）
	// db.AutoMigrate(&YourModel{}) // 更换为你的具体模型结构体

	return db, nil
}

// 然后在需要的地方可以直接使用 GlobalDb
// type UserRepository struct {
//     DB *gorm.DB
// }

// func NewUserRepository() *UserRepository {
//     return &UserRepository{DB: GlobalDb}
// }

// func (ur UserRepository) GetUserById(id uint) (*User, error) {
//     var user User
//     result := ur.DB.First(&user, id)
//     if result.Error != nil {
//         return nil, result.Error
//     }
//     return &user, nil
// }

func Test() {
	logrus.Debugf("Test for db init")
}
