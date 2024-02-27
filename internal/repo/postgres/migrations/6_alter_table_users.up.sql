ALTER TABLE users
    RENAME username TO telegram_username;

ALTER TABLE users
    DROP COLUMN first_name,
    DROP COLUMN last_name,
    ADD COLUMN public bool,
    ADD COLUMN nickname TEXT;