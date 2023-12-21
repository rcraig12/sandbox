package main

func SetProvider( name string, url string) *provider {
	provider := provider{ name: name, url: url}
	return &provider
}