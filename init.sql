CREATE TABLE book (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    title VARCHAR (255) NOT NULL,
    author VARCHAR (255) NOT NULL
);
CREATE TABLE admin (
    id UUID PRIMARY KEY,
    username VARCHAR (255) NOT NULL,
    password VARCHAR (255) NOT NULL
);