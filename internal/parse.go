package internal

import (
	"net/url"
	"strings"

	"github.com/lormars/crawlmap/common"
)

func ParseURL(rawURL string) (domain string, subdomains, paths []string, queryParams common.QueryParams, err error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return
	}
	hostParts := strings.Split(parsedURL.Hostname(), ".")
	if len(hostParts) > 2 {
		domain = strings.Join(hostParts[len(hostParts)-2:], ".")
		for i := len(hostParts) - 2; i >= 0; i-- {
			subdomains = append(subdomains, strings.Join(hostParts[i:], "."))
		}
	} else {
		domain = parsedURL.Hostname()
		subdomains = []string{domain}
	}

	paths = strings.Split(strings.Trim(parsedURL.Path, "/"), "/")

	queryParams = make(common.QueryParams)
	for key, values := range parsedURL.Query() {
		queryParams[key] = values
	}

	return
}
