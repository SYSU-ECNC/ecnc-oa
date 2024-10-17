CREATE TABLE
    IF NOT EXISTS refresh_tokens (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        user_id UUID NOT NULL,
        token TEXT NOT NULL,
        expires_at TIMESTAMPTZ NOT NULL,
        revoke BOOLEAN DEFAULT FALSE,
        created_at TIMESTAMPTZ DEFAULT NOW (),
        updated_at TIMESTAMPTZ DEFAULT NOW (),
        CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );