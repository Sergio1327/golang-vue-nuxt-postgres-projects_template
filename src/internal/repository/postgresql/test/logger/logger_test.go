package logger_test

import (
	"fmt"
	"os"
	"projects_template/config"
	"projects_template/internal/repository/postgresql"
	"projects_template/internal/transaction"
	"projects_template/tools/pgdb"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveLog(t *testing.T) {
	r := require.New(t)

	config, err := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NoError(err)

	fmt.Println(config.PostgresURL())

	db := pgdb.NewPostgresDB(config.PostgresURL())
	r.NoError(db.Ping())

	repo := postgresql.NewLogger()

	ts := transaction.NewSQLSession(db)
	err = ts.Start()
	r.NoError(err)
	defer ts.Rollback()

	err = repo.SaveLog(ts, "test")
	r.NoError(err)
}
