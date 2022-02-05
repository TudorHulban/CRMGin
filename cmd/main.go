package main

import (
	"fmt"
	"os"

	"github.com/TudorHulban/GinCRM/cmd/setup"

	"github.com/TudorHulban/GinCRM/pkg/httpinterface"
	"github.com/TudorHulban/GinCRM/pkg/ostop"
	"github.com/TudorHulban/GinCRM/pkg/persistence/cgorm"
	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "version" {
		fmt.Println(setup.Version) // fmt used instead of log for nicer output.
		os.Exit(0)
	}

	ctx, cancel := ostop.NewOSCancellableCtx()
	defer cancel()

	// app logger - sets application log level
	appLogger := log.New(log.DEBUG, os.Stderr, true)

	// clean up persistence
	setup.CleanRDBMS()

	// populate persistence
	cgorm.MigrateDBSchema() // creating RDBMS schema
	if errPo := cgorm.PopulateSchemaSecurityRoles(appLogger); errPo != nil {
		appLogger.Info("Error RDBMS schema population: ", errPo)
		os.Exit(ostop.RDBMSPopulationOfSchema)
	}

	// creating an error group to keep dependencies in sync, only Gin dependency now though.
	g, ctx := errgroup.WithContext(ctx)

	// creating HTTP layer
	g.Go(func() error {
		cfg, errConfig := httpinterface.CreateConfig("0.0.0.0:8080", "0.1", appLogger, 3)
		if errConfig != nil {
			return errors.WithMessage(errConfig, "errors creating HTTP Server configuration")
		}
		gin, errCo := httpinterface.NewGinServer(cfg)
		if errCo != nil {
			return errors.WithMessage(errCo, "could not create HTTP Server instance")
		}
		return gin.Run(ctx)
	})

	if errWait := g.Wait(); errWait != nil {
		appLogger.Info("Error group runtime: ", errWait)
		os.Exit(ostop.RUNTIME)
	}
}
