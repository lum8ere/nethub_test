export interface Device {
    id?: string;
    hostname: string;
    ip: string;
    location?: string;
    platform_code: string;
    is_active: boolean;
    created_at?: string;
}

export interface PaginatedResponse<T> {
    data: T[];
    total: number;
    page: number;
    limit: number;
}
