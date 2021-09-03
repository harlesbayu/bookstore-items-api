package services

import (
	"github.com/harlesbayu/bookstore-items-api/src/domain/items"
	"github.com/harlesbayu/bookstore-items-api/src/domain/queries"
	"github.com/harlesbayu/bookstore-utils-go/rest_errors"
)

var (
	ItemService itemsServciceInterface = &itemService{}
)

type itemsServciceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
	Search(queries.EsQuery) ([]items.Item, rest_errors.RestErr)
}

type itemService struct {}

func (s *itemService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	if  err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{Id: id}

	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}



