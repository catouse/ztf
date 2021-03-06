package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

type ProductReqPaginate struct {
	domain.PaginateReq
	Keywords string `json:"keywords"`
	Enabled  string `json:"enabled"`
}
