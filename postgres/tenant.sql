-- create a tenant table

CREATE TABLE tenant (
    id serial PRIMARY KEY,
    name_space VARCHAR(12),
    db_host VARCHAR(32),
    db_user VARCHAR(32),
    db_password VARCHAR(32),
    db_name VARCHAR(32),
    allowed_origin VARCHAR(64),
    manager_role VARCHAR(32),
    userRole VARCHAR(32),
    keycloak_client_id VARCHAR(32),
    keycloak_server VARCHAR(128),
    keycloak_jwks_url VARCHAR(128),
    created_at bigint not null default extract(epoch from now())
);

INSERT INTO tenant (name_space, db_host, db_user, db_password, db_name, allowed_origin, manager_role, userRole, keycloak_client_id, keycloak_server, keycloak_jwks_url)
VALUES
    ('host1', 'host1', 'user1', 'password1', 'database1', 'origin1', 'manager1', 'user1', 'client_id1', 'keycloak_server1', 'jwks_url1'),
    ('host2', 'host2', 'user2', 'password2', 'database2', 'origin2', 'manager2', 'user2', 'client_id2', 'keycloak_server2', 'jwks_url2'),
    ('host3', 'host3', 'user3', 'password3', 'database3', 'origin3', 'manager3', 'user3', 'client_id3', 'keycloak_server3', 'jwks_url3'),
    ('host4', 'host4', 'user4', 'password4', 'database4', 'origin4', 'manager4', 'user4', 'client_id4', 'keycloak_server4', 'jwks_url4');
