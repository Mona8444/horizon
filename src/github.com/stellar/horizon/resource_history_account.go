package horizon

import (
	"github.com/jagregory/halgo"
	"github.com/stellar/horizon/db"
	"github.com/stellar/horizon/render/hal"
	"github.com/stellar/horizon/resource"
)

//NewHistoryAccountResourcePage creates a page of HistoryAccountResources
func NewHistoryAccountResourcePage(records []db.HistoryAccountRecord, query db.PageQuery) (hal.Page, error) {
	fmts := "/accounts?order=%s&limit=%d&cursor=%s"
	next, prev, err := query.GetContinuations(records)
	if err != nil {
		return hal.Page{}, err
	}

	resources := make([]interface{}, len(records))
	for i, record := range records {
		var res resource.HistoryAccount
		res.Populate(record)
		resources[i] = res
	}

	return hal.Page{
		Links: halgo.Links{}.
			Self(fmts, query.Order, query.Limit, query.Cursor).
			Link("next", fmts, next.Order, next.Limit, next.Cursor).
			Link("prev", fmts, prev.Order, prev.Limit, prev.Cursor),
		Records: resources,
	}, nil
}
