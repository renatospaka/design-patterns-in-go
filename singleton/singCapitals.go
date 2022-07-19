package singleton

import "sync"

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
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
