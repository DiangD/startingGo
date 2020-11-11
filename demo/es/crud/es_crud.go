package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"reflect"
	"strconv"
)

var client *elastic.Client

const host = "http://127.0.0.1:9200"

var employees []Employee

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

func init() {
	errorLog := log.New(os.Stdout, "APP_ES", log.LstdFlags)
	var err error
	client, err = elastic.NewClient(elastic.SetErrorLog(errorLog), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	employee := Employee{
		FirstName: "Qiu",
		LastName:  "zhihan",
		Age:       20,
		About:     "humble and hungry",
		Interests: []string{"basketball", "game", "music"},
	}
	employee1 := Employee{
		FirstName: "Li",
		LastName:  "si",
		Age:       26,
		About:     "I Love C and C++",
		Interests: []string{"basketball", "cook"},
	}
	employee2 := Employee{
		FirstName: "Zhang",
		LastName:  "san",
		Age:       28,
		About:     "I love Python and Javascript",
		Interests: []string{"football", "camper"},
	}
	employees = append(employees, employee)
	employees = append(employees, employee1)
	employees = append(employees, employee2)
}

//创建
func create() {
	for i, employee := range employees {
		resp, err := client.Index().Index("demo_crud").Id(strconv.Itoa(i + 1)).BodyJson(employee).Do(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Indexed tweet %s to index s%s, type %s\n", resp.Id, resp.Index, resp.Type)
	}
}

func query() {
	resp, err := client.Search("demo_crud").Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("=======Query All=======")
	printEmployee(resp, err)
	q := elastic.NewQueryStringQuery("first_name:Zhang")

	resp, err = client.Search("demo_crud").Query(q).Do(context.Background())
	fmt.Println("=======Query By FirstName=======")
	printEmployee(resp, err)

	fmt.Println("=======Query By Condition=======")
	b := elastic.NewBoolQuery()
	//Lte <=,Gte>=
	b.Filter(elastic.NewRangeQuery("age").Lte(26))
	resp, err = client.Search("demo_crud").Query(b).Do(context.Background())
	printEmployee(resp, err)

	fmt.Println("=======Query By Match=======")
	//分词
	p := elastic.NewMatchPhraseQuery("about", "Java")
	resp, err = client.Search("demo_crud").Query(p).Do(context.Background())
	printEmployee(resp, err)

	fmt.Println("=======Query By Aggregation=======")
	aggs := elastic.NewTermsAggregation().Field("interests.keyword")
	resp, err = client.Search("demo_crud").Aggregation("all_interests", aggs).Do(context.Background())
	printEmployee(resp, err)

}

func update() {
	resp, err := client.Update().Index("demo_crud").Id("2").Doc(map[string]interface{}{"about": "I love Java and Go!"}).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Update about %s\n", resp.Result)
}

func delete() {
	resp, err := client.Delete().Index("demo_crud").Id("1").Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("delete Qiu %s\n", resp.Result)
}

func findById(id string) {
	resp, err := client.Get().Index("demo_crud").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if resp.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", resp.Id, resp.Version, resp.Index, resp.Type)
	}
}

func page() {
	resp, err := client.Search("demo_crud").Size(2).From(0).Do(context.Background())
	printEmployee(resp, err)
}

//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
}

func main() {
	//create()
	//update()
	//delete()
	query()
	findById("2")
	page()
}
