-- +migrate Up
CREATE TABLE news_like(
    user_id uuid REFERENCES users(id) ON DELETE CASCADE,
    news_id uuid REFERENCES news(id) ON DELETE CASCADE,
    created_at time DEFAULT now(),

    PRIMARY KEY(user_id, news_id)
);

-- +migrate Down
DROP TABLE news_like;
