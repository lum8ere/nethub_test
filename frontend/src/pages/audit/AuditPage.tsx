import { Typography } from 'antd';
import { auditApi } from 'features/audit/api/audit.api';
import { AuditTable } from 'features/audit/components/AuditTable';
import React, { useEffect, useState } from 'react';
import { AuditLog } from 'shared/types/audit';

export const AuditPage: React.FC = () => {
    const [logs, setLogs] = useState<AuditLog[]>([]);
    const [loading, setLoading] = useState(false);
    const [total, setTotal] = useState(0);

    const fetchLogs = async (page = 1) => {
        setLoading(true);
        try {
            const { data } = await auditApi.list(page);
            setLogs(data.data || []);
            setTotal(data.total || 0);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchLogs();
    }, []);

    return (
        <div>
            <Typography.Title level={2}>Журнал аудита</Typography.Title>
            <AuditTable logs={logs} loading={loading} total={total} onPageChange={fetchLogs} />
        </div>
    );
};
