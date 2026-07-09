-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS workouts(
    id BIGSERIAL PRIMARY KEY,
    --TODO:user_id
    title VARCHAR(255) NOT NULL,
    description TEXT,
    duration_minutes INTEGER NOT NULL,
    calories_burnt INTEGER NOT NULL,
    created_at timestamp with TIME ZONE default  CURRENT_TIMESTAMP
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE workouts;
-- +goose StatementEnd
