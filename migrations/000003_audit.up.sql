CREATE TABLE IF NOT EXISTS audit_logs (
    id object_id PRIMARY KEY DEFAULT gen_random_uuid(),
    operation TEXT,
    table_name TEXT,
    record_id object_id,
    created_at datetime DEFAULT NOW()
);