-- Active: 1720072909115@@127.0.0.1@5432@hgrabernext
-- Количество значений атрибутов
SELECT COUNT(*), attr, value
FROM book_attributes
GROUP BY
    attr,
    value;

-- Объем и количество файлов дубликатов
SELECT
    COUNT(*) AS pairs,
    SUM(t.c - 1) AS duplicate_count,
    SUM((t.c - 1) * t.s) as duplicate_size
FROM (
        SELECT COUNT(*) AS c, MAX("size") AS s
        FROM files
        GROUP BY
            md5_sum, sha256_sum
        HAVING
            COUNT(*) > 1
    ) AS t;

-- Файлы дубликаты
SELECT f.*
FROM (
        SELECT COUNT(*) AS c, md5_sum, sha256_sum
        FROM files
        GROUP BY
            md5_sum, sha256_sum
        HAVING
            COUNT(*) > 1
    ) AS t
    INNER join files AS f ON f.md5_sum = t.md5_sum
    AND f.sha256_sum = t.sha256_sum
ORDER BY f.id;

-- Не привязанные файлы
SELECT *
FROM files
WHERE
    id NOT IN (
        SELECT file_id
        FROM pages
    );