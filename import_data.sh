#!/bin/bash

login() {
  echo "Attempting to log in..."
  response=$(curl -s -c cookie.txt -X POST http://localhost:8080/login \
    -H "Content-Type: application/json" \
    -d '{"user": "admin", "pass": "securePasswd"}')
  
  if [[ $(echo "$response" | grep -c "Login successful") -gt 0 ]]; then
    echo "Login successful."
  else
    echo "Login failed. Exiting."
    exit 1
  fi
}

add_book() {
  echo "Adding book: $1"
  curl -s -b cookie.txt -X POST http://localhost:8080/books \
    -H "Content-Type: application/json" \
    -d "$1"
}

add_student() {
  echo "Adding student: $1"
  curl -s -b cookie.txt -X POST http://localhost:8080/students \
    -H "Content-Type: application/json" \
    -d "$1"
}


sleep 3
login

add_book '{"title": "To Kill a Mockingbird", "description": "A novel by Harper Lee published in 1960.", "count": 2}'
add_book '{"title": "1984", "description": "A dystopian social science fiction novel by George Orwell.", "count": 5}'
add_book '{"title": "The Great Gatsby", "description": "A novel by F. Scott Fitzgerald.", "count": 3}'
add_book '{"title": "One Hundred Years of Solitude", "description": "A landmark 1967 novel by Colombian author Gabriel García Márquez.", "count": 4}'
add_book '{"title": "Moby Dick", "description": "A novel by Herman Melville.", "count": 2}'
add_book '{"title": "War and Peace", "description": "A novel by the Russian author Leo Tolstoy.", "count": 1}'
add_book '{"title": "The Catcher in the Rye", "description": "A novel by J. D. Salinger.", "count": 6}'
add_book '{"title": "The Hobbit", "description": "A children''s fantasy novel by J. R. R. Tolkien.", "count": 8}'
add_book '{"title": "Ulysses", "description": "A modernist novel by Irish writer James Joyce.", "count": 1}'
add_book '{"title": "The Odyssey", "description": "An ancient Greek epic poem attributed to Homer.", "count": 3}'
add_book '{"title": "The Divine Comedy", "description": "An Italian narrative poem by Dante Alighieri.", "count": 1}'
add_book '{"title": "Brave New World", "description": "A dystopian novel by English author Aldous Huxley.", "count": 5}'
add_book '{"title": "The Sound and the Fury", "description": "A novel by American author William Faulkner.", "count": 2}'
add_book '{"title": "Catch-22", "description": "A satirical novel by American author Joseph Heller.", "count": 4}'
add_book '{"title": "The Grapes of Wrath", "description": "An American realist novel written by John Steinbeck.", "count": 3}'
add_book '{"title": "Don Quixote", "description": "A Spanish epic novel by Miguel de Cervantes.", "count": 2}'
add_book '{"title": "Jane Eyre", "description": "A novel by English writer Charlotte Brontë.", "count": 7}'
add_book '{"title": "Wuthering Heights", "description": "A novel by Emily Brontë.", "count": 4}'
add_book '{"title": "Frankenstein", "description": "A novel written by English author Mary Shelley.", "count": 5}'
add_book '{"title": "The Adventures of Huckleberry Finn", "description": "A novel by Mark Twain.", "count": 6}'
add_book '{"title": "Alice''s Adventures in Wonderland", "description": "An 1865 novel by English author Lewis Carroll.", "count": 9}'
add_book '{"title": "The Picture of Dorian Gray", "description": "A philosophical novel by Oscar Wilde.", "count": 3}'
add_book '{"title": "Dracula", "description": "An 1897 Gothic horror novel by Irish author Bram Stoker.", "count": 4}'
add_book '{"title": "The Scarlet Letter", "description": "A work of historical fiction by American author Nathaniel Hawthorne.", "count": 2}'
add_book '{"title": "A Tale of Two Cities", "description": "An 1859 historical novel by Charles Dickens.", "count": 5}'
add_book '{"title": "The Brothers Karamazov", "description": "The final novel by Russian author Fyodor Dostoevsky.", "count": 1}'
add_book '{"title": "Crime and Punishment", "description": "A novel by the Russian author Fyodor Dostoevsky.", "count": 2}'
add_book '{"title": "The Idiot", "description": "A novel by the Russian author Fyodor Dostoevsky.", "count": 1}'
add_book '{"title": "Demons", "description": "A novel by Fyodor Dostoevsky.", "count": 1}'
add_book '{"title": "The Sun Also Rises", "description": "A 1926 novel by American writer Ernest Hemingway.", "count": 3}'
add_book '{"title": "For Whom the Bell Tolls", "description": "A novel by Ernest Hemingway published in 1940.", "count": 2}'
add_book '{"title": "A Farewell to Arms", "description": "A novel by Ernest Hemingway set during the Italian campaign of World War I.", "count": 3}'
add_book '{"title": "The Old Man and the Sea", "description": "A short novel written by the American author Ernest Hemingway in 1951.", "count": 4}'
add_book '{"title": "Heart of Darkness", "description": "A novella by Polish-English novelist Joseph Conrad.", "count": 2}'
add_book '{"title": "Lord of the Flies", "description": "A 1954 novel by Nobel Prize-winning British author William Golding.", "count": 6}'
add_book '{"title": "Fahrenheit 451", "description": "A 1953 dystopian novel by American writer Ray Bradbury.", "count": 7}'
add_book '{"title": "Slaughterhouse-Five", "description": "A science fiction-infused anti-war novel by Kurt Vonnegut.", "count": 4}'
add_book '{"title": "Cat''s Cradle", "description": "A satirical novel by Kurt Vonnegut.", "count": 3}'
add_book '{"title": "Breakfast of Champions", "description": "A 1973 novel by the American author Kurt Vonnegut.", "count": 2}'
add_book '{"title": "The Bell Jar", "description": "A semi-autobiographical novel by Sylvia Plath.", "count": 3}'
add_book '{"title": "The Stranger", "description": "A 1942 novel by French author Albert Camus.", "count": 2}'
add_book '{"title": "The Plague", "description": "A novel by Albert Camus, published in 1947.", "count": 1}'
add_book '{"title": "The Myth of Sisyphus", "description": "A 1942 philosophical essay by Albert Camus.", "count": 1}'
add_book '{"title": "The Fall", "description": "A philosophical novel by Albert Camus.", "count": 1}'
add_book '{"title": "Nausea", "description": "A 1938 philosophical novel by the existentialist philosopher Jean-Paul Sartre.", "count": 1}'
add_book '{"title": "No Exit", "description": "A 1944 existentialist play by Jean-Paul Sartre.", "count": 1}'
add_book '{"title": "Being and Nothingness", "description": "A 1943 book by the philosopher Jean-Paul Sartre.", "count": 1}'
add_book '{"title": "The Metamorphosis", "description": "A novella written by Franz Kafka which was first published in 1915.", "count": 3}'
add_book '{"title": "The Trial", "description": "A novel written by Franz Kafka between 1914 and 1915.", "count": 2}'
add_book '{"title": "The Castle", "description": "A 1926 novel by Franz Kafka.", "count": 1}'
add_book '{"title": "One Flew Over the Cuckoo''s Nest", "description": "A novel written by Ken Kesey.", "count": 5}'
add_book '{"title": "Sometimes a Great Notion", "description": "A 1964 novel by Ken Kesey.", "count": 2}'
add_book '{"title": "On the Road", "description": "A 1957 novel by American writer Jack Kerouac.", "count": 4}'
add_book '{"title": "The Dharma Bums", "description": "A 1958 novel by Beat Generation author Jack Kerouac.", "count": 2}'
add_book '{"title": "Big Sur", "description": "A 1962 novel by Jack Kerouac.", "count": 1}'
add_book '{"title": "Naked Lunch", "description": "A 1959 novel by American writer William S. Burroughs.", "count": 1}'
add_book '{"title": "Junkie", "description": "A 1953 novel by William S. Burroughs.", "count": 1}'
add_book '{"title": "Queer", "description": "A short novel by William S. Burroughs.", "count": 1}'
add_book '{"title": "Howl", "description": "A poem written by Allen Ginsberg in 1955.", "count": 1}'
add_book '{"title": "A Clockwork Orange", "description": "A dystopian satirical black comedy novel by English writer Anthony Burgess, published in 1962.", "count": 4}'
add_book '{"title": "The Man in the High Castle", "description": "An alternate history novel by American writer Philip K. Dick.", "count": 3}'
add_book '{"title": "Do Androids Dream of Electric Sheep?", "description": "A science fiction novel by American writer Philip K. Dick.", "count": 4}'
add_book '{"title": "A Scanner Darkly", "description": "A 1977 science fiction novel by American writer Philip K. Dick.", "count": 2}'
add_book '{"title": "VALIS", "description": "A 1981 science fiction novel by Philip K. Dick.", "count": 1}'
add_book '{"title": "Neuromancer", "description": "A 1984 science fiction novel by American-Canadian writer William Gibson.", "count": 3}'
add_book '{"title": "Count Zero", "description": "A science fiction novel by William Gibson.", "count": 2}'
add_book '{"title": "Mona Lisa Overdrive", "description": "A science fiction novel by William Gibson.", "count": 1}'
add_book '{"title": "Snow Crash", "description": "A science fiction novel by the American writer Neal Stephenson, published in 1992.", "count": 4}'
add_book '{"title": "Cryptonomicon", "description": "A 1999 novel by the American author Neal Stephenson.", "count": 2}'
add_book '{"title": "Anathem", "description": "A speculative fiction novel by Neal Stephenson.", "count": 1}'
add_book '{"title": "The Diamond Age", "description": "A science fiction novel by Neal Stephenson.", "count": 2}'
add_book '{"title": "Dune", "description": "A 1965 science-fiction novel by American author Frank Herbert.", "count": 7}'
add_book '{"title": "Dune Messiah", "description": "A science fiction novel by Frank Herbert.", "count": 4}'
add_book '{"title": "Children of Dune", "description": "A 1976 science fiction novel by Frank Herbert.", "count": 3}'
add_book '{"title": "God Emperor of Dune", "description": "A science fiction novel by Frank Herbert.", "count": 2}'
add_book '{"title": "Heretics of Dune", "description": "A 1984 science fiction novel by Frank Herbert.", "count": 1}'
add_book '{"title": "Chapterhouse: Dune", "description": "A 1985 science fiction novel by Frank Herbert.", "count": 1}'
add_book '{"title": "Foundation", "description": "A science fiction novel by American writer Isaac Asimov.", "count": 6}'
add_book '{"title": "Foundation and Empire", "description": "A science fiction novel by Isaac Asimov.", "count": 4}'
add_book '{"title": "Second Foundation", "description": "A science fiction novel by Isaac Asimov.", "count": 3}'
add_book '{"title": "I, Robot", "description": "A collection of nine science fiction short stories by Isaac Asimov.", "count": 5}'
add_book '{"title": "The Caves of Steel", "description": "A science fiction novel by Isaac Asimov.", "count": 3}'
add_book '{"title": "The Naked Sun", "description": "A science fiction novel by Isaac Asimov.", "count": 2}'
add_book '{"title": "The Robots of Dawn", "description": "A science fiction novel by Isaac Asimov.", "count": 1}'
add_book '{"title": "Robots and Empire", "description": "A science fiction novel by Isaac Asimov.", "count": 1}'
add_book '{"title": "The Hitchhiker''s Guide to the Galaxy", "description": "A comedy science fiction series created by Douglas Adams.", "count": 5}'
add_book '{"title": "The Restaurant at the End of the Universe", "description": "The second book in the Hitchhiker''s Guide to the Galaxy series.", "count": 4}'
add_book '{"title": "Life, the Universe and Everything", "description": "The third book in the Hitchhiker''s Guide to the Galaxy series.", "count": 3}'
add_book '{"title": "So Long, and Thanks for All the Fish", "description": "The fourth book in the Hitchhiker''s Guide to the Galaxy series.", "count": 2}'
add_book '{"title": "Mostly Harmless", "description": "The fifth book in the Hitchhiker''s Guide to the Galaxy series.", "count": 1}'
add_book '{"title": "The Colour of Magic", "description": "A 1983 fantasy novel by Terry Pratchett.", "count": 4}'
add_book '{"title": "The Light Fantastic", "description": "A 1986 fantasy novel by Terry Pratchett.", "count": 3}'
add_book '{"title": "Equal Rites", "description": "A 1987 fantasy novel by Terry Pratchett.", "count": 2}'
add_book '{"title": "Mort", "description": "A 1987 fantasy novel by Terry Pratchett.", "count": 3}'

