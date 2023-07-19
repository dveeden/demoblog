package main

import "time"

type Comment struct {
	Id      uint64
	Comment string
	Author  string
	Created time.Time
}
