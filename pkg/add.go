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

func init() {
	common.Nodemap = make(map[string]*common.Node)
}

func AddNode(input *NodeInput) {
	domain, subdomain, paths, queryParams, err := internal.ParseURL(input.Url)
	if err != nil {
		return
	}
	var currentNode *common.Node
	mu.Lock()
	defer mu.Unlock()
	if _, ok := common.Nodemap[domain]; !ok {
		common.Nodemap[domain] = common.NewNode(domain, input.Origin)
		currentNode = common.Nodemap[domain]
	} else {
		currentNode = common.Nodemap[domain]
		currentNode.Origins = append(currentNode.Origins, input.Origin)
	}

	if _, ok := currentNode.Children[subdomain]; !ok {
		currentNode.Children[subdomain] = common.NewNode(subdomain, input.Origin)
		currentNode = currentNode.Children[subdomain]
	} else {
		currentNode = currentNode.Children[subdomain]
		currentNode.Origins = append(currentNode.Origins, input.Origin)
	}

	for _, path := range paths {
		if _, ok := currentNode.Children[path]; !ok {
			currentNode.Children[path] = common.NewNode(path, input.Origin)
			currentNode = currentNode.Children[path]
		} else {
			currentNode = currentNode.Children[path]
			currentNode.Origins = append(currentNode.Origins, input.Origin)
		}

	}

	for key, values := range queryParams {
		currentNode.Params[key] = values
	}

	currentNode.StatusCode = input.StatusCode

}

func Save(dirName string) {
	internal.SaveToFile(dirName)
	mu.Lock()
	defer mu.Unlock()
	common.Nodemap = make(map[string]*common.Node)
}
