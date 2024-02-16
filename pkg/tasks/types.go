package tasks

import (
	"github.com/edgedb/edgedb-go"
)

type TaskModel struct {
	ID          edgedb.OptionalUUID      `edgedb:"id"`
	Title       string                   `edgedb:"title"`
	Description edgedb.OptionalStr       `edgedb:"description"`
	Due_date    edgedb.OptionalLocalDate `edgedb:"due_date"`
	Created_at  edgedb.OptionalDateTime  `edgedb:"created_at"`
}
