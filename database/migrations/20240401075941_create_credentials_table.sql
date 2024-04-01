-- +goose Up
-- +goose StatementBegin
CREATE TABLE credentials (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    client_id TEXT,
    client_secret TEXT,
    raw TEXT,
    app_id INTEGER,
    FOREIGN KEY (app_id) REFERENCES apps(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE credentials;
-- +goose StatementEnd
