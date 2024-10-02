package data

var Authors = []Author{
	{
		Name:    "Андрій",
		Surname: "Андрощук",
	},
	{
		Name:    "Богдан",
		Surname: "Боровик",
	},
	{
		Name:    "Вікторія",
		Surname: "Василенко",
	},
}
var Publishers = []Publisher{
	{
		Name: "Анархія-друк",
	},
	{
		Name: "Барка і штиль",
	},
	{
		Name: "Вокабюляри нової доби",
	},
}
var Books = []Book{
	{
		Title:        "Асканія-Нова. Історія заповідника",
		Author_ID:    1,
		Price:        380.00,
		Publisher_ID: 1,
		Published:    2024,
		ISBN:         "1111111111111",
	},
	{
		Title:        "Брати і кузени",
		Author_ID:    2,
		Price:        300.00,
		Publisher_ID: 2,
		Published:    2024,
		ISBN:         "1111111111112",
	},
	{
		Title:        "Віра і майна",
		Author_ID:    3,
		Price:        350.00,
		Publisher_ID: 3,
		Published:    2024,
		ISBN:         "1111111111113",
	},
}