# Add some students
add_student '{"card_id": "HVB001", "first_name": "John", "last_name": "Doe", "major": "Computer Science", "phone": "123-456-7890"}'
add_student '{"card_id": "HVB002", "first_name": "Jane", "last_name": "Smith", "major": "Physics", "phone": "098-765-4321"}'
add_student '{"card_id": "HVB003", "first_name": "Peter", "last_name": "Jones", "major": "Mathematics", "phone": "111-222-3333"}'
add_student '{"card_id": "HVB004", "first_name": "Mary", "last_name": "Johnson", "major": "Chemistry", "phone": "222-333-4444"}'
add_student '{"card_id": "HVB005", "first_name": "David", "last_name": "Williams", "major": "Biology", "phone": "333-444-5555"}'
add_student '{"card_id": "HVB006", "first_name": "Susan", "last_name": "Brown", "major": "English", "phone": "444-555-6666"}'
add_student '{"card_id": "HVB007", "first_name": "Michael", "last_name": "Davis", "major": "History", "phone": "555-666-7777"}'
add_student '{"card_id": "HVB008", "first_name": "Karen", "last_name": "Miller", "major": "Art", "phone": "666-777-8888"}'
add_student '{"card_id": "HVB009", "first_name": "William", "last_name": "Wilson", "major": "Music", "phone": "777-888-9999"}'
add_student '{"card_id": "HVB010", "first_name": "Linda", "last_name": "Moore", "major": "Geography", "phone": "888-999-0000"}'
add_student '{"card_id": "HVB011", "first_name": "James", "last_name": "Taylor", "major": "Philosophy", "phone": "999-000-1111"}'
add_student '{"card_id": "HVB012", "first_name": "Patricia", "last_name": "Anderson", "major": "Sociology", "phone": "000-111-2222"}'
add_student '{"card_id": "HVB013", "first_name": "Robert", "last_name": "Thomas", "major": "Anthropology", "phone": "111-222-3333"}'
add_student '{"card_id": "HVB014", "first_name": "Jennifer", "last_name": "Jackson", "major": "Psychology", "phone": "222-333-4444"}'
add_student '{"card_id": "HVB015", "first_name": "Charles", "last_name": "White", "major": "Economics", "phone": "333-444-5555"}'
add_student '{"card_id": "HVB016", "first_name": "Barbara", "last_name": "Harris", "major": "Political Science", "phone": "444-555-6666"}'
add_student '{"card_id": "HVB017", "first_name": "Richard", "last_name": "Martin", "major": "Linguistics", "phone": "555-666-7777"}'
add_student '{"card_id": "HVB018", "first_name": "Elizabeth", "last_name": "Thompson", "major": "Education", "phone": "666-777-8888"}'
add_student '{"card_id": "HVB019", "first_name": "Joseph", "last_name": "Garcia", "major": "Engineering", "phone": "777-888-9999"}'
add_student '{"card_id": "HVB020", "first_name": "Jessica", "last_name": "Martinez", "major": "Medicine", "phone": "888-999-0000"}'

rm cookie.txt

echo "Data import complete."