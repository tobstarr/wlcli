package main

func b2p(in bool) *bool {
	return &in
}

func s2p(in string) *string {
	return &in
}
