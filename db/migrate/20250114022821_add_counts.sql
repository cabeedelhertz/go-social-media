-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE posts ADD COLUMN like_count int NOT NULL DEFAULT 0;

ALTER TABLE posts ADD COLUMN comment_count int NOT NULL DEFAULT 0;

CREATE UNIQUE INDEX IF NOT EXISTS idx_post_likes_post_id_user_id ON post_likes (post_id, user_id) WHERE deleted_at IS NULL;


-- +goose StatementEnd



-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP INDEX IF EXISTS idx_post_likes_post_id_user_id;

ALTER TABLE posts DROP COLUMN like_count;

ALTER TABLE posts DROP COLUMN comment_count;
-- +goose StatementEnd
