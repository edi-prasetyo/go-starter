CREATE TABLE IF NOT EXISTS role_permissions (
role_id BIGINT,
permission_id BIGINT,
PRIMARY KEY (role_id, permission_id)
);