package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"io/ioutil"
	"net/http"
	"time"
)


func SolrSearchByKeyword(keyword string) interface{} {
	start := time.Now()
	response, err := http.Get("http://localhost:8983/solr/tests/select?fl=title%2Cpublish_time%2Cjournal%2Cabstract%2Cauthors&q=title%3A"+keyword+"*&rows=200000&wt=json")

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var data map[string]interface{}
	e := json.Unmarshal([]byte(body), &data)
	if e != nil {
		panic(e)
	}
	//fmt.Println(string(body))
	times := time.Since(start)
	fmt.Println(times)
	return data["response"]
}

func ElasticSearchByKeyword(keyword string) interface{} {
	ctx:=context.Background()
	start := time.Now()
	client,err:=elastic.NewClient()
	if err!=nil {
		panic(err)
	}
	//exists,err:=client.IndexExists("test2").Do(ctx)
	//if err!=nil {
	//	panic(err)
	//}
	termQuery := elastic.NewTermQuery("title",keyword)
	searchResult,e := client.Search().
		Index("test2").
		Query(termQuery).
		From(0).
		Pretty(true).
		Do(ctx)
	if e!=nil {
		panic(e)
	}
	times := time.Since(start)
	fmt.Println(times)
	return searchResult

}



