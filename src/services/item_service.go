package services

import (
	"github.com/harlesbayu/bookstore-items-api/src/domain/items"
	"github.com/harlesbayu/bookstore-utils-go/rest_errors"
	"net/http"
)

var (
	ItemService itemsServciceInterface = &itemService{}
)

type itemsServciceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemService struct {}

func (s *itemService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("invalid request bdoy", http.StatusNotImplemented, "not implemented", nil)
}

func (s *itemService) Get(string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me", http.StatusNotImplemented, "not implemented", nil)
}



