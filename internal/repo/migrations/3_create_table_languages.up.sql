CREATE TABLE IF NOT EXISTS languages (
    country_id INT,
    name INT,
    FOREIGN KEY (country_id) REFERENCES countries(id)
);