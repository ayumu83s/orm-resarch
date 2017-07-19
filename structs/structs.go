package structs

import "time"

type Actor struct {
	ActorId    int
	FirstName  string
	LastName   string
	LastUpdate time.Time
}

type Category struct {
	CategoryID int
	Name       string
	LastUpdate time.Time
}

type DataSource struct {
	DbName   string
	Host     string
	Port     int
	User     string
	Password string
}