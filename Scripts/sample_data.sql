-- Insert sample publishers
INSERT INTO publishers (name) VALUES
('Rainbow Books'),
('Tech Press'),
('Cloud Publications');

-- Insert sample authors
INSERT INTO authors (name, bio) VALUES
('David Jones', 'An experienced software engineer'),
('John Doe', 'Expert in cloud computing'),
('Alice Smith', 'Author of many programming books');

-- Insert sample books
INSERT INTO books (title, author_id, publisher_id) VALUES
('Golang For Beginners', 1, 1),
('Advanced Golang', 1, 3),
('Cloud Computing Basics', 2, 3),
('Microservices with Go', 1, 3),
('REST APIs in Go', 3, 1),
('Concurrency Patterns', 3, 3);
