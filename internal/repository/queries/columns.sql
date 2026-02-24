SELECT a.attname                                       AS "Column Name",
       pg_catalog.format_type(a.atttypid, a.atttypmod) AS "Data type",
       a.attlen                                        AS "Max Length",
       a.atttypmod                                     AS "Precision",
       NOT a.attnotnull                                AS is_nullable,
       COALESCE(ix.indisprimary, false)                AS "Primary Key"
FROM pg_attribute a
         JOIN pg_class t ON a.attrelid = t.oid
         JOIN pg_namespace n ON n.oid = t.relnamespace
         LEFT JOIN pg_index ix
                   ON ix.indrelid = t.oid
                       AND a.attnum = ANY (ix.indkey)
                       AND ix.indisprimary
WHERE t.relname = $1
  AND a.attnum > 0
  AND NOT a.attisdropped
ORDER BY a.attnum;
