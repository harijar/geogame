CREATE TABLE IF NOT EXISTS languages (
    country_id INT,
    name TEXT,
    FOREIGN KEY (country_id) REFERENCES countries(id)
);