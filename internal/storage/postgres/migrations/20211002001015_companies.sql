-- +goose Up
-- +goose StatementBegin
CREATE TABLE companies
(
    id         BIGSERIAL                NOT NULL PRIMARY KEY,
    alias      VARCHAR                  NOT NULL UNIQUE,
    name       VARCHAR                  NOT NULL,
    created_by BIGINT                   NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    FOREIGN KEY (created_by)
        REFERENCES users (id)
);

CREATE TABLE vacancies
(
    id         BIGSERIAL                NOT NULL PRIMARY KEY,
    company_id BIGINT                   NOT NULL,
    title      VARCHAR                  NOT NULL,
    created_by BIGINT                   NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    FOREIGN KEY (company_id)
        REFERENCES companies (id),
    FOREIGN KEY (created_by)
        REFERENCES users (id)
);

CREATE TABLE reviews
(
    id         BIGSERIAL                NOT NULL PRIMARY KEY,
    parent_id  BIGINT                   NULL,
    company_id BIGINT                   NOT NULL,
    content    TEXT                     NOT NULL,
    created_by BIGINT                   NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    FOREIGN KEY (parent_id)
        REFERENCES reviews (id),
    FOREIGN KEY (company_id)
        REFERENCES companies (id),
    FOREIGN KEY (created_by)
        REFERENCES users (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reviews;
DROP TABLE vacancies;
DROP TABLE companies;
-- +goose StatementEnd
