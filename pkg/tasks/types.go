package tasks

import (
	"github.com/edgedb/edgedb-go"
)

type TaskModel struct {
	ID          edgedb.OptionalUUID      `edgedb:"id" json:"id"`
	Title       string                   `edgedb:"title" json:"title" binding:"required"`
	Description edgedb.OptionalStr       `edgedb:"description" json:"description" binding:"required"`
	Due_date    edgedb.OptionalLocalDate `edgedb:"due_date" json:"due_date" binding:"required"`
	Created_at  edgedb.OptionalDateTime  `edgedb:"created_at" json:"created_at"`
}
