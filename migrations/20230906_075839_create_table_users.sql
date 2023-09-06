-- +migrate Up
CREATE TABLE users(
    id int,
    full_name text,
    email text,
    employee_code int,
    role text
);

-- +migrate Down
DROP TABLE users;