-- Create database
CREATE DATABASE user;

-- Switch to the newly created database
USE user;

-- Create assets table
CREATE TABLE assets (
    id VARCHAR(36) PRIMARY KEY,
    parentID VARCHAR(36),
    description VARCHAR(255)
);

-- Insert data into assets table
INSERT INTO assets (id, parentID, description)
VALUES
    ('550e8400-e29b-41d4-a716-446655440000', '650e8400-e29b-41d4-a716-446655440000', 'First Asset'),
    ('660e8400-e29b-41d4-a716-446655440000', '550e8400-e29b-41d4-a716-446655440000', 'Second Asset'),
    ('770e8400-e29b-41d4-a716-446655440000', '550e8400-e29b-41d4-a716-446655440000', 'Third Asset');

-- Create users table with VARCHAR(36) for id and an age column
CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(50),
    lastName VARCHAR(50),
    age INT
);

-- Insert data into users table with UUIDs and age values
INSERT INTO users (id, name, lastName, age)
VALUES
    ('880e8400-e29b-41d4-a716-446655440000', 'John', 'Doe', 30),
    ('990e8400-e29b-41d4-a716-446655440000', 'Jane', 'Smith', 25),
    ('aaa0e840-e29b-41d4-a716-446655440001', 'Alice', 'Johnson', 28);
