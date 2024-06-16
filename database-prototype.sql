-- CREATE ROLE hgrabernextuser WITH LOGIN PASSWORD 'hgrabernextpass';
-- CREATE DATABASE hgrabernext OWNER hgrabernextuser;

CREATE TABLE books (
    id UUID PRIMARY KEY,
    name TEXT,
    origin_url TEXT,
    page_count INT,
    attributes_parsed BOOLEAN NOT NULL DEFAULT FALSE,
    create_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE files (
    id UUID PRIMARY KEY,
    filename TEXT NOT NULL,
    ext TEXT NOT NULL,
    md5_sum TEXT,
    sha256_sum TEXT,
    size INT8,
    create_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE pages (
    book_id UUID NOT NULL REFERENCES books (id) ON UPDATE CASCADE ON DELETE CASCADE,
    page_number INT NOT NULL,
    ext TEXT NOT NULL,
    origin_url TEXT,
    create_at TIMESTAMPTZ NOT NULL,
    downloaded BOOL NOT NULL DEFAULT FALSE,
    load_at TIMESTAMPTZ,
    file_id UUID REFERENCES files (id) ON UPDATE CASCADE ON DELETE SET NULL,
    PRIMARY KEY (book_id, page_number)
);

CREATE TABLE attributes (
    code TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    plural_name TEXT NOT NULL,
    order INT NOT NULL DEFAULT 999,
    description TEXT
);

INSERT INTO
    attributes (
        code,
        name,
        plural_name,
        order
    )
VALUES ('tag', 'Тэг', 'Теги', 1),
    (
        'author',
        'Автор',
        'Авторы',
        3
    ),
    (
        'character',
        'Персонаж',
        'Персонажи',
        4
    ),
    (
        'language',
        'Языки',
        'Языки',
        6
    ),
    (
        'category',
        'Категория',
        'Категории',
        2
    ),
    (
        'parody',
        'Пародия',
        'Пародии',
        7
    ),
    (
        'group',
        'Группа',
        'Группы',
        5
    );

CREATE TABLE book_attributes (
    book_id UUID NOT NULL REFERENCES books (id) ON UPDATE CASCADE ON DELETE CASCADE,
    attr TEXT NOT NULL REFERENCES attributes (code) ON UPDATE CASCADE ON DELETE CASCADE,
    value TEXT NOT NULL
);

CREATE TABLE book_labels (
    book_id UUID NOT NULL REFERENCES books (id) ON UPDATE CASCADE ON DELETE CASCADE,
    page_number INT,
    name TEXT NOT NULL,
    value TEXT NOT NULL,
    create_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE agents (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    addr TEXT NOT NULL,
    token TEXT NOT NULL,
    can_parse BOOLEAN NOT NULL DEFAULT FALSE,
    can_export BOOLEAN NOT NULL DEFAULT FALSE,
    priority INT NOT NULL DEFAULT 0,
    create_at TIMESTAMPTZ NOT NULL
);