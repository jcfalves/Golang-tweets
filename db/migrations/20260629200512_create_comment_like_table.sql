-- migrate:up
CREATE TABLE IF NOT EXISTS comment_likes (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    comment_id BIGINT UNSIGNED NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_id_comment_likes FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_post_id_comment_likes FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE
);

-- migrate:down

DROP TABLE IF EXISTS comment_likes;
