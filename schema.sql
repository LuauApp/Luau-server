DROP SCHEMA luau;
CREATE SCHEMA luau DEFAULT CHARACTER SET = utf8mb4;

USE luau;


CREATE TABLE users (
    id bigint unsigned NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    password varchar(100) NOT NULL,

    primary key(id)
);
