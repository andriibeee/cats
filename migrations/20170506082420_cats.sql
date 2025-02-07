-- +goose Up
CREATE TABLE cats (
    id          UUID   PRIMARY KEY,
    name        varchar   NOT NULL,
    experience  int NOT NULL DEFAULT 0,
    breed       VARCHAR NOT NULL,
    salary      int NOT NULL DEFAULT 0
);

-- +goose Down
DROP TABLE cats;