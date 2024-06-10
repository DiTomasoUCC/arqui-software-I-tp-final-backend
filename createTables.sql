CREATE DATABASE uccdemy;

USE uccdemy;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    user_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    user_type BOOLEAN,
    password_hash VARCHAR(255) NOT NULL,
    creation_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create the courses table
CREATE TABLE courses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    instructor_id INT NOT NULL,
    category VARCHAR(255) NOT NULL,
    requirements TEXT NOT NULL,
    length INT NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    creation_time DATETIME NOT NULL,
    last_updated DATETIME NOT NULL,
    FOREIGN KEY (instructor_id) REFERENCES users(id)
);