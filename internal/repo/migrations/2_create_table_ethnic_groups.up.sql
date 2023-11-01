CREATE TABLE IF NOT EXISTS ethnic_groups (
    country_id INT,
    name TEXT,
    percentage FLOAT,
    FOREIGN KEY (country_id) REFERENCES countries(id)
);