-- +migrate Up
CREATE TABLE users(
    id text,
    full_name text,
    email text,
    employee_code int,
    role text,
    salt text,
    password text,
    created_at time DEFAULT now(),
    updated_at time DEFAULT now()
);

-- +migrate Down
DROP TABLE users;