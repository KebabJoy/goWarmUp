package main

import (
	"VSpace/internal/api"
	"VSpace/internal/conns"
	"context"
	"database/sql"
	"github.com/gorilla/mux"
	"strings"
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
		"_development",
	)
	if err != nil {
		logger.Fatal("ERROR OPENING DB CONNECTION", zap.Error(err))
	}

	router := mux.NewRouter()
	apiCon := api.New(
		conns.New(
			dbMain,
		),
		router,
		logger,
	)
	logger.Log(zap.ErrorLevel, "OK!")
	apiCon.ConfigServer()
	logger.Log(zap.ErrorLevel, "OK!")
}

func openDB(ctx context.Context, addr, user, pwd, dbname string) (*sql.DB, error) {
	cnf := mysql.NewConfig()
	cnf.Addr = addr
	cnf.User = user
	cnf.Passwd = pwd
	cnf.DBName = dbname

	cnf.AllowNativePasswords = true
	cnf.CheckConnLiveness = true
	cnf.RejectReadOnly = false
	cnf.MaxAllowedPacket = 0
	cnf.ParseTime = true

	conn, err := mysql.NewConnector(cnf)
	if err != nil {
		return nil, err
	}
	db := sql.OpenDB(conn)
	if err := db.Ping(); err != nil && strings.Contains(err.Error(), "insecure") {
		cnf.TLSConfig = "skip-verify"
		conn, err = mysql.NewConnector(cnf)
		if err != nil {
			return nil, err
		}
		db = sql.OpenDB(conn)
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
