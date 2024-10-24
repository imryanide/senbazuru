-- init.sql (Database Initialization)
CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL
);

INSERT INTO messages (content) VALUES ('Hello from PostgreSQL!');
