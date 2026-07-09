-- +goose Up
-- +goose StatementBegin
ALTER TABLE workout_entries ALTER COLUMN calories_burned DROP NOT NULL;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE workout_entries ALTER COLUMN calories_burned SET NOT NULL;
-- +goose StatementEnd
