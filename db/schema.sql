CREATE TABLE IF NOT EXISTS signup (
    id SERIAL PRIMARY KEY,
    user_email VARCHAR(255) NOT NULL,
    user_password VARCHAR(50) NOT NULL
);