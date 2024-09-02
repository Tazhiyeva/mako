package fetcher

import (
	model "mako/pkg/models"
	"net/url"
)

type ReviewFetcher interface {
	FetchReview(params url.Values) (*model.Response, error)
}
