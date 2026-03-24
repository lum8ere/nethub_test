export interface AuditLog {
    id: string;
    operation: 'CREATE' | 'UPDATE' | 'DELETE';
    table_name: string;
    record_id: string;
    created_at: string;
}
