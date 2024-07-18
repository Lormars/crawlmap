package pkg

import (
	"sync"

	"github.com/lormars/crawlmap/common"
	"github.com/lormars/crawlmap/internal"
)

type NodeInput struct {
	Url        string
	StatusCode int
	Origin     string
}

var mu sync.Mutex
var originMap map[string][]string

func init() {
	common.Nodemap = make(map[string]*common.Node)
	originMap = make(map[string][]string)
}

func AddNode(input *NodeInput) {
	domain, subdomains, paths, queryParams, err := internal.ParseURL(input.Url)
	if err != nil {
		return
	}
	var currentNode *common.Node
	mu.Lock()
	defer mu.Unlock()
	if _, ok := common.Nodemap[domain]; !ok {
		common.Nodemap[domain] = common.NewNode(domain)
		currentNode = common.Nodemap[domain]
	} else {
		currentNode = common.Nodemap[domain]
		// currentNode.Origins = append(currentNode.Origins, input.Origin)
	}

	var sub string
	for _, subdomain := range subdomains {
		if _, ok := currentNode.Children[subdomain]; !ok {
			currentNode.Children[subdomain] = common.NewNode(subdomain)
			currentNode = currentNode.Children[subdomain]
		} else {
			currentNode = currentNode.Children[subdomain]
			// currentNode.Origins = append(currentNode.Origins, input.Origin)
		}
		sub = subdomain
	}

	for _, path := range paths {
		if _, ok := currentNode.Children[path]; !ok {
			currentNode.Children[path] = common.NewNode(path)
			currentNode = currentNode.Children[path]
		} else {
			currentNode = currentNode.Children[path]
			// currentNode.Origins = append(currentNode.Origins, input.Origin)
		}

	}

	for key, values := range queryParams {
		currentNode.Params[key] = values
	}

	currentNode.StatusCode = input.StatusCode
	currentNode.Origins = append(currentNode.Origins, input.Origin)
	originMap[sub] = append(originMap[sub], input.Origin)

}

func Save(dirName string) {
	internal.SaveToFile(dirName)
	mu.Lock()
	defer mu.Unlock()
	common.Nodemap = make(map[string]*common.Node)
	originMap = make(map[string][]string)
}

func ReturnOrigin() map[string][]string {
	mu.Lock()
	defer mu.Unlock()
	return originMap
}
