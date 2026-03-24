import { DeleteOutlined, EditOutlined } from '@ant-design/icons';
import { Button, Popconfirm, Space, Table, Tag } from 'antd';
import React from 'react';
import { Device } from 'shared/types/device';

interface DeviceTableProps {
    devices: Device[];
    loading: boolean;
    total: number;
    currentPage: number;
    onPageChange: (page: number) => void;
    onEdit: (device: Device) => void;
    onDelete: (id: string) => void;
}

export const DeviceTable: React.FC<DeviceTableProps> = ({
    devices,
    loading,
    total,
    currentPage,
    onPageChange,
    onEdit,
    onDelete
}) => {
    const columns = [
        {
            title: 'Hostname',
            dataIndex: 'hostname',
            key: 'hostname',
            render: (text: string) => <b>{text}</b>
        },
        { title: 'IP Адрес', dataIndex: 'ip', key: 'ip' },
        {
            title: 'Платформа',
            dataIndex: 'platform_code',
            key: 'platform_code',
            render: (code: string) => <Tag color="blue">{code}</Tag>
        },
        {
            title: 'Статус',
            dataIndex: 'is_active',
            key: 'is_active',
            render: (isActive: boolean) => (
                <Tag color={isActive ? 'green' : 'red'}>{isActive ? 'Активен' : 'Отключен'}</Tag>
            )
        },
        {
            title: 'Действия',
            key: 'actions',
            width: 100,
            render: (_: any, record: Device) => (
                <Space>
                    <Button type="text" icon={<EditOutlined />} onClick={() => onEdit(record)} />
                    <Popconfirm
                        title="Удалить устройство?"
                        description="Это действие нельзя отменить (мягкое удаление)."
                        onConfirm={() => record.id && onDelete(record.id)}
                        okText="Да"
                        cancelText="Нет"
                    >
                        <Button type="text" danger icon={<DeleteOutlined />} />
                    </Popconfirm>
                </Space>
            )
        }
    ];

    return (
        <Table
            columns={columns}
            dataSource={devices}
            rowKey="id"
            loading={loading}
            pagination={{
                current: currentPage,
                total: total,
                pageSize: 10,
                onChange: onPageChange,
                showSizeChanger: false
            }}
        />
    );
};
