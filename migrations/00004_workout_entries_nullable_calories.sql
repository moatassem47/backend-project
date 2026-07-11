-- +goose Up
-- +goose StatementBegin
ALTER TABLE workout_entries ALTER COLUMN calories_burnt DROP NOT NULL;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE workout_entries ALTER COLUMN calories_burnt SET NOT NULL;
-- +goose StatementEnd
