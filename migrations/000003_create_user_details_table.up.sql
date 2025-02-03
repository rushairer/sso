CREATE TABLE
    IF NOT EXISTS user_details (
        account_id BINARY(16) NOT NULL,
        nick_name VARCHAR(50) NOT NULL,
        avatar VARCHAR(255) DEFAULT NULL,
        gender TINYINT DEFAULT 0 COMMENT '0: unknown, 1: male, 2: female',
        birthday DATE DEFAULT NULL,
        bio TEXT DEFAULT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (account_id),
        UNIQUE KEY user_details_nick_name_unique (nick_name),
        CONSTRAINT fk_user_details_account FOREIGN KEY (account_id) REFERENCES accounts (id) ON DELETE CASCADE
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;