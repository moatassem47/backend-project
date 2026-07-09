-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS workout_entries(
    id BIGSERIAL PRIMARY KEY,
    workout_id BIGINT NOT NULL references workouts(id) ON DELETE CASCADE,
    exercise_name VARCHAR(255) NOT NULL,
    sets_num INTEGER,
    reps_num INTEGER,
    duration_seconds INTEGER,
    weight DECIMAL(5,2),
    notes TEXT,
    order_index INTEGER NOT NULL,
    calories_burnt INTEGER NOT NULL,
    created_at timestamp with TIME ZONE default  CURRENT_TIMESTAMP,
    CONSTRAINT valid_workout_entry CHECK(
        (reps_num is not null or duration_seconds is not null )and
        (reps_num is null or duration_seconds is null)
    )
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE workout_entries;
-- +goose StatementEnd
