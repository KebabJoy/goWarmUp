package main

import (
	"Application/internal/api"
	"Application/internal/conns"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

func main() {
	logger := buildLogger()
	ctx := ctxzap.ToContext(context.Background(), logger)

	dbMain, err := openDB(ctx,
		"localhost:3306",
		"root",
		"",
		"zip_dev",
	)
	if err != nil {
		logger.Fatal("ERROR OPENING DB CONNECTION", zap.Error(err))
	}

	apiCon := api.New(
		conns.New(
			dbMain,
		),
		logger,
	)
	logger.Log(zap.ErrorLevel, "OK!")
	apiCon.Config()
	logger.Log(zap.ErrorLevel, "OK!")
}

func openDB(ctx context.Context, addr, user, pwd, dbname string) (*sqlx.DB, error) {
	cnf := &mysql.Config{
		User:                 user,
		Net:                  "tcp",
		Addr:                 addr,
		DBName:               dbname,
		AllowNativePasswords: true,
		RejectReadOnly:       false,
		MaxAllowedPacket:     0,
		ParseTime:            true,
	}

	dsn := cnf.FormatDSN()
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		fmt.Println("Can't connect to MySQL database on", addr)
		return nil, err
	}

	db.SetMaxOpenConns(60)
	db.SetMaxIdleConns(60)
	db.SetConnMaxIdleTime(60 * time.Second)
	return db, db.PingContext(ctx)
}

func buildLogger() *zap.Logger {
	zapCfg := zap.NewProductionConfig()

	logger, err := zapCfg.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
