CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY,
    nickname TEXT UNIQUE NOT NULL,
    telegram_username TEXT,
    public BOOL
);