package main

import (
	"social"
	"social/db"
	"social/pkg/common/logging"
	"social/pkg/service"
	"social/pkg/service/connect"
	"social/proto/gen/social/v1/socialv1connect"

	"social/internal/database"
	"social/internal/manager"
	"social/internal/server"
	"social/internal/store/databasedriver"
	"social/internal/store/s3"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"

	"github.com/spf13/cobra"
)

func NewServerCommand() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:  "server",
		RunE: runServer,
	}
	return serverCmd
}

func runServer(cmd *cobra.Command, _ []string) error {
	var config social.Config
	ctx := cmd.Context()

	log := logging.NewLogger("social")
	logging.NewContextLogger(ctx, log)

	err := db.MigrateDB(config.Dsn())
	if err != nil {
		return err
	}

	db, err := database.New(cmd.Context(), config)
	if err != nil {
		return err
	}

	awsConfig, err := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(config.AWSRegion))
	if err != nil {
		return err
	}

	s3Store := s3.New(&config, awsConfig)

	store := databasedriver.New(db)
	manager := manager.NewManager(store, s3Store)
	handler := server.NewServer(manager)

	svc, err := service.New(config.Config, connect.NewServer(config.Config, socialv1connect.NewSocialServiceHandler, socialv1connect.SocialServiceHandler(handler)))
	if err != nil {
		return err
	}
	return svc.Run(cmd.Context())
}
