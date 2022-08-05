package main

type Filter struct {
	Name    string
	Regex   string
	Options *FilterOpts
}

type FilterOpts struct {
	Replacement      string
	RemoveEmptyLines bool
}
