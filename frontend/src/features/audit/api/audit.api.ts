import { apiClient } from 'shared/api/client';
import { AuditLog } from 'shared/types/audit';
import { PaginatedResponse } from 'shared/types/device';

export const auditApi = {
    list: (page = 1, limit = 20) =>
        apiClient.get<PaginatedResponse<AuditLog>>(`/audit?page=${page}&limit=${limit}`)
};
