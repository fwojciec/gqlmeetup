CREATE TABLE IF NOT EXISTS agents (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    email text NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS authors (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    agent_id bigint NOT NULL,
    FOREIGN KEY (agent_id) REFERENCES agents (id) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS books (
    id bigserial PRIMARY KEY,
    title text NOT NULL
);

CREATE TABLE IF NOT EXISTS book_authors (
    book_id bigint,
    author_id bigint,
    FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE,
    FOREIGN KEY (author_id) REFERENCES authors (id) ON DELETE CASCADE,
    PRIMARY KEY (book_id, author_id)
);

CREATE TABLE IF NOT EXISTS users (
    email varchar(254) PRIMARY KEY,
    password varchar(60) NOT NULL,
    admin boolean
);

