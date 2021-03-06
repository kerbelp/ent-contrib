// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package todopulid

import (
	"github.com/facebookincubator/ent-contrib/entgql/internal/todopulid/ent/schema/pulid"
	"github.com/facebookincubator/ent-contrib/entgql/internal/todopulid/ent/todo"
)

type TodoInput struct {
	Status   todo.Status `json:"status"`
	Priority *int        `json:"priority"`
	Text     string      `json:"text"`
	Parent   *pulid.ID   `json:"parent"`
}
