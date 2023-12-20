package main

func newProvider( name string, url string) *provider {
	provider := provider{ name: name, url: url}
	return &provider
}