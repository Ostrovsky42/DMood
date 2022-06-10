CREATE TABLE IF NOT EXISTS public.users
(
    user_id bigint NOT NULL,
    setting json NOT NULL DEFAULT '{}'::json,
    user_name character varying COLLATE pg_catalog."default" NOT NULL DEFAULT 'Прекрасный человек'::character varying,
    enabled_statistic boolean NOT NULL DEFAULT true,
    request_time smallint DEFAULT 20,
    CONSTRAINT users_pkey PRIMARY KEY (user_id)
)

CREATE TABLE IF NOT EXISTS public.mood
(
    user_id integer,
    date date NOT NULL,
    mood_rating integer,
    discription character varying COLLATE pg_catalog."default",
    day_idea character varying COLLATE pg_catalog."default",
    CONSTRAINT fk_user FOREIGN KEY (user_id)
        REFERENCES public.users (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)