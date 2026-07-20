-- migrate:up
ALTER TABLE refresh_tokens
ADD  expired_at TIMESTAMP NULL;      

-- migrate:down

ALTER TABLE refresh_tokens
DROP COLUMN expired_at;