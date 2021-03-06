package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

type TestCaseReqPaginate struct {
	domain.PaginateReq
	Name string `json:"name"`

	Keywords string `json:"keywords"`
	Enabled  string `json:"enabled"`
}
