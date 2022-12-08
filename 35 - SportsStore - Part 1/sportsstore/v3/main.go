package main

import (
	"platform/http"
	"platform/http/handling"
	"platform/pipeline"
	"platform/pipeline/basic"
	"platform/services"
	"sportsstore/v3/models/repo"
	ssservices "sportsstore/v3/services"
	"sportsstore/v3/store"
	"sync"
)

func registerServices() {
	ssservices.RegisterServices()
	repo.RegisterMemoryRepoService()
}

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		handling.NewRouter(
			handling.HandlerEntry{"", store.ProductHandler{}},
		).AddMethodAlias("/", store.ProductHandler.GetProducts, 1).
			AddMethodAlias("/products", store.ProductHandler.GetProducts, 1),
	)
}

func main() {
	//fmt.Println(os.Getwd())
	registerServices()
	results, err := services.Call(http.Serve, createPipeline())
	if err == nil {
		(results[0].(*sync.WaitGroup)).Wait()
	} else {
		panic(err)
	}
}
