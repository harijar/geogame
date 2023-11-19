CREATE TABLE IF NOT EXISTS countries (
    id INT PRIMARY KEY,
    name TEXT,
    aliases TEXT ARRAY,

    un_not_member TEXT,
    unrecognised TEXT,

    capital TEXT,
    religion TEXT,
    religion_perc TEXT,

    population INT,
    area FLOAT,
    gdp INT,
    gdp_per_capita INT,
    hdi FLOAT,
    independent_from TEXT,

    agricultural_sector FLOAT,
    industrial_sector FLOAT,
    service_sector FLOAT,

    northernmost FLOAT,
    southernmost FLOAT,
    easternmost FLOAT,
    westernmost FLOAT,
    hemisphere_lat INT,
    hemisphere_long INT,

    monarchy BOOL,
    landlocked BOOL,
    island BOOL
);