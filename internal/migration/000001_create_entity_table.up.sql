CREATE TABLE IF NOT EXISTS entity
(
    id   SERIAL,
    name varchar(255),
    PRIMARY KEY (id)
);

INSERT INTO entity (name) values ('test_name')