CREATE TABLE IF NOT EXISTS guesses
(
    `user_id` int,
    `game_id` UUID,
    `country_id` int,
    `text` String,
    `guess_number` int,
    `right` boolean,
    `timestamp` DateTime
)
    ENGINE = MergeTree
    ORDER BY (user_id, timestamp);