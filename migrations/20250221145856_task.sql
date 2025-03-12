-- +goose Up
-- +goose StatementBegin
CREATE TYPE task_status AS ENUM ('discovered', 'todo', 'in_progress', 'done', 'cancelled', 'discarded');

CREATE TABLE tasks (
    id UUID PRIMARY KEY NOT NULL,
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    due_date TIMESTAMP,
    status task_status NOT NULL,
    user_id UUID REFERENCES users(id),
    source VARCHAR(100) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
DROP TYPE task_status;
-- +goose StatementEnd
