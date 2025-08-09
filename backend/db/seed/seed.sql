INSERT INTO books (title, author, published_date)
SELECT 'The Hobbit', 'J.R.R. Tolkien', '1937-09-21'
WHERE NOT EXISTS (SELECT 1 FROM books WHERE title = 'The Hobbit');

INSERT INTO books (title, author, published_date)
SELECT 'The Little Prince', 'Antoine de Saint-Exupéry', '1943-04-06'
WHERE NOT EXISTS (SELECT 1 FROM books WHERE title = 'The Little Prince');

INSERT INTO books (title, author, published_date)
SELECT 'The Lion, the Witch and the Wardrobe', 'C.S. Lewis', '1950-10-16'
WHERE NOT EXISTS (SELECT 1 FROM books WHERE title = 'The Lion, the Witch and the Wardrobe');

INSERT INTO books (title, author, published_date)
SELECT 'The Adventures of Tom Sawyer', 'Mark Twain', '1876-06-01'
WHERE NOT EXISTS (SELECT 1 FROM books WHERE title = 'The Adventures of Tom Sawyer');

INSERT INTO books (title, author, published_date)
SELECT 'The Adventures of Huckleberry Finn', 'Mark Twain', '1884-12-10'
WHERE NOT EXISTS (SELECT 1 FROM books WHERE title = 'The Adventures of Huckleberry Finn');

INSERT INTO books (title, author, published_date)
SELECT 'Pride and Prejudice', 'Jane Austen', '1813-01-28'
WHERE NOT EXISTS (SELECT 1 FROM books WHERE title = 'Pride and Prejudice');

INSERT INTO books (title, author, published_date)
SELECT 'Jane Eyre', 'Charlotte Brontë', '1847-10-16'
WHERE NOT EXISTS (SELECT 1 FROM books WHERE title = 'Jane Eyre');

INSERT INTO books (title, author, published_date)
SELECT 'Heart of Darkness', 'Joseph Conrad', '1902-02-01'
WHERE NOT EXISTS (SELECT 1 FROM books WHERE title = 'Heart of Darkness');

INSERT INTO books (title, author, published_date)
SELECT 'Brave New World', 'Aldous Huxley', '1932-08-01'
WHERE NOT EXISTS (SELECT 1 FROM books WHERE title = 'Brave New World');

INSERT INTO books (title, author, published_date)
SELECT '1984', 'George Orwell', '1949-06-08'
WHERE NOT EXISTS (SELECT 1 FROM books WHERE title = '1984');

INSERT INTO books (title, author, published_date)
SELECT 'The Stranger', 'Albert Camus', '1942-05-01'
WHERE NOT EXISTS (SELECT 1 FROM books WHERE title = 'The Stranger');

INSERT INTO books (title, author, published_date)
SELECT 'The Grapes of Wrath', 'John Steinbeck', '1939-04-14'
WHERE NOT EXISTS (SELECT 1 FROM books WHERE title = 'The Grapes of Wrath');

INSERT INTO books (title, author, published_date)
SELECT 'For Whom the Bell Tolls', 'Ernest Hemingway', '1940-10-21'
WHERE NOT EXISTS (SELECT 1 FROM books WHERE title = 'For Whom the Bell Tolls');

INSERT INTO books (title, author, published_date)
SELECT 'Ulysses', 'James Joyce', '1922-02-02'
WHERE NOT EXISTS (SELECT 1 FROM books WHERE title = 'Ulysses');

INSERT INTO books (title, author, published_date)
SELECT 'The Great Gatsby', 'F. Scott Fitzgerald', '1925-04-10'
WHERE NOT EXISTS (SELECT 1 FROM books WHERE title = 'The Great Gatsby');

INSERT INTO books (title, author, published_date)
SELECT 'Lolita', 'Vladimir Nabokov', '1955-09-15'
WHERE NOT EXISTS (SELECT 1 FROM books WHERE title = 'Lolita');

INSERT INTO books (title, author, published_date)
SELECT 'Catch-22', 'Joseph Heller', '1961-11-10'
WHERE NOT EXISTS (SELECT 1 FROM books WHERE title = 'Catch-22');

