-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         BIGSERIAL                NOT NULL PRIMARY KEY,
    name       VARCHAR                  NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
