package dal

import (
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	dsn := buildDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return fmt.Errorf("连接数据库失败: %w", err)
	}
	// 获取底层sql.DB对象
	sqldb, err := db.DB()
	if err != nil {
		return fmt.Errorf("获取sql.DB失败: %w", err)
	}
	// 设置连接池参数
	sqldb.SetMaxIdleConns(viper.GetInt("database.max_idle_conns"))
	sqldb.SetMaxOpenConns(viper.GetInt("database.max_open_conns"))
	sqldb.SetConnMaxLifetime(viper.GetDuration("database.conn_max_lifetime"))

	// 测试连接
	if err := sqldb.Ping(); err != nil {
		return fmt.Errorf("数据库连接测试失败: %w", err)
	}

	// 自动迁移数据库表
	if err := autoMigrate(); err != nil {
		return fmt.Errorf("数据库迁移失败: %w", err)
	}

	hlog.Info("数据库连接初始化成功")
	return nil
}

// autoMigrate 自动迁移数据库表
func autoMigrate() error {
	return db.AutoMigrate(

	// 在这里添加其他需要迁移的模型
	)
}

// CloseDB 关闭数据库连接
func Close() {
	sqldb, _ := db.DB()
	_ = sqldb.Close()
}

func buildDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.database"),
		viper.GetString("database.charset"),
		viper.GetBool("database.parse_time"),
		viper.GetString("database.loc"),
	)
}
