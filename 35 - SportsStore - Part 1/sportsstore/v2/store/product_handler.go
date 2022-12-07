package store

import (
	"platform/http/actionresults"
	"sportsstore/v2/models"
)

const pageSize = 4

type ProductHandler struct {
	Repository models.Repository
}

type ProductTemplateContext struct {
	Products []models.Product
}

func (handler ProductHandler) GetProducts() actionresults.ActionResult {
	return actionresults.NewTemplateAction("product_list.html",
		ProductTemplateContext{
			Products: handler.Repository.GetProducts(),
		})
}
