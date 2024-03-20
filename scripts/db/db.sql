-- DROP DATABASE IF EXISTS api;
-- CREATE DATABASE api;
-- USE api;
-- DROP DATABASE IF EXISTS api;
-- CREATE DATABASE api;
-- USE api;
CREATE TABLE customer (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE
);

INSERT INTO customer (id, name, active)
VALUES 
    (1, 'Big News Media Corp', TRUE),
    (2, 'Online Mega Store', TRUE),
    (3, 'Nachoroo Delivery', FALSE),
    (4, 'Euro Telecom Group', TRUE);

CREATE TABLE hourly_stats (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    time TIMESTAMP NOT NULL,
    request_count BIGINT NOT NULL DEFAULT 0,
    invalid_count BIGINT NOT NULL DEFAULT 0,
    UNIQUE (customer_id, time),
    CONSTRAINT fk_customer
        FOREIGN KEY (customer_id) 
        REFERENCES customer (id) 
        ON DELETE CASCADE
);

CREATE TABLE ip_blacklist (
    ip BIGINT PRIMARY KEY
);

INSERT INTO ip_blacklist (ip)
VALUES 
    (0),
    (2130706433),
    (4294967295);

CREATE TABLE ua_blacklist (
    ua VARCHAR(255) PRIMARY KEY
);

INSERT INTO ua_blacklist (ua)
VALUES 
    ('A6-Indexer'),
    ('Googlebot-News'),
    ('Googlebot');
CREATE TABLE customer (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE
);

INSERT INTO customer (id, name, active)
VALUES 
    (1, 'Big News Media Corp', TRUE),
    (2, 'Online Mega Store', TRUE),
    (3, 'Nachoroo Delivery', FALSE),
    (4, 'Euro Telecom Group', TRUE);

CREATE TABLE hourly_stats (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    time TIMESTAMP NOT NULL,
    request_count BIGINT NOT NULL DEFAULT 0,
    invalid_count BIGINT NOT NULL DEFAULT 0,
    UNIQUE (customer_id, time),
    CONSTRAINT fk_customer
        FOREIGN KEY (customer_id) 
        REFERENCES customer (id) 
        ON DELETE CASCADE
);

CREATE TABLE ip_blacklist (
    ip BIGINT PRIMARY KEY
);

INSERT INTO ip_blacklist (ip)
VALUES 
    (0),
    (2130706433),
    (4294967295);

CREATE TABLE ua_blacklist (
    ua VARCHAR(255) PRIMARY KEY
);

INSERT INTO ua_blacklist (ua)
VALUES 
    ('A6-Indexer'),
    ('Googlebot-News'),
    ('Googlebot');
