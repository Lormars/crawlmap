package main

import (
	"github.com/lormars/crawlmap/pkg"
)

func main() {
	// Define sample inputs
	sampleInputs := []pkg.NodeInput{
		// URLs for example.com
		{Url: "https://sub1.example.com/path1?param1=value1", StatusCode: 200, Origin: "script.js"},
		{Url: "https://sub4.sub2.example.com/path2?param2=value2", StatusCode: 404, Origin: "index.html"},
		{Url: "https://sub3.sub2.example.com/path1/path2?param3=value3&param4=value4", StatusCode: 500, Origin: "api.js"},
		{Url: "https://example.com/path1/path2/path3?param4=value4", StatusCode: 200, Origin: "homepage.html"},

		// URLs for abc.com
		{Url: "https://sub1.abc.com/path1?param1=value1", StatusCode: 200, Origin: "main.js"},
		{Url: "https://sub1.abc.com/path2?param2=value2", StatusCode: 302, Origin: "redirect.html"},
		{Url: "https://sub2.abc.com/path1/path2?param3=value3&param4=value4", StatusCode: 404, Origin: "notfound.html"},
		{Url: "https://abc.com/path1/path2/path3?param4=value4", StatusCode: 500, Origin: "server.js"},
		{Url: "https://abc.com/path4?param5=value5", StatusCode: 200, Origin: "test.js"},
		{Url: "https://sub3.abc.com/path5?param6=value6", StatusCode: 301, Origin: "redirect.js"},
	}

	// Add nodes
	for _, input := range sampleInputs {
		pkg.AddNode(&input)
	}

	// Save the nodemap to a file
	pkg.Save("output")

}
