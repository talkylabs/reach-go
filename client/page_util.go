package client

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Takes a limit on the max number of records to read and a max pageSize and calculates the max number of pages to read.
func ReadLimits(pageSize *int, limit *int) int {
	//don't care about pageSize
	if pageSize == nil {
		if limit == nil {
			//don't care about the limit either
			return 20 //default
		}
		//return the most efficient pageSize
		return min(*limit, 1000)
	} else {
		if limit == nil {
			//we care about the pageSize but not the limit
			return *pageSize
		}
		return min(*pageSize, *limit)
	}
}

func UrlWithoutPaginationInfo(baseUrl string, data url.Values) (string, error) {
	url1, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}
	query := ""
	q := []string{"page", "pageSize"}
	queryMap, _ := url.ParseQuery(url1.RawQuery)
	if queryMap != nil {
		for _, key := range q {
			_, ok := queryMap[key]
			if ok {
				delete(queryMap, key)
			}
		}
		query = queryMap.Encode()
	}
	if data != nil {
		dataCopy := make(url.Values)
		for k, v := range data {
			dataCopy[k] = v
		}
		for _, key := range q {
			_, ok := dataCopy[key]
			if ok {
				delete(dataCopy, key)
			}
		}
		val := dataCopy.Encode()
		if query != "" && val != "" {
			val = "&" + val
		}
		query = query + val
	}

	url1.RawQuery = query
	return url1.String(), nil

}

func GetNext(baseUrl string, response interface{}, getNextPage func(nextPageUri string) (interface{}, error)) (interface{}, error) {
	nextPageUrl, err := getNextPageUrl(baseUrl, response)
	if err != nil {
		return nil, err
	}

	return getNextPage(nextPageUrl)
}

func GetPrevious(baseUrl string, response interface{}, getPreviousPage func(previousPageUri string) (interface{}, error)) (interface{}, error) {
	previousPageUrl, err := getPreviousPageUrl(baseUrl, response)
	if err != nil {
		return nil, err
	}

	return getPreviousPage(previousPageUrl)
}

func toMap(s interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &payload)
	if err != nil {
		return nil, err
	}

	return payload, err
}

func getNextPageUrl(baseUrl string, response interface{}) (string, error) {
	payload, err := toMap(response)
	if err != nil {
		return "", err
	}
	if payload == nil {
		return "", nil
	}

	url1, err1 := url.Parse(baseUrl)
	if err1 != nil {
		return "", err1
	}

	currentPage := 0
	if tmp, ok := payload["page"]; ok {
		currentPage = int(tmp.(float64))
	}
	pageSize := 1
	if tmp, ok := payload["pageSize"]; ok {
		//pageSize = tmp.(int)
		pageSize = int(tmp.(float64))
	}
	totalPages := 1
	if tmp, ok := payload["totalPages"]; ok {
		totalPages = int(tmp.(float64))
	}
	outOfPageRange := true
	tmpOut, okOut := payload["outOfPageRange"]
	if okOut {
		outOfPageRange = tmpOut.(bool)
	}else{
		// because of omitempty
		outOfPageRange = false 

	}

	if !outOfPageRange && (currentPage+1) < totalPages {
		query := fmt.Sprintf("pageSize=%d&page=%d", pageSize, (currentPage + 1))
		if url1.RawQuery != "" {
			query = url1.RawQuery + "&" + query
		}
		url1.RawQuery = query
		return url1.String(), nil
	}

	return "", nil
}

func getPreviousPageUrl(baseUrl string, response interface{}) (string, error) {
	payload, err := toMap(response)
	if err != nil {
		return "", err
	}
	if payload == nil {
		return "", nil
	}

	url1, err1 := url.Parse(baseUrl)
	if err1 != nil {
		return "", err1
	}

	currentPage := 0
	if tmp, ok := payload["page"]; ok {
		currentPage = tmp.(int)
	}
	pageSize := 1
	if tmp, ok := payload["pageSize"]; ok {
		pageSize = tmp.(int)
	}

	if currentPage > 0 {
		query := fmt.Sprintf("pageSize=%d&page=%d", pageSize, (currentPage - 1))
		if url1.RawQuery != "" {
			query = url1.RawQuery + "&" + query
		}
		url1.RawQuery = query
		return url1.String(), nil
	}

	return "", nil
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
