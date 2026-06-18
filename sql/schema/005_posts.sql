-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    title TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    description TEXT,
    published_at timestamp without time zone,
    feed_id UUID NOT NULL,
    CONSTRAINT fk_feed_id
        FOREIGN KEY (feed_id)
        REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;