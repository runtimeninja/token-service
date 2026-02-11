CREATE TABLE IF NOT EXISTS api_keys (
  id          UUID PRIMARY KEY,
  owner       TEXT NOT NULL,
  name        TEXT NOT NULL,
  token_hash  TEXT NOT NULL,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
  revoked_at  TIMESTAMPTZ NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS ux_api_keys_owner_name
ON api_keys(owner, name)
WHERE revoked_at IS NULL;

CREATE INDEX IF NOT EXISTS ix_api_keys_token_hash_active
ON api_keys(token_hash)
WHERE revoked_at IS NULL;