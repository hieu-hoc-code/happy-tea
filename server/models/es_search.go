package models

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type ElasticSearch struct {
	es *elasticsearch.Client
}

func ConnectES() *elasticsearch.Client {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error when connect to ES database: %s", err)
	}
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error when get response: %s", err)
	}

	defer res.Body.Close()

	return es
}

func NewES(es *elasticsearch.Client) *ElasticSearch {
	return &ElasticSearch{
		es: es,
	}
}

func (es *ElasticSearch) Create(in interface{}) {
	data := in.(Product)
	filterData := map[string]interface{}{
		"name":         data.Name,
		"desc":         data.Descrip,
		"price":        data.Price,
		"image":        data.Image,
		"topping_id":   data.Topping_ID,
		"category_id":  data.Category_ID,
		"inventory_id": data.Inventory_ID,
		"discount_id":  data.Discount_ID,
	}

	var id string = strconv.FormatUint(uint64(data.Id), 10)
	line, _ := json.Marshal(filterData)
	req := esapi.IndexRequest{
		Index:        "ecommerce",
		DocumentType: "_doc",
		DocumentID:   id,
		Body:         bytes.NewReader(line),
	}

	req.Do(context.Background(), es.es)
}

func (es *ElasticSearch) Update(in interface{}) {
	data := in.(Product)
	filterData := map[string]interface{}{
		"name":        data.Name,
		"desc":        data.Descrip,
		"price":       data.Price,
		"image":       data.Image,
		"category_id": data.Category_ID,
		"inventory":   data.Inventory_ID,
		"discount_id": data.Discount_ID,
	}

	var id string = strconv.FormatUint(uint64(data.Id), 10)
	line, _ := json.Marshal(filterData)
	req := esapi.UpdateRequest{
		Index:      "ecommerce",
		DocumentID: id,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, line))),
	}

	req.Do(context.Background(), es.es)
}

func (es *ElasticSearch) Delete(in uint) {
	var id string = strconv.FormatUint(uint64(in), 10)

	req := esapi.DeleteRequest{
		Index:      "ecommerce",
		DocumentID: id,
	}

	req.Do(context.Background(), es.es)
}

func (es *ElasticSearch) Search(str string) interface{} {

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  str,
				"fields": []string{"name", "desc"},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := es.es.Search(
		es.es.Search.WithContext(context.Background()),
		es.es.Search.WithIndex("ecommerce"),
		es.es.Search.WithBody(&buf),
		es.es.Search.WithTrackTotalHits(true),
		es.es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	var r map[string]interface{}
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the ID and document source for each hit.
	var data []interface{}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"]

		item := map[string]interface{}{
			"id":          hit.(map[string]interface{})["_id"],
			"name":        source.(map[string]interface{})["name"],
			"desc":        source.(map[string]interface{})["desc"],
			"price":       source.(map[string]interface{})["price"],
			"image":       source.(map[string]interface{})["image"],
			"category_id": source.(map[string]interface{})["category_id"],
			"inventory":   source.(map[string]interface{})["inventory"],
			"discount_id": source.(map[string]interface{})["discount_id"],
		}
		data = append(data, item)
	}

	return data
}
