--
-- PostgreSQL database dump
--

-- Dumped from database version 10.0
-- Dumped by pg_dump version 10.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: entitlement_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE entitlement_groups (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE entitlement_groups OWNER TO postgres;

--
-- Name: entitlements; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE entitlements (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    permanent boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE entitlements OWNER TO postgres;

--
-- Name: grants; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE grants (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    entitlement_group_id uuid NOT NULL,
    entitlement_id uuid NOT NULL,
    allow boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE grants OWNER TO postgres;

--
-- Name: identity_entitlements; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE identity_entitlements (
    id uuid NOT NULL,
    identity character varying(255) NOT NULL,
    entitlement_id uuid NOT NULL,
    allow boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE identity_entitlements OWNER TO postgres;

--
-- Name: identity_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE identity_groups (
    id uuid NOT NULL,
    identity character varying(255) NOT NULL,
    entitlement_group_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE identity_groups OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE schema_migration OWNER TO postgres;

--
-- Name: entitlement_groups entitlement_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY entitlement_groups
    ADD CONSTRAINT entitlement_groups_pkey PRIMARY KEY (id);


--
-- Name: entitlements entitlements_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY entitlements
    ADD CONSTRAINT entitlements_pkey PRIMARY KEY (id);


--
-- Name: grants grants_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY grants
    ADD CONSTRAINT grants_pkey PRIMARY KEY (id);


--
-- Name: identity_entitlements identity_entitlements_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY identity_entitlements
    ADD CONSTRAINT identity_entitlements_pkey PRIMARY KEY (id);


--
-- Name: identity_groups identity_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY identity_groups
    ADD CONSTRAINT identity_groups_pkey PRIMARY KEY (id);


--
-- Name: version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX version_idx ON schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

