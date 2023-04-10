package proxy

import "log"

type Driven interface {
	Drive()
}

type Car struct {}

func (c *Car) Drive() {
	log.Println("Car is being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car Car
	driver *Driver
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 18 {
		c.car.Drive()
	} else {
		log.Println("Driver too young")
	}
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{
		Car{},
		driver,
	}
}