import { apiClient } from 'shared/api/client';
import { Location } from 'shared/types/location';

export const locationApi = {
    listAll: () => apiClient.get<{ data: Location[] }>('/locations?limit=1000')
};
