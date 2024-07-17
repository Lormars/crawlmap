package internal

import (
	"net/url"
	"strings"

	"github.com/lormars/crawlmap/common"
)

func ParseURL(rawURL string) (domain, subdomain string, paths []string, queryParams common.QueryParams, err error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return
	}
	hostParts := strings.Split(parsedURL.Hostname(), ".")
	if len(hostParts) > 2 {
		domain = strings.Join(hostParts[len(hostParts)-2:], ".")
		subdomain = parsedURL.Hostname()
	} else {
		domain = parsedURL.Hostname()
		subdomain = domain
	}

	paths = strings.Split(strings.Trim(parsedURL.Path, "/"), "/")

	queryParams = make(common.QueryParams)
	for key, values := range parsedURL.Query() {
		queryParams[key] = values
	}

	return
}
