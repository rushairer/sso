CREATE TABLE
    IF NOT EXISTS users (
        id binary(16) NOT NULL,
        email varchar(255) NOT NULL,
        `password` varchar(255) NULL DEFAULT NULL,
        `status` tinyint NOT NULL DEFAULT 0,
        created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        email_verified_at timestamp NULL DEFAULT NULL,
        PRIMARY KEY (id),
        UNIQUE KEY users_email_unique (email)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;