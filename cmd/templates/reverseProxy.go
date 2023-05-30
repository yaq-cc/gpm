package main

import (
	"net/http/httputil"
	"net/url"
)

type ReverseProxy struct {
	proxy *httputil.ReverseProxy
}

func New(targetURL string) (*ReverseProxy, error) {
	URL, err := url.Parse(targetURL)
	if err != nil {
		return nil, err
	}
	
	proxy := &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			pr.SetURL(URL)
			pr.SetXForwarded()
			// Add more logic here!
			
		},
	}
	
	reverseProxy := &ReverseProxy{
		proxy: proxy,
	}

	return reverseProxy, nil
}