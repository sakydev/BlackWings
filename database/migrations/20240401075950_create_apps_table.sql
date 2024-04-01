-- +goose Up
-- +goose StatementBegin
CREATE TABLE apps (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR,
    provider VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE apps;
-- +goose StatementEnd
