-- +goose Up
-- +goose StatementBegin
CREATE TABLE settings (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name varchar UNIQUE NOT NULL,
  value TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS settings;
-- +goose StatementEnd
