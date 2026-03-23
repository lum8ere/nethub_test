CREATE TABLE IF NOT EXISTS audit_logs (
    id BIGSERIAL PRIMARY KEY,
    operation VARCHAR(50),
    table_name VARCHAR(50),
    record_id BIGINT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);