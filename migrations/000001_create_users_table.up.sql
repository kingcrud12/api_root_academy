CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    roles TEXT NOT NULL DEFAULT '',
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    instructor_title TEXT NOT NULL DEFAULT '',
    instructor_bio TEXT NOT NULL DEFAULT '',
    avatar TEXT NOT NULL DEFAULT ''
);
