package config

import (
	"context"
	"math/rand"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"

	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/models"
	"BRSBackend/pkg/services"
)

func SeedData(bookService services.BookService, studentService services.StudentService, rentService services.RentService) {
	seedBooks(bookService)
	seedStudents(studentService)
	seedRents(rentService, bookService, studentService)
}

func seedRents(rentService services.RentService, bookService services.BookService, studentService services.StudentService) {
	books, err := bookService.GetAllBooks(context.Background(), dto.PaginationParams{Limit: 100, Offset: 0})
	if err != nil {
		log.Errorf("Failed to get books for seeding rents: %v", err)
		return
	}

	students, err := studentService.GetAllStudents(context.Background(), dto.PaginationParams{Limit: 100, Offset: 0})
	if err != nil {
		log.Errorf("Failed to get students for seeding rents: %v", err)
		return
	}

	if len(books.Results) == 0 || len(students.Results) == 0 {
		log.Info("No books or students to seed rents with.")
		return
	}

	for i := 0; i < 35; i++ {
		student := students.Results[rand.Intn(len(students.Results))]
		book := books.Results[rand.Intn(len(books.Results))]

		rentRequest := dto.CreateRentRequest{
			StudentID: student.Id,
			BookIDs:   []uuid.UUID{book.Id},
		}

		_, err := rentService.CreateRentTransaction(context.Background(), rentRequest)
		if err != nil {
			log.Errorf("Failed to create rent transaction: %v", err)
		}
	}
}

