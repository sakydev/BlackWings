-- +goose Up
-- +goose StatementBegin
INSERT INTO apps (name, provider) VALUES ('Gmail', 'Google'), ('Drive', 'Google');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM apps WHERE name = 'Gmail' OR name = 'Drive';
-- +goose StatementEnd
