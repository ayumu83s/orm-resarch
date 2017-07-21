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

type FilmActor struct {
	ActorId    int
	FilmId     int
	LastUpdate time.Time
}

type FilmText struct {
	FilmId      int
	Title       string
	Description string
}

type City struct {
	CityID     int
	City       string
	CountryID  string
	LastUpdate time.Time
}

// type FilmActorDetail struct {
// 	FilmActor `xorm:"extends"`
// 	//FilmActor
// 	FirstName string
// 	LastName  string
// }
