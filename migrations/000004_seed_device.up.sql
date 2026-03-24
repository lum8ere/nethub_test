INSERT INTO devices (hostname, ip, platform_code, location, is_active, name)
VALUES 
('srv-web-01', '192.168.1.10', 'LINUX', (SELECT id FROM locations WHERE name = 'Datacenter-1' LIMIT 1), true, 'Основной веб-сервер'),
('srv-db-prod', '192.168.1.20', 'LINUX', (SELECT id FROM locations WHERE name = 'Datacenter-1' LIMIT 1), true, 'База данных (Prod)'),
('workstation-dev-01', '10.0.0.50', 'WINDOWS', (SELECT id FROM locations WHERE name = 'Datacenter-1' LIMIT 1), false, 'Рабочая станция разработчика'),
('srv-proxy-nginx', '172.16.0.5', 'LINUX', (SELECT id FROM locations WHERE name = 'Datacenter-1' LIMIT 1), true, 'Прокси-сервер');