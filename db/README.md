# sql

## create database
```
CREATE DATABASE name;
```

## drop database
```
ROP DATABASE name;
```

## create table
```
CREATE TABLE table_name (
  id SERIAL PRIMARY KEY NOT NULL,
  code character varying NOT NULL,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
);
```

## drop table
```
DROP TABLE table_name;
```

## table column comment
```
COMMENT ON COLUMN table_name.column IS 'xxxxx';
```

## add table column
```
ALTER TABLE table_name
ADD COLUMN new_column_name data_type,
ADD COLUMN new_column_name data_type;
```

## rename table column
```
ALTER TABLE table_name
RENAME column_name TO new_column_name;
```

## change table column type
```
ALTER TABLE table_name
ALTER COLUMN column_name new_data_type;
```

## change table column not null
```
ALTER TABLE table_name
ALTER COLUMN column_name SET NOT NULL;
```

## change table column default
```
ALTER TABLE table_name
ALTER COLUMN column_name SET DEFAULT value;
```

## drop table column
```
ALTER TABLE table_name
DROP COLUMN column_name;
```

## add index
```
CREATE INDEX index_name ON table_name
[USING method]
(
  column_name [ASC | DESC] [NULLS {FIRST | LAST }],
  ...
);

example: CREATE UNIQUE INDEX index_name ON table_name USING btree (column_name, column_name);
```
