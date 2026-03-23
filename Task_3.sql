-- список активных устройств с количеством связанных configs
SELECT 
    d.id, 
    d.hostname, 
    COUNT(c.id) as configs_count
FROM devices d
LEFT JOIN configs c ON d.id = c.device_id
WHERE d.is_active = TRUE 
  AND d.deleted_at IS NULL
GROUP BY d.id, d.hostname;

-- CREATE INDEX idx_configs_device_id ON configs(device_id);

-- Либо через подзапрос, если записей миллион
SELECT 
    d.id, 
    d.hostname, 
    COALESCE(conf.total, 0) as configs_count
FROM devices d
LEFT JOIN (
    SELECT device_id, COUNT(*) as total 
    FROM configs 
    GROUP BY device_id
) conf ON d.id = conf.device_id
WHERE d.is_active = TRUE 
  AND d.deleted_at IS NULL;

-- последних N записей из logs для конкретного устройства;
SELECT * 
FROM logs 
WHERE device_id = 123
ORDER BY created_at DESC 
LIMIT 10;

-- CREATE INDEX idx_logs_device_id_created_at ON logs(device_id, created_at DESC);