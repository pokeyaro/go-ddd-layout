package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"server/infrastructure/persistence/user/po"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Database table name prefix
const (
	dbTablePrefix = "t_"
)

// Database parameter settings
const (
	sqlBatchSize    = 1000
	maxIdleConns    = 10
	maxOpenConns    = 100
	connMaxLifetime = 3600 * time.Second
)

type DBConfig struct {
	DB                       *gorm.DB
	host                     string
	port                     int
	dbUser, dbPasswd, dbName string
}

// NewDBConfig creates a new DBConfig instance with the given parameters.
func NewDBConfig(host string, port int, user string, passwd string, DBName string) *DBConfig {
	return &DBConfig{host: host, port: port, dbUser: user, dbPasswd: passwd, dbName: DBName}
}

// InitDB initializes the database connection, pooling, and table migration.
func InitDB(host string, port int, user, passwd, DBName string) *gorm.DB {
	return NewDBConfig(host, port, user, passwd, DBName).
		connect().
		pool().
		migrateTables().
		DB
}

// dbOptions returns the GORM database options.
func (c *DBConfig) dbOptions() gorm.Option {
	// Set the GORM global logger
	logConfig := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // IO writer
		logger.Config{
			SlowThreshold:             time.Minute, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error
			Colorful:                  true,        // Enable colorful printing
		},
	)

	dbConfig := &gorm.Config{
		// Set table naming strategy
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   dbTablePrefix, // Add a common prefix
			SingularTable: true,          // Use singular form
		},
		// Set logger
		Logger: logConfig,
		// If your project doesn't use transactions, you can set it to improve performance by more than 30%
		SkipDefaultTransaction: false,
		// Set the maximum number of batch created items for a single SQL statement
		CreateBatchSize: sqlBatchSize,
	}

	return dbConfig
}

// dsn returns the Data Source Name (DSN) for the database connection.
func (c *DBConfig) dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.dbUser, c.dbPasswd, c.host, c.port, c.dbName)
}

// connect establishes the database connection.
func (c *DBConfig) connect() *DBConfig {
	db, err := gorm.Open(mysql.Open(c.dsn()), c.dbOptions())
	if err != nil {
		log.Fatal("Init mysql pool failed...")
	}
	log.Println("Connected to MySQL!")

	c.DB = db

	return c
}

// pool configures the database connection pooling settings.
func (c *DBConfig) pool() *DBConfig {
	sqlDB, _ := c.DB.DB()
	// SetMaxIdleConns sets the maximum number of idle connections in the connection pool
	sqlDB.SetMaxIdleConns(maxIdleConns)
	// SetMaxOpenConns sets the maximum number of open connections to the database
	sqlDB.SetMaxOpenConns(maxOpenConns)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	return c
}

// migrateTables performs table migration for the necessary models.
func (c *DBConfig) migrateTables() *DBConfig {
	err := c.DB.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	if err != nil {
		log.Fatalln("Migrate mysql failed...")
	}

	return c
}
