import { PlusOutlined } from '@ant-design/icons';
import { Button, Input, Select, Space } from 'antd';
import React from 'react';

interface DeviceFiltersProps {
    onSearch: (value: string) => void;
    onStatusChange: (value?: boolean) => void;
    onAddClick: () => void;
}

export const DeviceFilters: React.FC<DeviceFiltersProps> = ({
    onSearch,
    onStatusChange,
    onAddClick
}) => {
    return (
        <Space
            style={{
                marginBottom: 16,
                display: 'flex',
                justifyContent: 'space-between',
                width: '100%'
            }}
        >
            <Space>
                <Input.Search
                    placeholder="Поиск по hostname"
                    allowClear
                    onSearch={onSearch}
                    style={{ width: 250 }}
                />
                <Select
                    placeholder="Фильтр по статусу"
                    allowClear
                    style={{ width: 150 }}
                    onChange={(val) => onStatusChange(val === undefined ? undefined : val)}
                >
                    <Select.Option value={true}>Активные</Select.Option>
                    <Select.Option value={false}>Отключенные</Select.Option>
                </Select>
            </Space>
            <Button type="primary" icon={<PlusOutlined />} onClick={onAddClick}>
                Добавить устройство
            </Button>
        </Space>
    );
};
