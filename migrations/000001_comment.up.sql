CREATE TABLE comments (
                          id         BIGSERIAL PRIMARY KEY,
                          parent_id  BIGINT REFERENCES comments(id) ON DELETE CASCADE,
                          body       TEXT NOT NULL,
                          created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
