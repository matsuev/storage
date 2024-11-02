-- Table: public.user

CREATE TABLE IF NOT EXISTS public."user"
(
    id bigint NOT NULL DEFAULT nextval('user_id_seq'::regclass),
    username character varying(50) COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
    email character varying(100) COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
    CONSTRAINT user_pkey PRIMARY KEY (id),
    CONSTRAINT user_username_key UNIQUE (username)
)

TABLESPACE pg_default;
