--
-- PostgreSQL database dump
--

\restrict 35KYz4xSIdaeyqZb3ouz6fANQhiWm9ZfmXaQ1kcB2iqikTxQKgsJhTiBJ8ixekX

-- Dumped from database version 18.4 (Debian 18.4-1.pgdg13+1)
-- Dumped by pg_dump version 18.4 (Debian 18.4-1.pgdg13+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: update_modified_column(); Type: FUNCTION; Schema: public; Owner: root
--

CREATE FUNCTION public.update_modified_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;

$$;


ALTER FUNCTION public.update_modified_column() OWNER TO root;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: games; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.games (
    id integer NOT NULL,
    title character varying(100),
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);


ALTER TABLE public.games OWNER TO root;

--
-- Name: games_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.games_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.games_id_seq OWNER TO root;

--
-- Name: games_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.games_id_seq OWNED BY public.games.id;


--
-- Name: games id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.games ALTER COLUMN id SET DEFAULT nextval('public.games_id_seq'::regclass);


--
-- Name: games update_games_modtime; Type: TRIGGER; Schema: public; Owner: root
--

CREATE TRIGGER update_games_modtime BEFORE UPDATE ON public.games FOR EACH ROW EXECUTE FUNCTION public.update_modified_column();


--
-- PostgreSQL database dump complete
--

\unrestrict 35KYz4xSIdaeyqZb3ouz6fANQhiWm9ZfmXaQ1kcB2iqikTxQKgsJhTiBJ8ixekX
