CREATE TABLE posts(
    id                TEXT PRIMARY KEY,
    title             TEXT UNIQUE NOT NULL,
    content           TEXT NOT NULL,
    reactions         TEXT,
    created_at        TIMESTAMP NOT NULL,
    updated_at        TIMESTAMP NOT NULL
);