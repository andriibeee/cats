-- +goose Up
CREATE TABLE missions (
    id UUID PRIMARY KEY,
    targets JSONB NOT NULL DEFAULT '[]',
    assignee_id UUID NULL DEFAULT NULL,
    complete BOOLEAN NOT NULL DEFAULT FALSE,
    CONSTRAINT fk_assignee
    FOREIGN KEY(assignee_id)
    REFERENCES cats(id)
);

-- +goose Down
DROP TABLE missions;