CREATE DOMAIN "public"."object_id" AS uuid;
CREATE DOMAIN "public"."code" AS TEXT;
CREATE DOMAIN "public"."datetime" AS timestamptz;
CREATE DOMAIN "public"."coordinate" AS point;

CREATE TABLE IF NOT EXISTS users (
    id object_id PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at datetime NOT NULL DEFAULT now(),
    deleted_at datetime,
    created_by object_id,
    username TEXT,
    email TEXT UNIQUE,
    last_name TEXT,
    first_name TEXT,
    middle_name TEXT,
    birth_date date,
    avatar_file_path TEXT
);

-- Это для разделения устройств по системам: WINDOWS, IOS, LINUX, lunux aka какая-нибудь русская обвертка и тд.
CREATE TABLE IF NOT EXISTS platforms (
    id object_id PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at datetime NOT NULL DEFAULT now(),
    ver bigint NOT NULL DEFAULT 0,
    deleted_at datetime,
    code code UNIQUE NOT NULL,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS locations (
    id object_id PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at datetime NOT NULL DEFAULT now(),
    created_by object_id REFERENCES users(id),
    deleted_at datetime,
    var bigint NOT NULL DEFAULT 0,
    name TEXT,
    is_active boolean NOT NULL DEFAULT true,
    coordinate coordinate
);

CREATE TABLE IF NOT EXISTS devices (
    id object_id PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at datetime NOT NULL DEFAULT now(),
    created_by object_id REFERENCES users(id),
    deleted_at datetime,
    name TEXT, -- это человеко читаемое

    -- Поля из моего опыта, которые пригодятся для дальнейшего улучшения системы
    user_id object_id REFERENCES users(id) ON DELETE SET NULL,
    platform_code code NOT NULL REFERENCES platforms(code),

    -- поля по ТЗ
    ip TEXT,
    hostname TEXT, -- то, что дается/устанавлявается устройством
    location object_id REFERENCES locations(id) ON DELETE SET NULL,
    is_active BOOLEAN NOT NULL DEFAULT true
);