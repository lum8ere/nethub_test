import { message } from 'antd';

import { useCallback, useState } from 'react';
import { Device } from 'shared/types/device';
import { deviceApi } from '../api/device.api';

export const useDevices = () => {
    const [devices, setDevices] = useState<Device[]>([]);
    const [loading, setLoading] = useState(false);
    const [total, setTotal] = useState(0);

    const fetchDevices = useCallback(async (params: any) => {
        setLoading(true);
        try {
            const { data } = await deviceApi.list(params);
            setDevices(data.data || []);
            setTotal(data.total || 0);
        } catch (err) {
            message.error('Ошибка загрузки');
        } finally {
            setLoading(false);
        }
    }, []);

    return { devices, loading, total, fetchDevices };
};
