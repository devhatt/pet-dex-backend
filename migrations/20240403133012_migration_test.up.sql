CREATE TABLE IF NOT EXISTS users
(
    id         uuid primary key default UUID(),
    name       varchar(120),
    type       varchar(20) check(type IN ('fisica', 'juridica')),
    birthdate  date,
    document VARCHAR(14),
    avatarUrl  varchar(255),
    email      varchar(128) not null,
    pass VARCHAR(100) NOT NULL,
    phone      varchar(12)  not null,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    UNIQUE INDEX idx_email (email)
    );