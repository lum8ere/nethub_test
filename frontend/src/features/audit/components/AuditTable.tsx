import { Table, Tag, Typography } from 'antd';
import dayjs from 'dayjs';
import React from 'react';
import { AuditLog } from 'shared/types/audit';

interface AuditTableProps {
    logs: AuditLog[];
    loading: boolean;
    total: number;
    onPageChange: (page: number) => void;
}

export const AuditTable: React.FC<AuditTableProps> = ({ logs, loading, total, onPageChange }) => {
    const columns = [
        {
            title: 'Дата и время',
            dataIndex: 'created_at',
            key: 'created_at',
            render: (date: string) => dayjs(date).format('DD.MM.YYYY HH:mm:ss')
        },
        {
            title: 'Операция',
            dataIndex: 'operation',
            key: 'operation',
            render: (op: string) => (
                <Tag color={op === 'CREATE' ? 'green' : op === 'UPDATE' ? 'blue' : 'red'}>{op}</Tag>
            )
        },
        {
            title: 'Таблица',
            dataIndex: 'table_name',
            key: 'table_name',
            render: (name: string) => <Typography.Text code>{name}</Typography.Text>
        },
        {
            title: 'ID записи',
            dataIndex: 'record_id',
            key: 'record_id',
            render: (id: string) => (
                <Typography.Text type="secondary" style={{ fontSize: '12px' }}>
                    {id}
                </Typography.Text>
            )
        }
    ];

    return (
        <Table
            dataSource={logs}
            columns={columns}
            rowKey="id"
            loading={loading}
            pagination={{
                total: total,
                pageSize: 20,
                onChange: onPageChange
            }}
        />
    );
};