func seedBooks(bookService services.BookService) {
	books := []models.Book{
		{Title: "To Kill a Mockingbird", Description: "A novel by Harper Lee published in 1960.", Count: 2},
		{Title: "1984", Description: "A dystopian social science fiction novel by George Orwell.", Count: 5},
		{Title: "The Great Gatsby", Description: "A novel by F. Scott Fitzgerald.", Count: 3},
		{Title: "One Hundred Years of Solitude", Description: "A landmark 1967 novel by Colombian author Gabriel García Márquez.", Count: 4},
		{Title: "Moby Dick", Description: "A novel by Herman Melville.", Count: 2},
		{Title: "War and Peace", Description: "A novel by the Russian author Leo Tolstoy.", Count: 1},
		{Title: "The Catcher in the Rye", Description: "A novel by J. D. Salinger.", Count: 6},
		{Title: "The Hobbit", Description: "A children's fantasy novel by J. R. R. Tolkien.", Count: 8},
		{Title: "Ulysses", Description: "A modernist novel by Irish writer James Joyce.", Count: 1},
		{Title: "The Odyssey", Description: "An ancient Greek epic poem attributed to Homer.", Count: 3},
		{Title: "The Divine Comedy", Description: "An Italian narrative poem by Dante Alighieri.", Count: 1},
		{Title: "Brave New World", Description: "A dystopian novel by English author Aldous Huxley.", Count: 5},
		{Title: "The Sound and the Fury", Description: "A novel by American author William Faulkner.", Count: 2},
		{Title: "Catch-22", Description: "A satirical novel by American author Joseph Heller.", Count: 4},
		{Title: "The Grapes of Wrath", Description: "An American realist novel written by John Steinbeck.", Count: 3},
		{Title: "Don Quixote", Description: "A Spanish epic novel by Miguel de Cervantes.", Count: 2},
		{Title: "Jane Eyre", Description: "A novel by English writer Charlotte Brontë.", Count: 7},
		{Title: "Wuthering Heights", Description: "A novel by Emily Brontë.", Count: 4},
		{Title: "Frankenstein", Description: "A novel written by English author Mary Shelley.", Count: 5},
		{Title: "The Adventures of Huckleberry Finn", Description: "A novel by Mark Twain.", Count: 6},
		{Title: "Alice's Adventures in Wonderland", Description: "An 1865 novel by English author Lewis Carroll.", Count: 9},
		{Title: "The Picture of Dorian Gray", Description: "A philosophical novel by Oscar Wilde.", Count: 3},
		{Title: "Dracula", Description: "An 1897 Gothic horror novel by Irish author Bram Stoker.", Count: 4},
		{Title: "The Scarlet Letter", Description: "A work of historical fiction by American author Nathaniel Hawthorne.", Count: 2},
		{Title: "A Tale of Two Cities", Description: "An 1859 historical novel by Charles Dickens.", Count: 5},
		{Title: "The Brothers Karamazov", Description: "The final novel by Russian author Fyodor Dostoevsky.", Count: 1},
		{Title: "Crime and Punishment", Description: "A novel by the Russian author Fyodor Dostoevsky.", Count: 2},
		{Title: "The Idiot", Description: "A novel by the Russian author Fyodor Dostoevsky.", Count: 1},
		{Title: "Demons", Description: "A novel by Fyodor Dostoevsky.", Count: 1},
		{Title: "The Sun Also Rises", Description: "A 1926 novel by American writer Ernest Hemingway.", Count: 3},
		{Title: "For Whom the Bell Tolls", Description: "A novel by Ernest Hemingway published in 1940.", Count: 2},
		{Title: "A Farewell to Arms", Description: "A novel by Ernest Hemingway set during the Italian campaign of World War I.", Count: 3},
		{Title: "The Old Man and the Sea", Description: "A short novel written by the American author Ernest Hemingway in 1951.", Count: 4},
		{Title: "Heart of Darkness", Description: "A novella by Polish-English novelist Joseph Conrad.", Count: 2},
		{Title: "Lord of the Flies", Description: "A 1954 novel by Nobel Prize-winning British author William Golding.", Count: 6},
		{Title: "Fahrenheit 451", Description: "A 1953 dystopian novel by American writer Ray Bradbury.", Count: 7},
		{Title: "Slaughterhouse-Five", Description: "A science fiction-infused anti-war novel by Kurt Vonnegut.", Count: 4},
		{Title: "Cat's Cradle", Description: "A satirical novel by Kurt Vonnegut.", Count: 3},
		{Title: "Breakfast of Champions", Description: "A 1973 novel by the American author Kurt Vonnegut.", Count: 2},
		{Title: "The Bell Jar", Description: "A semi-autobiographical novel by Sylvia Plath.", Count: 3},
		{Title: "The Stranger", Description: "A 1942 novel by French author Albert Camus.", Count: 2},
		{Title: "The Plague", Description: "A novel by Albert Camus, published in 1947.", Count: 1},
		{Title: "The Myth of Sisyphus", Description: "A 1942 philosophical essay by Albert Camus.", Count: 1},
		{Title: "The Fall", Description: "A philosophical novel by Albert Camus.", Count: 1},
		{Title: "Nausea", Description: "A 1938 philosophical novel by the existentialist philosopher Jean-Paul Sartre.", Count: 1},
		{Title: "No Exit", Description: "A 1944 existentialist play by Jean-Paul Sartre.", Count: 1},
		{Title: "Being and Nothingness", Description: "A 1943 book by the philosopher Jean-Paul Sartre.", Count: 1},
		{Title: "The Metamorphosis", Description: "A novella written by Franz Kafka which was first published in 1915.", Count: 3},
		{Title: "The Trial", Description: "A novel written by Franz Kafka between 1914 and 1915.", Count: 2},
		{Title: "The Castle", Description: "A 1926 novel by Franz Kafka.", Count: 1},
		{Title: "One Flew Over the Cuckoo's Nest", Description: "A novel written by Ken Kesey.", Count: 5},
		{Title: "Sometimes a Great Notion", Description: "A 1964 novel by Ken Kesey.", Count: 2},
		{Title: "On the Road", Description: "A 1957 novel by American writer Jack Kerouac.", Count: 4},
		{Title: "The Dharma Bums", Description: "A 1958 novel by Beat Generation author Jack Kerouac.", Count: 2},
		{Title: "Big Sur", Description: "A 1962 novel by Jack Kerouac.", Count: 1},
		{Title: "Naked Lunch", Description: "A 1959 novel by American writer William S. Burroughs.", Count: 1},
		{Title: "Junkie", Description: "A 1953 novel by William S. Burroughs.", Count: 1},
		{Title: "Queer", Description: "A short novel by William S. Burroughs.", Count: 1},
		{Title: "Howl", Description: "A poem written by Allen Ginsberg in 1955.", Count: 1},
		{Title: "A Clockwork Orange", Description: "A dystopian satirical black comedy novel by English writer Anthony Burgess, published in 1962.", Count: 4},
		{Title: "The Man in the High Castle", Description: "An alternate history novel by American writer Philip K. Dick.", Count: 3},
		{Title: "Do Androids Dream of Electric Sheep?", Description: "A science fiction novel by American writer Philip K. Dick.", Count: 4},
		{Title: "A Scanner Darkly", Description: "A 1977 science fiction novel by American writer Philip K. Dick.", Count: 2},
		{Title: "VALIS", Description: "A 1981 science fiction novel by Philip K. Dick.", Count: 1},
		{Title: "Neuromancer", Description: "A 1984 science fiction novel by American-Canadian writer William Gibson.", Count: 3},
		{Title: "Count Zero", Description: "A science fiction novel by William Gibson.", Count: 2},
		{Title: "Mona Lisa Overdrive", Description: "A science fiction novel by William Gibson.", Count: 1},
		{Title: "Snow Crash", Description: "A science fiction novel by the American writer Neal Stephenson, published in 1992.", Count: 4},
		{Title: "Cryptonomicon", Description: "A 1999 novel by the American author Neal Stephenson.", Count: 2},
		{Title: "Anathem", Description: "A speculative fiction novel by Neal Stephenson.", Count: 1},
		{Title: "The Diamond Age", Description: "A science fiction novel by Neal Stephenson.", Count: 2},
		{Title: "Dune", Description: "A 1965 science-fiction novel by American author Frank Herbert.", Count: 7},
		{Title: "Dune Messiah", Description: "A science fiction novel by Frank Herbert.", Count: 4},
		{Title: "Children of Dune", Description: "A 1976 science fiction novel by Frank Herbert.", Count: 3},
		{Title: "God Emperor of Dune", Description: "A science fiction novel by Frank Herbert.", Count: 2},
		{Title: "Heretics of Dune", Description: "A 1984 science fiction novel by Frank Herbert.", Count: 1},
		{Title: "Chapterhouse: Dune", Description: "A 1985 science fiction novel by Frank Herbert.", Count: 1},
		{Title: "Foundation", Description: "A science fiction novel by American writer Isaac Asimov.", Count: 6},
		{Title: "Foundation and Empire", Description: "A science fiction novel by Isaac Asimov.", Count: 4},
		{Title: "Second Foundation", Description: "A science fiction novel by Isaac Asimov.", Count: 3},
		{Title: "I, Robot", Description: "A collection of nine science fiction short stories by Isaac Asimov.", Count: 5},
		{Title: "The Caves of Steel", Description: "A science fiction novel by Isaac Asimov.", Count: 3},
		{Title: "The Naked Sun", Description: "A science fiction novel by Isaac Asimov.", Count: 2},
		{Title: "The Robots of Dawn", Description: "A science fiction novel by Isaac Asimov.", Count: 1},
		{Title: "Robots and Empire", Description: "A science fiction novel by Isaac Asimov.", Count: 1},
		{Title: "The Hitchhiker's Guide to the Galaxy", Description: "A comedy science fiction series created by Douglas Adams.", Count: 5},
		{Title: "The Restaurant at the End of the Universe", Description: "The second book in the Hitchhiker's Guide to the Galaxy series.", Count: 4},
		{Title: "Life, the Universe and Everything", Description: "The third book in the Hitchhiker's Guide to the Galaxy series.", Count: 3},
		{Title: "So Long, and Thanks for All the Fish", Description: "The fourth book in the Hitchhiker's Guide to the Galaxy series.", Count: 2},
		{Title: "Mostly Harmless", Description: "The fifth book in the Hitchhiker's Guide to the Galaxy series.", Count: 1},
		{Title: "The Colour of Magic", Description: "A 1983 fantasy novel by Terry Pratchett.", Count: 4},
		{Title: "The Light Fantastic", Description: "A 1986 fantasy novel by Terry Pratchett.", Count: 3},
		{Title: "Equal Rites", Description: "A 1987 fantasy novel by Terry Pratchett.", Count: 2},
		{Title: "Mort", Description: "A 1987 fantasy novel by Terry Pratchett.", Count: 3},
	}

	for _, book := range books {
		if err := bookService.CreateBook(context.Background(), &book); err != nil {
			log.Errorf("Failed to create book: %v", err)
		}
	}
}

