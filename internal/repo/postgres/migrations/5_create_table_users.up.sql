CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    username TEXT,
    public BOOL
);