package postgresql

import (
	"github.com/jmoiron/sqlx"
	"projects_template/internal/transaction"
)

func SqlxTx(ts transaction.Session) *sqlx.Tx {
	tx := ts.Tx()
	return tx.(*sqlx.Tx)
}
