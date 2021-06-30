package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	c "lru-cache/src/cache"
)

const calculatorServer = "http://localhost:8080/calc"
const cacheSize = 2

func bodyToString(body io.Reader) string {
	jsonBytes, err := io.ReadAll(body)

	if err != nil {
		log.Fatal(err)
	}

	return string(jsonBytes)
}

func queryCache(operation string, cache *c.LRUCache) (*c.CacheNode, bool) {
	if node, ok := cache.CacheMap[operation]; ok {
		return node, true
	}

	return &c.CacheNode{}, false
}

func main() {
	cache := c.CreateCache(cacheSize)

	http.HandleFunc("/calc", func(res http.ResponseWriter, req *http.Request) {
		operation := bodyToString(req.Body)
		node, cached := queryCache(operation, cache)

		if cached {
			node.Next.Previous = node.Previous
			node.Previous.Next = node.Next
			cache.Enqueue(node)
			fmt.Fprintln(res, "Cached:"+node.Value)
		} else {
			calcResponse, _ := http.Post(calculatorServer, "application/json", bytes.NewBufferString(operation))
			answer := bodyToString(calcResponse.Body)

			cache.AddNewNode(operation, answer)
			fmt.Fprintln(res, "Uncached: "+answer)
		}
	})

	log.Panic(http.ListenAndServe(":8081", nil))
}
