-- Add UUID extension
CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
SET
TIMEZONE="Asia/Jakarta";

-- Create Roles Table
CREATE TABLE roles
(
    id          UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL
);

-- Create Users Table
CREATE TABLE users
(
    id              UUID                     DEFAULT uuid_generate_v4() PRIMARY KEY,
    full_name       VARCHAR(255) NOT NULL,
    email           VARCHAR(255) NOT NULL,
    password        VARCHAR(255) NOT NULL,
    forgot_password VARCHAR(255) NULL,
    role_id         UUID         NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at      TIMESTAMP NULL,
    CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES roles (id)
);

-- Create Status Table
CREATE TABLE status
(
    id   UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Create Todos Table
CREATE TABLE todos
(
    id         UUID                     DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id    UUID         NOT NULL,
    title      VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);

-- Create Todolists Table
CREATE TABLE todolists
(
    id         UUID                     DEFAULT uuid_generate_v4() PRIMARY KEY,
    todo_id    UUID         NOT NULL,
    task       VARCHAR(255) NOT NULL,
    status_id  UUID         NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP NULL,
    CONSTRAINT fk_todo_id FOREIGN KEY (todo_id) REFERENCES todos (id),
    CONSTRAINT fk_status_id FOREIGN KEY (status_id) REFERENCES status (id)
);

-- Create Attachments Table
CREATE TABLE attachments
(
    id          UUID                     DEFAULT uuid_generate_v4() PRIMARY KEY,
    todolist_id UUID         NOT NULL,
    url         VARCHAR(255) NOT NULL,
    caption     VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at  TIMESTAMP NULL,
    CONSTRAINT fk_todolist_id FOREIGN KEY (todolist_id) REFERENCES todolists (id)
);