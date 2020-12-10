package main

type mbr struct {
	Tamano     int64
	Fecha      [25]byte
	ID         int64
	Fit        [25]byte
	Particion1 particion
	Particion2 particion
	Particion3 particion
	Particion4 particion
}

type particion struct {
	Start  int64
	Size   int64
	Status bool
	Tipo   [25]byte
	Fit    [25]byte
	Name   [25]byte
}
