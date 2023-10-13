package measure_test

import (
	"fmt"
	"os"
	"projects_template/config"
	"projects_template/internal/transaction"
	"projects_template/rimport"
	"projects_template/tools/logger"
	"projects_template/tools/measure"
	"projects_template/tools/pgdb"

	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMeasureFmt(t *testing.T) {
	writer := measure.NewFmtWriter()
	m := measure.NewMeasure(writer)

	m.Start("func A")
	time.Sleep(time.Millisecond * 500)
	m.Stop("func A")

	m.Start("func B")
	time.Sleep(time.Millisecond * 500)
	m.Stop("func B")

	total := m.Result()
	require.Equal(t, 1, int(total.Seconds()))
}

func TestMeasureLorus(t *testing.T) {

	writer := measure.NewLogrusWriter(logger.NewFileLogger("measure"))
	m := measure.NewMeasure(writer)

	t.Run("func A", func(t *testing.T) {
		m.Start("func A")
		time.Sleep(time.Millisecond * 500)
		m.Stop("func A")
	})

	t.Run("func B", func(t *testing.T) {
		m.Start("func B")
		time.Sleep(time.Millisecond * 500)
		m.Stop("func B")
	})

	total := m.Result()
	require.Equal(t, 1, int(total.Seconds()))
}

func TestMeasureLogger(t *testing.T) {

	config, err := config.NewConfig(os.Getenv("CONF_PATH"))
	require.NoError(t, err)

	db := pgdb.NewPostgresDB(config.PostgresURL())
	defer db.Close()

	sm := transaction.NewSQLSessionManager(db)
	ri := rimport.NewRepositoryImports(sm)

	log := logger.NewFileLogger("measure")
	dblog := logger.NewDBLog("measure", ri)

	writer := measure.NewFmtWriter()
	m := measure.NewMeasure(writer)

	m.Start("fmt")
	fmt.Println("testtesttest")
	m.Stop("fmt")

	m.Start("log")
	log.Infoln("testtesttest")
	m.Stop("log")

	m.Start("db")
	dblog.Infoln("testtesttest")
	m.Stop("db")

	m.Result()

}
