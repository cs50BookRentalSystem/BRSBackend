package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	sqliteGo "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"BRSBackend/pkg/models"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(dbPath string) (*Database, error) {

	const CustomDriverName = "sqlite3_extended"
	sql.Register(CustomDriverName,
		&sqliteGo.SQLiteDriver{
			ConnectHook: func(conn *sqliteGo.SQLiteConn) error {
				err := conn.RegisterFunc(
					"gen_random_uuid",
					func(arguments ...interface{}) (string, error) {
						return uuid.New().String(), nil
					},
					true,
				)
				return err
			},
		},
	)

	conn, err := sql.Open(CustomDriverName, dbPath)
	if err != nil {
		panic(err)
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(sqlite.Dialector{
		DriverName: CustomDriverName,
		DSN:        dbPath,
		Conn:       conn,
	}, &gorm.Config{
		Logger:                   newLogger,
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &Database{DB: db}, nil
}

func (db *Database) AutoMigrate() error {
	if err := db.DB.Exec("PRAGMA foreign_keys = ON").Error; err != nil {
		return fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	err := db.DB.AutoMigrate(
		&models.Librarian{},
		&models.Book{},
		&models.Student{},
		&models.Cart{},
		&models.Rent{},
		&models.Session{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migrated successfully")
	return nil
}

func (db *Database) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (db *Database) SeedData() error {
	var count int64
	db.DB.Model(&models.Librarian{}).Count(&count)
	if count > 0 {
		log.Println("Database already seeded")
		return nil
	}

	librarian := models.Librarian{
		User: "admin",
		Pass: []byte("password"),
	}

	if err := db.DB.Create(&librarian).Error; err != nil {
		return fmt.Errorf("failed to seed librarian data: %w", err)
	}

	log.Println("Librarian seeded successfully")
	return nil
}
