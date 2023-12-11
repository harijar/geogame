CREATE TABLE IF NOT EXISTS guesses
(
    `user_id` Int64,
    `game_id` UUID,
    `country_id` UInt8,
    `text` String,
    `guess_number` UInt8,
    `right` boolean,
    `timestamp` DateTime
)
ENGINE = MergeTree
ORDER BY (user_id, timestamp);