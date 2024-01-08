ALTER TABLE IF EXISTS users
    RENAME username TO telegram_username;

ALTER TABLE IF EXISTS users
    DROP COLUMN first_name,
    DROP COLUMN last_name,
    ADD COLUMN public bool,
    ADD COLUMN nickname TEXT;