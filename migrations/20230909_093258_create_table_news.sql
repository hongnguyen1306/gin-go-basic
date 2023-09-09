-- +migrate Up
CREATE TABLE news(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    title text,
    content text,
    creator_id uuid REFERENCES users(id) ON DELETE SET NULL,
    created_at time DEFAULT now(),
    updated_at time DEFAULT now(),
    UNIQUE(id)
);

-- +migrate Down
DROP TABLE news;
