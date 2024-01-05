ALTER TABLE IF EXISTS users
    RENAME telegram_username TO username;

ALTER TABLE IF EXISTS users
    ADD COLUMN first_name TEXT,
    ADD COLUMN last_name TEXT,
    DROP COLUMN public;