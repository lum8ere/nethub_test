-- Создаем дефолтную платформу и локацию
INSERT INTO platforms (code, name) VALUES ('LINUX', 'Linux OS'), ('WINDOWS', 'Windows OS'), ('MAC', 'Mac OS');
INSERT INTO locations (name, is_active) VALUES ('Datacenter-1', true);