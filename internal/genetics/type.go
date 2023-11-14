package main

import (
	"time"
)

type Waktu map[string]string

func (w Waktu) toTime() map[string]time.Time {
	buffer := map[string]time.Time{}
	format := "15:04 MST"
	for point, waktu := range w {
		res, _ := time.Parse(format, waktu)
		buffer[point] = res
	}
	return buffer
}

func (w Waktu) collision(t Waktu) bool {
	src, dst := w.toTime(), t.toTime()
	var checker bool = true
	a, b := "mulai", "selesai"
	checker = checker && src[a].Before(dst[b])
	checker = checker && src[b].After(dst[a])
	return checker
}

type Jadwal map[string][]Waktu

var Hari = map[int]string{
	0: "Senin",
	1: "Selasa",
	2: "Rabu",
	3: "Kamis",
	4: "Jum'at",
	5: "Sabtu",
}

var Asisten = map[int]string{
	0: "Rizuki",
	1: "Makise",
	2: "Iruma",
}

var MataPraktikum = map[int]string{
	0: "Pemrograman Dasar",
	1: "Pemrograman Mobile",
	2: "Pemrograman Website",
	3: "Pemrograman Desktop",
}

var Kelas = map[int]string{
	0: "IF-A 2021",
	1: "IF-B 2021",
	2: "IF-C 2021",
	3: "IF-A 2023",
	4: "IF-B 2023",
	5: "IF-C 2023",
}

var Ruangan = map[int]string{
	0: "L-KOM I",
	1: "L-KOM II",
	2: "L-KOM III",
}

var Durasi = map[int]Waktu{
	0: {"mulai": "08:00 WIB", "selesai": "09:40 WIB"},
	1: {"mulai": "09:40 WIB", "selesai": "11:20 WIB"},
	2: {"mulai": "11:20 WIB", "selesai": "13:00 WIB"},
	3: {"mulai": "13:00 WIB", "selesai": "14:40 WIB"},
	4: {"mulai": "14:40 WIB", "selesai": "16:20 WIB"},
	5: {"mulai": "16:20 WIB", "selesai": "18:00 WIB"},
}
