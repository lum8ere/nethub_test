import { apiClient } from 'shared/api/client';
import { Platform } from 'shared/types/platform';

export const platformApi = {
    listAll: () => apiClient.get<Platform[]>('/platforms')
};
