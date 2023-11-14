package main

var Constraint = map[string]Jadwal{
	"IF-A 2023": {
		"Senin": {
			{"mulai": "10:30 WIB", "selesai": "13:00 WIB"},
			{"mulai": "13:50 WIB", "selesai": "15:30 WIB"},
		},
		"Selasa": {
			{"mulai": "08:00 WIB", "selesai": "10:30 WIB"},
			{"mulai": "10:30 WIB", "selesai": "13:00 WIB"},
		},
		"Rabu": {
			{"mulai": "08:00 WIB", "selesai": "10:30 WIB"},
		},
		"Jum'at": {
			{"mulai": "08:00 WIB", "selesai": "09:40 WIB"},
			{"mulai": "13:30 WIB", "selesai": "16:00 WIB"},
		},
	},
	"IF-B 2023": {
		"Selasa": {
			{"mulai": "08:00 WIB", "selesai": "10:30 WIB"},
			{"mulai": "10:30 WIB", "selesai": "13:00 WIB"},
		},
		"Rabu": {
			{"mulai": "10:30 WIB", "selesai": "13:00 WIB"},
		},
		"Jum'at": {
			{"mulai": "08:00 WIB", "selesai": "10:30 WIB"},
			{"mulai": "13:30 WIB", "selesai": "16:00 WIB"},
		},
		"Sabtu": {
			{"mulai": "09:40 WIB", "selesai": "11:20 WIB"},
			{"mulai": "11:20 WIB", "selesai": "13:00 WIB"},
		},
	},
	"IF-C 2023": {
		"Senin": {
			{"mulai": "12:10 WIB", "selesai": "13:50 WIB"},
		},
		"Selasa": {
			{"mulai": "13:00 WIB", "selesai": "15:30 WIB"},
		},
		"Kamis": {
			{"mulai": "10:30 WIB", "selesai": "13:00 WIB"},
			{"mulai": "13:00 WIB", "selesai": "15:30 WIB"},
		},
		"Jum'at": {
			{"mulai": "09:40 WIB", "selesai": "11:20 WIB"},
		},
		"Sabtu": {
			{"mulai": "08:00 WIB", "selesai": "10:30 WIB"},
			{"mulai": "10:30 WIB", "selesai": "13:00 WIB"},
		},
	},
	"IF-A 2021": {
		"Senin": {
			{"mulai": "08:00 WIB", "selesai": "10:30 WIB"},
		},
		"Selasa": {
			{"mulai": "08:00 WIB", "selesai": "10:30 WIB"},
			{"mulai": "10:30 WIB", "selesai": "13:00 WIB"},
		},
		"Rabu": {
			{"mulai": "08:00 WIB", "selesai": "10:30 WIB"},
			{"mulai": "12:10 WIB", "selesai": "14:40 WIB"},
		},
		"Kamis": {
			{"mulai": "08:00 WIB", "selesai": "10:30 WIB"},
			{"mulai": "10:30 WIB", "selesai": "13:00 WIB"},
		},
	},
	"IF-B 2021": {
		"Senin": {
			{"mulai": "10:30 WIB", "selesai": "13:00 WIB"},
		},
		"Selasa": {
			{"mulai": "08:00 WIB", "selesai": "10:30 WIB"},
			{"mulai": "10:30 WIB", "selesai": "13:00 WIB"},
		},
		"Rabu": {
			{"mulai": "10:30 WIB", "selesai": "13:00 WIB"},
			{"mulai": "14:40 WIB", "selesai": "17:10 WIB"},
		},
		"Kamis": {
			{"mulai": "10:30 WIB", "selesai": "13:00 WIB"},
			{"mulai": "13:00 WIB", "selesai": "15:30 WIB"},
		},
	},
	"IF-C 2021": {
		"Senin": {
			{"mulai": "08:00 WIB", "selesai": "10:30 WIB"},
			{"mulai": "13:00 WIB", "selesai": "15:30 WIB"},
		},
		"Selasa": {
			{"mulai": "13:00 WIB", "selesai": "15:30 WIB"},
		},
		"Rabu": {
			{"mulai": "08:00 WIB", "selesai": "10:30 WIB"},
			{"mulai": "13:00 WIB", "selesai": "15:30 WIB"},
		},
		"Kamis": {
			{"mulai": "10:30 WIB", "selesai": "13:00 WIB"},
			{"mulai": "13:00 WIB", "selesai": "15:30 WIB"},
		},
	},
}
