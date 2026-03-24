import { message } from 'antd';
import { deviceApi } from 'features/devices/api/device.api';
import { DeviceFilters } from 'features/devices/components/DeviceFilters';
import { DeviceModal } from 'features/devices/components/DeviceModal';
import { DeviceTable } from 'features/devices/components/DeviceTable';
import React, { useCallback, useEffect, useState } from 'react';
import { Device } from 'shared/types/device';

export const DevicesPage: React.FC = () => {
    const [devices, setDevices] = useState<Device[]>([]);
    const [loading, setLoading] = useState(false);
    const [total, setTotal] = useState(0);
    const [currentPage, setCurrentPage] = useState(1);

    const [search, setSearch] = useState('');
    const [status, setStatus] = useState<boolean | undefined>(undefined);

    const [isModalOpen, setIsModalOpen] = useState(false);
    const [selectedDevice, setSelectedDevice] = useState<Device | null>(null);
    const [submitLoading, setSubmitLoading] = useState(false);

    const loadData = useCallback(
        async (page: number = 1) => {
            setLoading(true);
            try {
                const response = await deviceApi.list({
                    hostname: search,
                    is_active: status,
                    page: page,
                    limit: 10
                });
                setDevices(response.data.data || []);
                setTotal(response.data.total || 0);
                setCurrentPage(page);
            } catch (error) {
                console.error(error);
                message.error('Не удалось загрузить список устройств');
            } finally {
                setLoading(false);
            }
        },
        [search, status]
    );

    useEffect(() => {
        loadData(1);
    }, [loadData]);

    const handleAdd = () => {
        setSelectedDevice(null);
        setIsModalOpen(true);
    };

    const handleEdit = (device: Device) => {
        setSelectedDevice(device);
        setIsModalOpen(true);
    };

    const handleDelete = async (id: string) => {
        try {
            await deviceApi.delete(id);
            message.success('Устройство успешно удалено');
            loadData(currentPage);
        } catch (error) {
            message.error('Ошибка при удалении устройства');
        }
    };

    const handleSave = async (values: Device) => {
        setSubmitLoading(true);
        try {
            if (selectedDevice?.id) {
                await deviceApi.update(selectedDevice.id, values);
                message.success('Устройство обновлено');
            } else {
                await deviceApi.create(values);
                message.success('Устройство успешно создано');
            }
            setIsModalOpen(false);
            loadData(currentPage);
        } catch (error) {
            message.error('Ошибка при сохранении данных');
        } finally {
            setSubmitLoading(false);
        }
    };

    return (
        <div className="devices-page">
            <DeviceFilters
                onSearch={(val) => {
                    setSearch(val);
                }}
                onStatusChange={(val) => {
                    setStatus(val);
                }}
                onAddClick={handleAdd}
            />
            <DeviceTable
                devices={devices}
                loading={loading}
                total={total}
                currentPage={currentPage}
                onPageChange={(page) => loadData(page)}
                onEdit={handleEdit}
                onDelete={handleDelete}
            />
            <DeviceModal
                open={isModalOpen}
                device={selectedDevice}
                confirmLoading={submitLoading}
                onClose={() => setIsModalOpen(false)}
                onSave={handleSave}
            />
        </div>
    );
};
