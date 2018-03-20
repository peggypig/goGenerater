package models

import "time"

type users struct {
	username	string
	password	string
	mail	string
	age	int
	birth	time.Time
}
