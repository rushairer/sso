CREATE TABLE
    IF NOT EXISTS account_roles (
        account_id BINARY(16) NOT NULL,
        role_id BINARY(16) NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (account_id, role_id),
        KEY account_roles_account_id_index (account_id),
        KEY account_roles_role_id_index (role_id),
        CONSTRAINT fk_account_roles_account FOREIGN KEY (account_id) REFERENCES accounts (id) ON DELETE CASCADE,
        CONSTRAINT fk_account_roles_role FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;