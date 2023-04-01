package core

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/shadowscatcher/shodan"
	"github.com/shadowscatcher/shodan/search"
)

var ips = make([]string, 0)

func (c *Core) Search() {

	for _, i := range c.Config.SearchSettings.Query.Text {
		go c.search(1, i)
		time.Sleep(5 * time.Second)
	}

	fmt.Printf(color.HiBlackString("\nOutput file: ")+color.HiGreenString("%s\n"), c.Outfile)

}

func (c *Core) search(offset int, text string) {

	var q = c.Config.SearchSettings.Query

	client, err := shodan.GetClient(c.Config.ShodanApiKey, http.DefaultClient, true)

	if err != nil {
		fmt.Println("homo search: invalid API key")
		return
	}

	searchConfig := search.Params{
		Page: uint(offset),
		Query: search.Query{
			Text: text,
		},
	}

	if !c.ShodanFree {
		searchConfig.Query.OS = q.Os
		searchConfig.Query.Hash = q.Hash
		searchConfig.Query.HasSSL = q.HasSSL
		searchConfig.Query.Port = q.Port
		searchConfig.Query.Net = q.Net
		searchConfig.Query.Org = q.Org
		searchConfig.Query.Tag = q.Tag
		searchConfig.Query.Region = q.Region
	}

	result, err := client.Search(c.ctx, searchConfig)
	if err != nil {
		fmt.Printf("\nhomo search: " + err.Error())
	}

	for _, match := range result.Matches {
		fmt.Printf("[ %s ] %s\n", color.HiGreenString(time.Now().Format("15:04:05")), color.HiBlackString(match.IpAndPort()))
		ips = append(ips, match.IpAndPort())

	}

	if c.Outfile != "null" {
		file, _ := os.OpenFile(c.Outfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
		for _, i := range ips {
			file.Write([]byte(i + "\n"))
		}
		defer file.Close()
	}
}
func CheckPlan(api string, ctx context.Context) (isFree bool, err error) {

	client, err := shodan.GetClient(api, http.DefaultClient, true)

	if err != nil {
		fmt.Println("homo search: invalid API key")
		return false, err
	}

	result, err := client.ApiInfo(ctx)

	if err != nil {
		fmt.Println("homo search: Error: " + err.Error())
		return false, err
	}

	if result.Plan == "oss" {
		isFree = true
	} else {
		isFree = false
	}

	return isFree, nil
}
