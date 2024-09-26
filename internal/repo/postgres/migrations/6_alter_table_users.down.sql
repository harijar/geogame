ALTER TABLE users
    RENAME telegram_username TO username;

ALTER TABLE users
    ADD COLUMN first_name TEXT,
    ADD COLUMN last_name TEXT,
    DROP COLUMN public;