package main

type mbr struct {
	tamano     int
	fecha      string
	id         int
	fit        string
	particion1 particion
	particion2 particion
	particion3 particion
	particion4 particion
}

type particion struct {
	status bool
	tipo   string
	fit    string
	start  int
	size   int
	name   string
}
