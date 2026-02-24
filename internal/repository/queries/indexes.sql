SELECT i.relname                                         AS index_name,
       am.amname ||
       CASE WHEN ix.indisunique THEN ', UNIQUE' ELSE '' END ||
       CASE WHEN ix.indisprimary THEN ', PRIMARY KEY' ELSE '' END
                                                         AS index_description,
       STRING_AGG(a.attname, ', ' ORDER BY k.ordinality) AS index_keys
FROM pg_class t
         JOIN pg_index ix ON t.oid = ix.indrelid
         JOIN pg_class i ON i.oid = ix.indexrelid
         JOIN pg_am am ON i.relam = am.oid
         JOIN LATERAL unnest(ix.indkey) WITH ORDINALITY AS k(attnum, ordinality)
ON true
    JOIN pg_attribute a ON a.attrelid = t.oid AND a.attnum = k.attnum
    JOIN pg_namespace n ON n.oid = t.relnamespace
WHERE t.relkind = 'r'
  AND t.relname = $1
GROUP BY i.relname, am.amname, ix.indisunique, ix.indisprimary
ORDER BY i.relname;
