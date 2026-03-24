import { apiClient } from 'shared/api/client';
import { Device, PaginatedResponse } from 'shared/types/device';

export const deviceApi = {
    list: (params: any) => apiClient.get<PaginatedResponse<Device>>('/devices', { params }),
    create: (data: Device) => apiClient.post<Device>('/devices', data),
    update: (id: string, data: Device) => apiClient.put<Device>(`/devices/${id}`, data),
    delete: (id: string) => apiClient.delete(`/devices/${id}`)
};
