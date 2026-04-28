CREATE TABLE IF NOT EXISTS user_roles (
user_id BIGINT,
role_id BIGINT,
PRIMARY KEY (user_id, role_id)
);