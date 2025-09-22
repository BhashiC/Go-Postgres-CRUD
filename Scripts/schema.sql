-- Drop tables if they exist
DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS authors;
DROP TABLE IF EXISTS publishers;

-- Create publishers table
CREATE TABLE publishers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

-- Create authors table
CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    bio TEXT
);

-- Create books table
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    author_id INT REFERENCES authors(id) ON UPDATE CASCADE ON DELETE SET NULL,
    publisher_id INT REFERENCES publishers(id) ON UPDATE CASCADE ON DELETE SET NULL
);
