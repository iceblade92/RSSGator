-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    user_id UUID NOT NULL,
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;