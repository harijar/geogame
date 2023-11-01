CREATE TABLE IF NOT EXISTS funfacts (
    country_id INT,
    text TEXT,
    FOREIGN KEY (country_id) REFERENCES countries(id)
);