DROP TABLE IF EXISTS oauth_clients;
CREATE TABLE oauth_clients (
	id INT AUTO_INCREMENT NOT NULL,
	secret VARCHAR(128),
	isActive BOOLEAN,
	scopes VARCHAR(256),
	PRIMARY KEY (`id`)
);

INSERT INTO oauth_clients
    (id, secret, isActive, scopes)
VALUES
    ('1', 'secret1', false, 'scope1,scope2,scope3'),
    ('2', 'secret2', true, 'scope1,scope2,scope3'),
    ('3', 'secret3', false, 'scope1,scope2,scope3'),
    ('4', 'secret4', true, 'scope1,scope2,scope3');


DROP TABLE IF EXISTS oauth_users;
CREATE TABLE oauth_users (
	id INT AUTO_INCREMENT NOT NULL,
	email VARCHAR(128) NOT NULL,
	password VARCHAR(128) NOT NULL,
	PRIMARY KEY (`id`)

);


INSERT INTO oauth_users
    (id, email, password)
VALUES
    ('1', 'm@gmail.com', 'kkksks'),
    ('2', 's@gmail.com', 'kla;a;'),
    ('3', 'j@gmail.com', 'p.;w92922');



