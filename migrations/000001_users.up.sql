CREATE TABLE IF NOT EXISTS users (
    full_name varchar(100) NOT NULL,
    email varchar(50) NOT NULL UNIQUE)
