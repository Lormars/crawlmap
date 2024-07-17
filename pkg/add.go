package pkg

import (
	"github.com/lormars/crawlmap/common"
	"github.com/lormars/crawlmap/internal"
)

type NodeInput struct {
	Url        string
	StatusCode int
}

func init() {
	common.Nodemap = make(map[string]*common.Node)
}

func AddNode(input *NodeInput) {
	domain, subdomain, paths, queryParams, err := internal.ParseURL(input.Url)
	if err != nil {
		return
	}

	if _, ok := common.Nodemap[domain]; !ok {
		common.Nodemap[domain] = common.NewNode(domain)
	}

	currentNode := common.Nodemap[domain]

	if _, ok := currentNode.Children[subdomain]; !ok {
		currentNode.Children[subdomain] = common.NewNode(subdomain)
	}

	currentNode = currentNode.Children[subdomain]

	for _, path := range paths {
		if _, ok := currentNode.Children[path]; !ok {
			currentNode.Children[path] = common.NewNode(path)
		}

		currentNode = currentNode.Children[path]
	}

	for key, values := range queryParams {
		currentNode.Params[key] = values
	}

	currentNode.StatusCode = input.StatusCode

}

func Save(fileName string) {
	internal.SaveToFile(fileName)
}
