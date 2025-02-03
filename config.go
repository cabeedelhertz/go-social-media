package social

import (
	"fmt"
	"social/pkg/service"
)

func init() {}

type Config struct {
	service.Config     `config:",squash"`
	DBName             string `config:"db_name"`
	DBPassword         string `config:"db_password"`
	DBHost             string `config:"db_host"`
	DBUser             string `config:"db_user"`
	DBPort             int    `config:"db_port"`
	PGSSLMode          string `config:"pgsslmode"`
	AWSRegion          string `config:"aws_region"`
	AWSAccessKeyId     string `config:"aws_access_key_id"`
	AWSSecretAccessKey string `config:"aws_secret_access_key"`
	BucketName         string `config:"bucket_name"`
}

func (c Config) Dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		c.DBHost,
		c.DBUser,
		c.DBPassword,
		c.DBName,
		c.DBPort,
		c.PGSSLMode,
	)
}
