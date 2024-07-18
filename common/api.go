package common

type QueryParams map[string][]string

type Node struct {
	Name       string
	Children   map[string]*Node
	Params     QueryParams
	StatusCode int
	Origins    []string
}

func NewNode(name string) *Node {
	return &Node{
		Name:       name,
		Children:   make(map[string]*Node),
		Params:     make(QueryParams),
		StatusCode: 0,
		Origins:    []string{},
	}
}

var Nodemap map[string]*Node