func seedStudents(studentService services.StudentService) {
	students := []models.Student{
		{FirstName: "John", LastName: "Doe", CardId: "HVB001", Major: "Computer Science", Phone: "123-456-7890"},
		{FirstName: "Jane", LastName: "Smith", CardId: "HVB002", Major: "Physics", Phone: "098-765-4321"},
		{FirstName: "Peter", LastName: "Jones", CardId: "HVB003", Major: "Mathematics", Phone: "111-222-3333"},
		{FirstName: "Mary", LastName: "Johnson", CardId: "HVB004", Major: "Chemistry", Phone: "222-333-4444"},
		{FirstName: "David", LastName: "Williams", CardId: "HVB005", Major: "Biology", Phone: "333-444-5555"},
		{FirstName: "Susan", LastName: "Brown", CardId: "HVB006", Major: "English", Phone: "444-555-6666"},
		{FirstName: "Michael", LastName: "Davis", CardId: "HVB007", Major: "History", Phone: "555-666-7777"},
		{FirstName: "Karen", LastName: "Miller", CardId: "HVB008", Major: "Art", Phone: "666-777-8888"},
		{FirstName: "William", LastName: "Wilson", CardId: "HVB009", Major: "Music", Phone: "777-888-9999"},
		{FirstName: "Linda", LastName: "Moore", CardId: "HVB010", Major: "Geography", Phone: "888-999-0000"},
		{FirstName: "James", LastName: "Taylor", CardId: "HVB011", Major: "Philosophy", Phone: "999-000-1111"},
		{FirstName: "Patricia", LastName: "Anderson", CardId: "HVB012", Major: "Sociology", Phone: "000-111-2222"},
		{FirstName: "Robert", LastName: "Thomas", CardId: "HVB013", Major: "Anthropology", Phone: "111-222-3333"},
		{FirstName: "Jennifer", LastName: "Jackson", CardId: "HVB014", Major: "Psychology", Phone: "222-333-4444"},
		{FirstName: "Charles", LastName: "White", CardId: "HVB015", Major: "Economics", Phone: "333-444-5555"},
		{FirstName: "Barbara", LastName: "Harris", CardId: "HVB016", Major: "Political Science", Phone: "444-555-6666"},
		{FirstName: "Richard", LastName: "Martin", CardId: "HVB017", Major: "Linguistics", Phone: "555-666-7777"},
		{FirstName: "Elizabeth", LastName: "Thompson", CardId: "HVB018", Major: "Education", Phone: "666-777-8888"},
		{FirstName: "Joseph", LastName: "Garcia", CardId: "HVB019", Major: "Engineering", Phone: "777-888-9999"},
		{FirstName: "Jessica", LastName: "Martinez", CardId: "HVB020", Major: "Medicine", Phone: "888-999-0000"},
	}

	for _, student := range students {
		if err := studentService.CreateStudent(context.Background(), &student); err != nil {
			log.Errorf("Failed to create student: %v", err)
		}
	}
}
