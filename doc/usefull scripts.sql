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

-- Рассчет потенциальной компресии по автору
SELECT pageinfo.*, fsinfo.*, ROUND(
        (pageinfo."sum" - fsinfo."sum") * 100 / pageinfo."sum", 2
    )
FROM (
        SELECT ap.value, count(*), SUM(ap."size")
        FROM (
                SELECT bo.value, f."size"
                FROM (
                        SELECT b.id, ba.value
                        FROM
                            books b
                            LEFT JOIN book_attributes ba ON ba.book_id = b.id
                            AND ba.attr = 'author'
                        WHERE
                            ba.book_id IS NOT NULL
                    ) AS bo
                    INNER JOIN pages p ON bo.id = p.book_id
                    INNER JOIN files f ON f.id = p.file_id
            ) AS ap
        GROUP BY
            ap.value
        ORDER BY SUM(ap."size") DESC
    ) AS pageinfo
    LEFT JOIN (
        SELECT ap.value, count(*), SUM(ap."size")
        FROM (
                SELECT bo.value, f."size"
                FROM (
                        SELECT b.id, ba.value
                        FROM
                            books b
                            LEFT JOIN book_attributes ba ON ba.book_id = b.id
                            AND ba.attr = 'author'
                        WHERE
                            ba.book_id IS NOT NULL
                    ) AS bo
                    INNER JOIN pages p ON bo.id = p.book_id
                    INNER JOIN files f ON f.id = p.file_id
                GROUP BY
                    bo.value, f.md5_sum, f.sha256_sum, f."size"
            ) AS ap
        GROUP BY
            ap.value
        ORDER BY SUM(ap."size") DESC
    ) AS fsinfo ON pageinfo.value = fsinfo.value;