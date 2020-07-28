CREATE TABLE admins (
    id SERIAL PRIMARY KEY NOT NULL,
    code character varying NOT NULL,
    name character varying NOT NULL,
    mobile character varying NOT NULL,
    encrypted_password character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

COMMENT ON COLUMN admins.code IS '编号';
COMMENT ON COLUMN admins.name IS '姓名';
COMMENT ON COLUMN admins.mobile IS '手机';
COMMENT ON COLUMN admins.encrypted_password IS '密码';
