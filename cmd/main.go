package main

import (
	"github.com/lormars/crawlmap/pkg"
)

func main() {
	// Define sample inputs
	sampleInputs := []pkg.NodeInput{
		{Url: "https://sub1.example.com/path1?param1=value1", StatusCode: 200},
		{Url: "https://sub1.example.com/path2?param2=value2", StatusCode: 404},
		{Url: "https://sub2.example.com/path1/path2?param3=value3&param4=value4", StatusCode: 500},
		{Url: "https://example.com/path1/path2/path3?param4=value4", StatusCode: 200},
	}

	// Add nodes
	for _, input := range sampleInputs {
		pkg.AddNode(&input)
	}

	// Save the nodemap to a file
	pkg.Save("output.json")

}
