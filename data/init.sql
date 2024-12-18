CREATE DATABASE IF NOT EXISTS robinhood;
USE robinhood;
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO users (username, email, password)
VALUES (
        'sun',
        'sun@example.com',
        '$2a$10$N67KDkgfUsQWk3lZ36G6OeD25buVGGptgFtHtWmvbiYTiIfJayjoK'
    );
CREATE TABLE threads (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    detail TEXT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'todo',
    createdBy INT NOT NULL,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedBy INT NOT NULL,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE INDEX idx_updated_at ON threads (updatedAt);

CREATE TABLE thread_histories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    thread_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    detail TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    createdBy INT NOT NULL,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedBy INT NOT NULL,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE INDEX idx_thread_id ON thread_histories (thread_id);

CREATE TABLE comments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    thread_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    detail TEXT NOT NULL,
    createdBy INT NOT NULL,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedBy INT NOT NULL,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE INDEX idx_thread_id ON comments (thread_id);