CREATE TABLE users (
    id SERIAL PRIMARY KEY NOT NULL,
    code character varying NOT NULL,
    name character varying NOT NULL,
    mobile character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

COMMENT ON COLUMN users.code IS '编号';
COMMENT ON COLUMN users.name IS '姓名';
COMMENT ON COLUMN users.mobile IS '手机';
