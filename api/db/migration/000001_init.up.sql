BEGIN;
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
    id   BIGSERIAL PRIMARY KEY,
    username  TEXT NOT NULL UNIQUE,
    firstname TEXT NOT NULL,
    lastname  TEXT NOT NULL,
    email     TEXT NOT NULL UNIQUE,
    password  TEXT NOT NULL
);

CREATE TABLE roles (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE resources (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE actions (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULl UNIQUE
);

CREATE TABLE permissions (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    resource_id BIGINT REFERENCES resources(id) NOT NULL,
    action_id BIGINT REFERENCES actions(id) NOT NULL
);

CREATE TABLE user_roles (
    user_id BIGINT REFERENCES users(id) NOT NULL,
    roles_id BIGINT REFERENCES roles(id) NOT NULL,
    UNIQUE (user_id, roles_id)
);

CREATE TABLE role_permissions (
    role_id BIGINT REFERENCES roles(id) NOT NULL,
    permission_id BIGINT REFERENCES permissions(id) NOT NULL
);

/*
    Seed roles, resources, actions, and permissions
*/
INSERT INTO roles (name) VALUES ('user');
INSERT INTO roles (name) VALUES ('admin');
INSERT INTO resources (name) VALUES ('users');
INSERT INTO resources (name) VALUES ('self');
INSERT INTO actions (name) VALUES ('create');
INSERT INTO actions (name) VALUES ('read');
INSERT INTO permissions (name, resource_id, action_id) VALUES (
   'read_self',
   (SELECT id FROM resources WHERE name = 'self'),
   (SELECT id FROM actions WHERE name = 'read')
);
INSERT INTO permissions (name, resource_id, action_id) VALUES (
    'read_users',
    (SELECT id FROM resources WHERE name = 'users'),
    (SELECT id FROM actions WHERE name = 'read')
);
INSERT INTO role_permissions(role_id, permission_id) VALUES (
     (SELECT id FROM roles WHERE name = 'user'),
     (SELECT id FROM permissions WHERE name = 'read_self')
);
INSERT INTO role_permissions(role_id, permission_id) VALUES (
    (SELECT id FROM roles WHERE name = 'admin'),
    (SELECT id FROM permissions WHERE name = 'read_users')
);

COMMIT;

/* Possibility to have trigger function to handle role updates rather than client
CREATE OR REPLACE FUNCTION assign_user_role() RETURNS TRIGGER AS $table$
BEGIN
    INSERT INTO user_roles(user_id, roles_id) VALUES (new.ID, (SELECT id FROM roles WHERE name = 'user'));
    RETURN NEW;
END;
$table$ LANGUAGE plpgsql;

CREATE TRIGGER assign_user_role_trigger
    AFTER INSERT ON users
    FOR EACH ROW
EXECUTE FUNCTION assign_user_role();
 */
