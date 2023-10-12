package template

import (
	"projects_template/tools/sqlnull"
	"time"
)

// CreateParam параметры создания, обязательно окончание Param
type CreateParam struct {
	Name      string           `json:"name"`
	StartDate time.Time        `json:"start_date" db:"start_date"`
	EndDate   sqlnull.NullTime `json:"end_date" db:"end_date"`
}
