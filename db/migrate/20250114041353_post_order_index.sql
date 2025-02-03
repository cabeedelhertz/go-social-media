-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

-- add index to posts for ordering by created_at desc
CREATE INDEX IF NOT EXISTS idx_posts_latest_created_at ON posts (created_at DESC);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP INDEX IF EXISTS idx_posts_latest_created_at;

-- +goose StatementEnd
