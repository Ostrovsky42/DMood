CREATE TABLE IF NOT EXISTS public.users
(
    user_id integer NOT NULL,
    setting json NOT NULL DEFAULT '{}'::json,
    enabled_statistic bool NOT NULL DEFAULT true 
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