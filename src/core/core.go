package core

import (
	"context"
	"searcher/src/config"
	"sync"
)

type Core struct {
	Config     *config.Config
	ShodanFree bool

	Outfile string
	ctx     context.Context
	wg      *sync.WaitGroup
}

func Launch(config *config.Config, out string) {
	ctx := context.Background()

	plan, err := CheckPlan(config.ShodanApiKey, ctx)
	if err != nil {
		return
	}
	core := Core{
		Config:     config,
		wg:         &sync.WaitGroup{},
		ctx:        ctx,
		ShodanFree: plan,
		Outfile:    out,
	}

	core.Search()

}
