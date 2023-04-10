package singleton

import "sync"

type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int {
			"alpha": 1,
			"beta": 2,
			"gamma": 3,
		}
	}
	return d.dummyData[name]
}

var once sync.Once	// thread safety
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	// garantee there is only run ONCE
	once.Do(func() {
		caps, err := readData("capitals.txt")
		db := singletonDatabase{caps}
		if err == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
	} //DIP
	return result
}

func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	} 
	return result
}
