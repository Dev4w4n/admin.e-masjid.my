-- create a tenant table

CREATE TABLE tenant (
    id serial PRIMARY KEY,
    db_host VARCHAR(32),
    db_user VARCHAR(32),
    db_password VARCHAR(32),
    db_name VARCHAR(32),
    allowed_origin VARCHAR(64),
    manager_role VARCHAR(32),
    userRole VARCHAR(32),
    keycloak_client_id VARCHAR(32),
    keycloak_server VARCHAR(128),
    keycloak_jwks_url VARCHAR(128)
);
