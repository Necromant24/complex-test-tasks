--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2
-- Dumped by pg_dump version 17.2

-- Started on 2024-12-31 01:48:45

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

DROP DATABASE IF EXISTS "complexTest";
--
-- TOC entry 4809 (class 1262 OID 16405)
-- Name: complexTest; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE "complexTest" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Russian_Russia.1251';


ALTER DATABASE "complexTest" OWNER TO postgres;

\connect "complexTest"

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
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 4810 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 218 (class 1259 OID 16427)
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id uuid NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.products OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16434)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16420)
-- Name: users-requests; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."users-requests" (
    id uuid NOT NULL,
    "userId" uuid NOT NULL,
    "productId" uuid NOT NULL,
    created date NOT NULL,
    status character varying
);


ALTER TABLE public."users-requests" OWNER TO postgres;

--
-- TOC entry 4653 (class 2606 OID 16433)
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- TOC entry 4650 (class 2606 OID 16426)
-- Name: users-requests users-requests_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."users-requests"
    ADD CONSTRAINT "users-requests_pkey" PRIMARY KEY (id);


--
-- TOC entry 4656 (class 2606 OID 16440)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4651 (class 1259 OID 16451)
-- Name: index_products; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX index_products ON public.products USING btree (id);


--
-- TOC entry 4654 (class 1259 OID 16452)
-- Name: index_users; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX index_users ON public.users USING btree (id);


--
-- TOC entry 4648 (class 1259 OID 16453)
-- Name: index_users-requests; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "index_users-requests" ON public."users-requests" USING btree (id);


--
-- TOC entry 4657 (class 2606 OID 16446)
-- Name: users-requests productForeignKey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."users-requests"
    ADD CONSTRAINT "productForeignKey" FOREIGN KEY ("productId") REFERENCES public.products(id) NOT VALID;


--
-- TOC entry 4658 (class 2606 OID 16441)
-- Name: users-requests userForeignKey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."users-requests"
    ADD CONSTRAINT "userForeignKey" FOREIGN KEY ("userId") REFERENCES public.users(id) NOT VALID;


-- Completed on 2024-12-31 01:48:45

--
-- PostgreSQL database dump complete
--

