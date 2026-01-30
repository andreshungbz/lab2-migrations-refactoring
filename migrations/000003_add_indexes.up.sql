-- Filename: 000003_add_indexes.up.sql
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_posts_user_id ON posts(user_id);
