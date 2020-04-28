package loader

// BooksLiteralOld is a slice literal of Bookdata struct pointers, containing a subset of the real book data
// Once you have created your CSV loader, you will not need this file.
var BooksLiteralOld = []*BookData{
	{
		BookID:        "1",
		Title:         "Harry Potter and the Half-Blood Prince (Harry Potter  #6)",
		Authors:       "J.K. Rowling-Mary GrandPré",
		AverageRating: 4.56,
		ISBN:          "0439785960",
		ISBN13:        "9780439785969",
		LanguageCode:  "eng",
		NumPages:      652,
		Ratings:       1944099,
		Reviews:       26249,
	},
	{
		BookID:        "2",
		Title:         "Harry Potter and the Order of the Phoenix (Harry Potter  #5)",
		Authors:       "J.K. Rowling-Mary GrandPré",
		AverageRating: 4.49,
		ISBN:          "0439358078",
		ISBN13:        "9780439358071",
		LanguageCode:  "eng",
		NumPages:      870,
		Ratings:       1996446,
		Reviews:       27613,
	},
	{
		BookID:        "3",
		Title:         "Harry Potter and the Sorcerer's Stone (Harry Potter  #1)",
		Authors:       "J.K. Rowling-Mary GrandPré",
		AverageRating: 4.47,
		ISBN:          "0439554934",
		ISBN13:        "9780439554930",
		LanguageCode:  "eng",
		NumPages:      320,
		Ratings:       5629932,
		Reviews:       70390,
	},
	{
		BookID:        "4",
		Title:         "Harry Potter and the Chamber of Secrets (Harry Potter  #2)",
		Authors:       "J.K. Rowling",
		AverageRating: 4.41,
		ISBN:          "0439554896",
		ISBN13:        "9780439554893",
		LanguageCode:  "eng",
		NumPages:      352,
		Ratings:       6267,
		Reviews:       272,
	},
	{
		BookID:        "5",
		Title:         "Harry Potter and the Prisoner of Azkaban (Harry Potter  #3)",
		Authors:       "J.K. Rowling-Mary GrandPré",
		AverageRating: 4.55,
		ISBN:          "043965548X",
		ISBN13:        "9780439655484",
		LanguageCode:  "eng",
		NumPages:      435,
		Ratings:       2149872,
		Reviews:       33964,
	},
	{
		BookID:        "8",
		Title:         "Harry Potter Boxed Set  Books 1-5 (Harry Potter  #1-5)",
		Authors:       "J.K. Rowling-Mary GrandPré",
		AverageRating: 4.78,
		ISBN:          "0439682584",
		ISBN13:        "9780439682589",
		LanguageCode:  "eng",
		NumPages:      2690,
		Ratings:       38872,
		Reviews:       154,
	},
	{
		BookID:        "9",
		Title:         "Unauthorized Harry Potter Book Seven News: Half-Blood Prince Analysis and Speculation",
		Authors:       "W. Frederick Zimmerman",
		AverageRating: 3.69,
		ISBN:          "0976540606",
		ISBN13:        "9780976540601",
		LanguageCode:  "en-US",
		NumPages:      152,
		Ratings:       18,
		Reviews:       1,
	},
	{
		BookID:        "10",
		Title:         "Harry Potter Collection (Harry Potter  #1-6)",
		Authors:       "J.K. Rowling",
		AverageRating: 4.73,
		ISBN:          "0439827604",
		ISBN13:        "9780439827607",
		LanguageCode:  "eng",
		NumPages:      3342,
		Ratings:       27410,
		Reviews:       820,
	},
	{
		BookID:        "12",
		Title:         "The Ultimate Hitchhiker's Guide: Five Complete Novels and One Story (Hitchhiker's Guide to the Galaxy  #1-5)",
		Authors:       "Douglas Adams",
		AverageRating: 4.38,
		ISBN:          "0517226952",
		ISBN13:        "9780517226957",
		LanguageCode:  "eng",
		NumPages:      815,
		Ratings:       3602,
		Reviews:       258,
	},
	{
		BookID:        "13",
		Title:         "The Ultimate Hitchhiker's Guide to the Galaxy",
		Authors:       "Douglas Adams",
		AverageRating: 4.38,
		ISBN:          "0345453743",
		ISBN13:        "9780345453747",
		LanguageCode:  "eng",
		NumPages:      815,
		Ratings:       240189,
		Reviews:       3954,
	},
	{
		BookID:        "14",
		Title:         "The Hitchhiker's Guide to the Galaxy (Hitchhiker's Guide to the Galaxy  #1)",
		Authors:       "Douglas Adams",
		AverageRating: 4.22,
		ISBN:          "1400052920",
		ISBN13:        "9781400052929",
		LanguageCode:  "eng",
		NumPages:      215,
		Ratings:       4416,
		Reviews:       408,
	},
	{
		BookID:        "16",
		Title:         "The Hitchhiker's Guide to the Galaxy (Hitchhiker's Guide to the Galaxy  #1)",
		Authors:       "Douglas Adams-Stephen Fry",
		AverageRating: 4.22,
		ISBN:          "0739322206",
		ISBN13:        "9780739322208",
		LanguageCode:  "eng",
		NumPages:      6,
		Ratings:       1222,
		Reviews:       253,
	},
	{
		BookID:        "18",
		Title:         "The Ultimate Hitchhiker's Guide (Hitchhiker's Guide to the Galaxy #1-5)",
		Authors:       "Douglas Adams",
		AverageRating: 4.38,
		ISBN:          "0517149257",
		ISBN13:        "9780517149256",
		LanguageCode:  "en-US",
		NumPages:      815,
		Ratings:       2801,
		Reviews:       192,
	},
	{
		BookID:        "21",
		Title:         "A Short History of Nearly Everything",
		Authors:       "Bill Bryson-William Roberts",
		AverageRating: 4.2,
		ISBN:          "076790818X",
		ISBN13:        "9780767908184",
		LanguageCode:  "eng",
		NumPages:      544,
		Ratings:       228522,
		Reviews:       8840,
	},
	{
		BookID:        "22",
		Title:         "Bill Bryson's African Diary",
		Authors:       "Bill Bryson",
		AverageRating: 3.43,
		ISBN:          "0767915062",
		ISBN13:        "9780767915069",
		LanguageCode:  "eng",
		NumPages:      55,
		Ratings:       6993,
		Reviews:       470,
	},
	{
		BookID:        "23",
		Title:         "Bryson's Dictionary of Troublesome Words: A Writer's Guide to Getting It Right",
		Authors:       "Bill Bryson",
		AverageRating: 3.88,
		ISBN:          "0767910435",
		ISBN13:        "9780767910439",
		LanguageCode:  "eng",
		NumPages:      256,
		Ratings:       2020,
		Reviews:       124,
	},
	{
		BookID:        "24",
		Title:         "In a Sunburned Country",
		Authors:       "Bill Bryson",
		AverageRating: 4.07,
		ISBN:          "0767903862",
		ISBN13:        "9780767903868",
		LanguageCode:  "eng",
		NumPages:      335,
		Ratings:       68213,
		Reviews:       4077,
	},
	{
		BookID:        "25",
		Title:         "I'm a Stranger Here Myself: Notes on Returning to America After Twenty Years Away",
		Authors:       "Bill Bryson",
		AverageRating: 3.9,
		ISBN:          "076790382X",
		ISBN13:        "9780767903820",
		LanguageCode:  "eng",
		NumPages:      304,
		Ratings:       47490,
		Reviews:       2153,
	},
	{
		BookID:        "26",
		Title:         "The Lost Continent: Travels in Small Town America",
		Authors:       "Bill Bryson",
		AverageRating: 3.83,
		ISBN:          "0060920084",
		ISBN13:        "9780060920081",
		LanguageCode:  "en-US",
		NumPages:      299,
		Ratings:       43779,
		Reviews:       2146,
	},
	{
		BookID:        "27",
		Title:         "Neither Here nor There: Travels in Europe",
		Authors:       "Bill Bryson",
		AverageRating: 3.87,
		ISBN:          "0380713802",
		ISBN13:        "9780380713806",
		LanguageCode:  "eng",
		NumPages:      254,
		Ratings:       46397,
		Reviews:       2127,
	},
	{
		BookID:        "28",
		Title:         "Notes from a Small Island",
		Authors:       "Bill Bryson",
		AverageRating: 3.92,
		ISBN:          "0380727501",
		ISBN13:        "9780380727506",
		LanguageCode:  "eng",
		NumPages:      324,
		Ratings:       76476,
		Reviews:       3159,
	},
	{
		BookID:        "29",
		Title:         "The Mother Tongue: English and How It Got That Way",
		Authors:       "Bill Bryson",
		AverageRating: 3.94,
		ISBN:          "0380715430",
		ISBN13:        "9780380715435",
		LanguageCode:  "eng",
		NumPages:      270,
		Ratings:       26672,
		Reviews:       1986,
	},
	{
		BookID:        "30",
		Title:         "J.R.R. Tolkien 4-Book Boxed Set: The Hobbit and The Lord of the Rings",
		Authors:       "J.R.R. Tolkien",
		AverageRating: 4.59,
		ISBN:          "0345538374",
		ISBN13:        "9780345538376",
		LanguageCode:  "eng",
		NumPages:      1728,
		Ratings:       97731,
		Reviews:       1536,
	},
	{
		BookID:        "31",
		Title:         "The Lord of the Rings (The Lord of the Rings  #1-3)",
		Authors:       "J.R.R. Tolkien",
		AverageRating: 4.49,
		ISBN:          "0618517650",
		ISBN13:        "9780618517657",
		LanguageCode:  "eng",
		NumPages:      1184,
		Ratings:       1670,
		Reviews:       91,
	},
	{
		BookID:        "32",
		Title:         "The Lord of the Rings (The Lord of the Rings  #1-3)",
		Authors:       "J.R.R. Tolkien",
		AverageRating: 4.49,
		ISBN:          "0618346244",
		ISBN13:        "9780618346240",
		LanguageCode:  "eng",
		NumPages:      1137,
		Ratings:       2819,
		Reviews:       139,
	},
	{
		BookID:        "34",
		Title:         "The Fellowship of the Ring (The Lord of the Rings  #1)",
		Authors:       "J.R.R. Tolkien",
		AverageRating: 4.35,
		ISBN:          "0618346252",
		ISBN13:        "9780618346257",
		LanguageCode:  "eng",
		NumPages:      398,
		Ratings:       2009749,
		Reviews:       12784,
	},
	{
		BookID:        "35",
		Title:         "The Lord of the Rings (The Lord of the Rings  #1-3)",
		Authors:       "J.R.R. Tolkien-Alan  Lee",
		AverageRating: 4.49,
		ISBN:          "0618260587",
		ISBN13:        "9780618260584",
		LanguageCode:  "en-US",
		NumPages:      1216,
		Ratings:       1606,
		Reviews:       139,
	},
}
