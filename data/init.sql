CREATE DATABASE IF NOT EXISTS robinhood;
USE robinhood;
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO users (username, email, password)
VALUES (
        'sun',
        'sun@example.com',
        '$2a$10$N67KDkgfUsQWk3lZ36G6OeD25buVGGptgFtHtWmvbiYTiIfJayjoK'
    );