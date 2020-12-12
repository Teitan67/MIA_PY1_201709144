package main

type mbr struct {
	Tamano  int64
	Fecha   [25]byte
	ID      int64
	Fit     [25]byte
	Memoria int64
	MbrSize int64
}

type particion struct {
	Start     int64
	Size      int64
	Status    bool
	Tipo      [25]byte
	Fit       [25]byte
	Name      [25]byte
	Siguiente int64
}

type disco struct {
	Path          [25]byte
	Mbr           mbr
	ID            [25]byte
	EstadoBorrado bool
	P1            particion
	P2            particion
	P3            particion
	P4            particion
	DiscoMontado  [25]byte
}
