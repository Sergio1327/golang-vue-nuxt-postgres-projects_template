package main

import (
	"os"
	"projects_template/bimport"
	"projects_template/config"
	"projects_template/external"
	"projects_template/internal/transaction"
	"projects_template/rimport"
	"projects_template/tools/logger"
	"projects_template/tools/pgdb"
	"projects_template/uimport"
)

const (
	module = "project_template"
)

var (
	version string = os.Getenv("VERSION")
)

func main() {
	log := logger.NewNoFileLogger(module)
	log.Infoln("version", version)

	config, err := config.NewConfig(os.Getenv("CONF_PATH"))
	if err != nil {
		log.Fatal(err)
	}
	log.Infoln(config.Template.TmplString)
	db := pgdb.NewPostgresDB(config.PostgresURL())
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	sm := transaction.NewSQLSessionManager(db)
	ri := rimport.NewRepositoryImports(sm)

	dblog := logger.NewDBLog(module, ri)

	bi := bimport.NewEmptyBridge()

	ui := uimport.NewUsecaseImports(log, dblog, ri, bi, sm)
	bi.InitBridge(
		ui.Usecase.Template,
	)

	cron := external.NewCron(log, dblog, ui)
	cron.Run()
}
