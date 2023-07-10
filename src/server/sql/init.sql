DROP TABLE IF EXISTS users;

-- Create the 'users' table
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP DEFAULT NOW()
);