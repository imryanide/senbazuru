-- init.sql (Database Initialization)
CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL
);

INSERT INTO messages (content) VALUES 
('Hello, world!'),
('Welcome to our API!'),
('This is a sample message.'),
('Message 4 from PostgreSQL!'),
('Message 5 from PostgreSQL!'),
('Message 6 from PostgreSQL!'),
('Message 7 from PostgreSQL!'),
('Message 8 from PostgreSQL!'),
('Message 9 from PostgreSQL!'),
('Message 10 from PostgreSQL!'),
('Message 11 from PostgreSQL!'),
('Message 12 from PostgreSQL!'),
('Message 13 from PostgreSQL!'),
('Message 14 from PostgreSQL!'),
('Message 15 from PostgreSQL!'),
('Message 16 from PostgreSQL!'),
('Message 17 from PostgreSQL!'),
('Message 18 from PostgreSQL!'),
('Message 19 from PostgreSQL!'),
('Message 20 from PostgreSQL!');
