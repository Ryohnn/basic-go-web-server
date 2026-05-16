--
-- PostgreSQL database dump
--

\restrict nVjuHfOouLl15a60lVpnZn0gJlZM6q0i0SSSAdtnc7rlaITIg77ZB6X1Nj1Q0pa

-- Dumped from database version 18.3 (Debian 18.3-1.pgdg13+1)
-- Dumped by pg_dump version 18.3 (Debian 18.3-1.pgdg13+1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: games; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.games (
    id integer NOT NULL,
    title character varying(100)
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
-- PostgreSQL database dump complete
--

\unrestrict nVjuHfOouLl15a60lVpnZn0gJlZM6q0i0SSSAdtnc7rlaITIg77ZB6X1Nj1Q0pa

