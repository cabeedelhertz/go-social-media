-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE users ADD COLUMN bio varchar NOT NULL DEFAULT '';
ALTER TABLE users ADD COLUMN follower_count int NOT NULL DEFAULT 0;
ALTER TABLE users ADD COLUMN following_count int NOT NULL DEFAULT 0;

CREATE TABLE IF NOT EXISTS followers (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    user_id uuid NOT NULL REFERENCES users(id),
    follower_id uuid NOT NULL REFERENCES users(id),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_followers_user_id_follower_id ON followers (user_id, follower_id) WHERE deleted_at IS NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

ALTER TABLE users DROP COLUMN bio;
ALTER TABLE users DROP COLUMN follower_count;
ALTER TABLE users DROP COLUMN following_count;

DROP TABLE IF EXISTS followers;

-- +goose StatementEnd
