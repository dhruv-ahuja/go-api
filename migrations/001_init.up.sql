CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY,
		isbn INTEGER,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		genre TEXT NOT NULL,
		year INTEGER
		);