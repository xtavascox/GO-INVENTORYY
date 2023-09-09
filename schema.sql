
CREATE DATABASE inventory;


\c inventory;


CREATE TABLE USERS
(
    id       SERIAL       PRIMARY KEY,
    email    VARCHAR(50)  NOT NULL UNIQUE,
    password VARCHAR(50)  NOT NULL,
    name     VARCHAR(255) NOT NULL
);


CREATE TABLE PRODUCTS
(
    id          SERIAL       PRIMARY KEY,
    name        VARCHAR(50)  NOT NULL,
    description VARCHAR(255) NOT NULL,
    price       FLOAT        NOT NULL,
    created_by  INT          NOT NULL,
    created_at  TIMESTAMP    NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (created_by) REFERENCES USERS (id)
);


CREATE TABLE ROLES
(
    id   SERIAL       PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);


CREATE TABLE USER_ROLES
(
    id      SERIAL PRIMARY KEY,
    user_id INT    NOT NULL,
    role_id INT    NOT NULL,
    FOREIGN KEY (user_id) REFERENCES USERS (id),
    FOREIGN KEY (role_id) REFERENCES ROLES (id)
);
