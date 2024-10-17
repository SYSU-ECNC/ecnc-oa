CREATE TABLE
    IF NOT EXISTS users (
        id UUID PRIMARY KEY,
        username TEXT NOT NULL UNIQUE,
        password_hash TEXT NOT NULL,
        full_name TEXT NOT NULL,
        role TEXT NOT NULL,
        is_active BOOLEAN NOT NULL DEFAULT TRUE,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
    );

INSERT INTO
    users (id, username, password_hash, full_name, role)
VALUES
    (
        'f03cd9cc-9383-49b3-84f9-cfea75074e82',
        'admin',
        '$2a$10$EXiXrDu1nb1.Gu/gSnIPCu/aUWAaTCmNRtR8U.WP2JPc1bbfQCib2',
        '初始管理员',
        '黑心'
    );