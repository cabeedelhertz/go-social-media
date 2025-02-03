package database

import (
	"context"
	"log"
	"os"
	"social"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"
)

var logMap = map[string]logger.LogLevel{
	"SILENT": logger.Silent,
	"INFO":   logger.Info,
}

// NewDatabase creates a new database with given config
func New(ctx context.Context, config social.Config) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			// ParameterizedQueries:      paramQueries,     // Don't include params in the SQL log
			Colorful: true, // Disable color
		},
	)
	var db *gorm.DB
	var err error
	for i := 0; i <= 30; i++ {
		db, err = gorm.Open(postgres.Open(config.Dsn()), &gorm.Config{Logger: newLogger})
		if err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	if err != nil {
		return nil, err
	}
	if err = db.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		return nil, err
	}

	origin, err := db.DB()

	if err != nil {
		return nil, err
	}

	origin.SetMaxOpenConns(50)
	origin.SetMaxIdleConns(5)
	origin.SetConnMaxLifetime(time.Duration(86400) * time.Second)

	return db, nil
}
