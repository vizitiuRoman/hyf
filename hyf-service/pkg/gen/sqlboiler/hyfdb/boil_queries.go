// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to hyf-service re-generated in place and/or deleted at any time.

package hyfdb

import (
	"regexp"

	"github.com/volatiletech/sqlboiler/v4/drivers"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var dialect = drivers.Dialect{
	LQ: 0x22,
	RQ: 0x22,

	UseIndexPlaceholders:    true,
	UseLastInsertID:         false,
	UseSchema:               false,
	UseDefaultKeyword:       true,
	UseAutoColumns:          false,
	UseTopClause:            false,
	UseOutputClause:         false,
	UseCaseWhenExistsClause: false,
}

// This is a dummy variable to prevent unused regexp import error
var _ = &regexp.Regexp{}

// NewQuery initializes a new Query using the passed in QueryMods
func NewQuery(mods ...qm.QueryMod) *queries.Query {
	q := &queries.Query{}
	queries.SetDialect(q, &dialect)
	qm.Apply(q, mods...)

	return q
}
