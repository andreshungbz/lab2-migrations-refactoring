-- Filename: migrations/000003_add_indexes.down.sql
DROP INDEX IF EXISTS idx_users_username;
DROP INDEX IF EXISTS idx_posts_user_id;
