--
-- PostgreSQL database dump
--

-- Dumped from database version 12.2 (Debian 12.2-1.pgdg100+1)
-- Dumped by pg_dump version 12.2 (Debian 12.2-1.pgdg100+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: converttvkdau(text); Type: FUNCTION; Schema: public; Owner: user
--

CREATE FUNCTION public.converttvkdau(x text) RETURNS text
    LANGUAGE plpgsql
    AS $$
DECLARE
 cdau text; kdau text; r text;
BEGIN
 cdau = 'áàảãạâấầẩẫậăắằẳẵặđéèẻẽẹêếềểễệíìỉĩịóòỏõọôốồổỗộơớờởỡợúùủũụưứừửữựýỳỷỹỵÁÀẢÃẠÂẤẦẨẪẬĂẮẰẲẴẶĐÉÈẺẼẸÊẾỀỂỄỆÍÌỈĨỊÓÒỎÕỌÔỐỒỔỖỘƠỚỜỞỠỢÚÙỦŨỤƯỨỪỬỮỰÝỲỶỸỴ';
 kdau = 'aaaaaaaaaaaaaaaaadeeeeeeeeeeeiiiiiooooooooooooooooouuuuuuuuuuuyyyyyaaaaaaaaaaaaaaaaadeeeeeeeeeeeiiiiiooooooooooooooooouuuuuuuuuuuyyyyy';
 r = x;
 FOR i IN 0..length(cdau)
 LOOP
 r = replace(r, substr(cdau,i,1), substr(kdau,i,1));
 END LOOP;
 RETURN r;
END;
$$;


ALTER FUNCTION public.converttvkdau(x text) OWNER TO "user";

--
-- Name: next_id(text); Type: FUNCTION; Schema: public; Owner: user
--

CREATE FUNCTION public.next_id(OUT result bigint, seq text) RETURNS bigint
    LANGUAGE plpgsql
    AS $$
DECLARE
    our_epoch bigint := 1314220021721;
    seq_id bigint;
    now_millis bigint;
    shard_id int := 5;
BEGIN
    SELECT nextval(seq) % 1024 INTO seq_id;
    SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp())) INTO now_millis;
    result := (now_millis - our_epoch)*1000 << 23;
    result := result | (shard_id <<10);
    result := result | (seq_id);
    
END;
    $$;


ALTER FUNCTION public.next_id(OUT result bigint, seq text) OWNER TO "user";

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: admin_users; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.admin_users (
    id bigint NOT NULL,
    email text NOT NULL,
    first_name text,
    last_name text,
    password text,
    username text,
    role text,
    last_login timestamp with time zone,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.admin_users OWNER TO "user";

--
-- Name: admin_users_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.admin_users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.admin_users_id_seq OWNER TO "user";

--
-- Name: admin_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: user
--

ALTER SEQUENCE public.admin_users_id_seq OWNED BY public.admin_users.id;


--
-- Name: api_keys; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.api_keys (
    id bigint DEFAULT public.next_id('api_keys_id_seq'::text) NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    value text,
    type text
);


ALTER TABLE public.api_keys OWNER TO "user";

--
-- Name: api_keys_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.api_keys_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.api_keys_id_seq OWNER TO "user";

--
-- Name: consumers; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.consumers (
    id bigint DEFAULT public.next_id('consumers_id_seq'::text) NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    email text,
    name text,
    phone text,
    address text,
    password text
);


ALTER TABLE public.consumers OWNER TO "user";

--
-- Name: consumers_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.consumers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.consumers_id_seq OWNER TO "user";

--
-- Name: customers; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.customers (
    id bigint DEFAULT public.next_id('customers_id_seq'::text) NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    start_time timestamp without time zone,
    end_time timestamp without time zone,
    email text,
    name text,
    phone text,
    address text
);


ALTER TABLE public.customers OWNER TO "user";

--
-- Name: customers_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.customers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.customers_id_seq OWNER TO "user";

--
-- Name: migrations; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.migrations (
    id character varying(255) NOT NULL
);


ALTER TABLE public.migrations OWNER TO "user";

--
-- Name: movie_session_in_theaters; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.movie_session_in_theaters (
    id bigint DEFAULT public.next_id('movie_session_in_theaters_id_seq'::text) NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    start_time timestamp without time zone,
    end_time timestamp without time zone,
    movie_id bigint,
    theater_id bigint
);


ALTER TABLE public.movie_session_in_theaters OWNER TO "user";

--
-- Name: movie_session_in_theaters_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.movie_session_in_theaters_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.movie_session_in_theaters_id_seq OWNER TO "user";

--
-- Name: movies; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.movies (
    id bigint DEFAULT public.next_id('movies_id_seq'::text) NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    name text,
    description text,
    image text,
    trailer text,
    started_at timestamp without time zone,
    duration integer,
    rating double precision,
    type integer
);


ALTER TABLE public.movies OWNER TO "user";

--
-- Name: movies_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.movies_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.movies_id_seq OWNER TO "user";

--
-- Name: payment_partners; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.payment_partners (
    id bigint DEFAULT public.next_id('payment_partners_id_seq'::text) NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    name text,
    api_key_id bigint,
    type integer
);


ALTER TABLE public.payment_partners OWNER TO "user";

--
-- Name: payment_partners_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.payment_partners_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.payment_partners_id_seq OWNER TO "user";

--
-- Name: payments; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.payments (
    id bigint DEFAULT public.next_id('payments_id_seq'::text) NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    amount integer,
    ticket_id bigint,
    payment_partner_id bigint
);


ALTER TABLE public.payments OWNER TO "user";

--
-- Name: payments_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.payments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.payments_id_seq OWNER TO "user";

--
-- Name: products; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.products (
    id bigint DEFAULT public.next_id('products_id_seq'::text) NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    name text,
    price bigint
);


ALTER TABLE public.products OWNER TO "user";

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.products_id_seq OWNER TO "user";

--
-- Name: qor_admin_settings; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.qor_admin_settings (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    key text,
    resource text,
    user_id text,
    value text
);


ALTER TABLE public.qor_admin_settings OWNER TO "user";

--
-- Name: qor_admin_settings_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.qor_admin_settings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qor_admin_settings_id_seq OWNER TO "user";

--
-- Name: qor_admin_settings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: user
--

ALTER SEQUENCE public.qor_admin_settings_id_seq OWNED BY public.qor_admin_settings.id;


--
-- Name: theaters; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.theaters (
    id bigint DEFAULT public.next_id('theaters_id_seq'::text) NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    name text,
    state text,
    city text,
    district text,
    ward text,
    street text,
    description text
);


ALTER TABLE public.theaters OWNER TO "user";

--
-- Name: theaters_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.theaters_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.theaters_id_seq OWNER TO "user";

--
-- Name: tickets; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.tickets (
    id bigint DEFAULT public.next_id('tickets_id_seq'::text) NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    value integer,
    number_seat text,
    movie_session_in_theater_id bigint,
    type integer,
    number_theater integer,
    customer_id bigint
);


ALTER TABLE public.tickets OWNER TO "user";

--
-- Name: tickets_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.tickets_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tickets_id_seq OWNER TO "user";

--
-- Name: admin_users id; Type: DEFAULT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.admin_users ALTER COLUMN id SET DEFAULT nextval('public.admin_users_id_seq'::regclass);


--
-- Name: qor_admin_settings id; Type: DEFAULT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.qor_admin_settings ALTER COLUMN id SET DEFAULT nextval('public.qor_admin_settings_id_seq'::regclass);


--
-- Data for Name: admin_users; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.admin_users (id, email, first_name, last_name, password, username, role, last_login, created_at, updated_at, deleted_at) FROM stdin;
12365	root@gmail.com	root	root	$2a$10$mWh3vR3bCYY6u2HtX3mc5uiKwzzlIoRDLm6pawtqLhhuNE36J9PvC	root	root	\N	2020-02-25 09:09:44.377641+00	2020-02-25 09:09:44.374956+00	\N
2	trungtin2qn1@gmail.com	Tin	Huynh	$2a$10$zU7RybNU.pq5vA6/dijfh.AGwxOHxzDbUzxJduM0JApu1rl4bDwbu	trungtin2qn1	root	\N	2020-02-25 09:11:01.173871+00	2020-02-25 09:11:01.173871+00	\N
\.


--
-- Data for Name: api_keys; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.api_keys (id, created_at, updated_at, deleted_at, value, type) FROM stdin;
1506949409907151873	2020-02-28 05:13:57.944503	2020-02-28 05:13:57.944503	\N	test	normal
\.


--
-- Data for Name: consumers; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.consumers (id, created_at, updated_at, deleted_at, email, name, phone, address, password) FROM stdin;
1506930938192335873	2020-02-27 11:38:43.419276	2020-02-27 11:38:43.419276	\N	trungtin2qn1@gmail.com	Tin Huynh	+84935403649	Quy Nhon, Binh Dinh	$2a$10$uvDMan7UskCFzth2CpwoM.hlFQ7x4QmYlLhQqu5K4WQNE5jfSMXCi
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.customers (id, created_at, updated_at, deleted_at, start_time, end_time, email, name, phone, address) FROM stdin;
\.


--
-- Data for Name: migrations; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.migrations (id) FROM stdin;
init_admin
\.


--
-- Data for Name: movie_session_in_theaters; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.movie_session_in_theaters (id, created_at, updated_at, deleted_at, start_time, end_time, movie_id, theater_id) FROM stdin;
1506920259494351873	2020-02-27 11:17:30.816497	2020-02-27 11:17:30.816497	\N	2020-02-27 12:30:00	2020-02-27 14:30:00	1506919873618383873	1506354062008783873
1507034822713807874	2020-02-27 15:05:06.996461	2020-02-27 15:07:39.716401	\N	2020-02-27 14:30:00	2020-02-29 16:30:00	1506919873618383873	1506354062008783873
1507046910697935875	2020-02-27 15:29:07.961394	2020-02-27 15:29:07.961394	\N	2020-02-27 17:00:00	2020-02-27 19:00:00	1507034118070735874	1506354062008783873
1507047153967567876	2020-02-27 15:29:37.379954	2020-02-27 15:29:37.379954	\N	2020-02-27 18:30:00	2020-02-27 20:30:00	1507034118070735874	1506354062008783873
\.


--
-- Data for Name: movies; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.movies (id, created_at, updated_at, deleted_at, name, description, image, trailer, started_at, duration, rating, type) FROM stdin;
1506919873618383873	2020-02-27 11:16:44.651324	2020-02-27 11:16:44.651324	\N	Movie name	Movie description	movie image link	movie trailer link	2020-03-01 15:30:00	120	8.5	1
1507034118070735874	2020-02-27 15:03:42.835573	2020-02-27 15:03:42.835573	\N	name 1	Description	image	trailer	2020-02-29 00:00:00	97	5.599999904632568	2
1507034352951759875	2020-02-27 15:04:11.259439	2020-02-27 15:04:11.259439	\N	name 2	Description	movie image link	movie trailer link	2020-02-29 00:00:00	98	7.300000190734863	1
\.


--
-- Data for Name: payment_partners; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.payment_partners (id, created_at, updated_at, deleted_at, name, api_key_id, type) FROM stdin;
1507673522938319873	2020-02-28 12:14:06.853056	2020-02-28 12:14:06.853056	\N	momo	1506949409907151873	1
\.


--
-- Data for Name: payments; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.payments (id, created_at, updated_at, deleted_at, amount, ticket_id, payment_partner_id) FROM stdin;
1507685938078159873	2020-02-28 12:38:46.590899	2020-02-28 12:38:46.590899	\N	30	1507680233824719874	1507673522938319873
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.products (id, created_at, updated_at, deleted_at, name, price) FROM stdin;
1505155866796495873	2020-02-25 00:51:58.813	2020-02-25 00:51:58.813	\N	name	20000
1505156017791439874	2020-02-25 00:52:16.8796	2020-02-25 00:52:27.349765	\N	update name 1	25000
1505156965704143875	2020-02-25 00:54:09.723544	2020-02-25 00:54:09.723544	\N	tin huynh	236000
\.


--
-- Data for Name: qor_admin_settings; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.qor_admin_settings (id, created_at, updated_at, deleted_at, key, resource, user_id, value) FROM stdin;
\.


--
-- Data for Name: theaters; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.theaters (id, created_at, updated_at, deleted_at, name, state, city, district, ward, street, description) FROM stdin;
1506354062008783873	2020-02-26 16:32:34.616089	2020-02-27 11:15:49.956434	\N	theater name	state	city	district	ward	street	Description
\.


--
-- Data for Name: tickets; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.tickets (id, created_at, updated_at, deleted_at, value, number_seat, movie_session_in_theater_id, type, number_theater, customer_id) FROM stdin;
1506922323091919873	2020-02-27 11:21:36.800974	2020-02-27 11:21:36.800974	\N	50000	F6	1506920259494351873	1	2	\N
1507683631210959875	2020-02-28 12:34:11.773936	2020-02-28 12:34:11.773936	\N	30	F3	1507047153967567876	2	2	\N
1507680233824719874	2020-02-28 12:27:26.871366	2020-02-28 12:38:46.637936	\N	30	F3	1507047153967567876	1	2	1506930938192335873
\.


--
-- Name: admin_users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.admin_users_id_seq', 2, true);


--
-- Name: api_keys_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.api_keys_id_seq', 1, true);


--
-- Name: consumers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.consumers_id_seq', 1, true);


--
-- Name: customers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.customers_id_seq', 1, false);


--
-- Name: movie_session_in_theaters_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.movie_session_in_theaters_id_seq', 4, true);


--
-- Name: movies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.movies_id_seq', 3, true);


--
-- Name: payment_partners_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.payment_partners_id_seq', 1, true);


--
-- Name: payments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.payments_id_seq', 1, true);


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.products_id_seq', 3, true);


--
-- Name: qor_admin_settings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.qor_admin_settings_id_seq', 1, false);


--
-- Name: theaters_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.theaters_id_seq', 1, true);


--
-- Name: tickets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.tickets_id_seq', 3, true);


--
-- Name: admin_users admin_users_email_key; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.admin_users
    ADD CONSTRAINT admin_users_email_key UNIQUE (email);


--
-- Name: admin_users admin_users_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.admin_users
    ADD CONSTRAINT admin_users_pkey PRIMARY KEY (id);


--
-- Name: api_keys api_keys_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.api_keys
    ADD CONSTRAINT api_keys_pkey PRIMARY KEY (id);


--
-- Name: consumers consumers_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.consumers
    ADD CONSTRAINT consumers_pkey PRIMARY KEY (id);


--
-- Name: customers customers_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);


--
-- Name: migrations migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.migrations
    ADD CONSTRAINT migrations_pkey PRIMARY KEY (id);


--
-- Name: movie_session_in_theaters movie_session_in_theaters_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.movie_session_in_theaters
    ADD CONSTRAINT movie_session_in_theaters_pkey PRIMARY KEY (id);


--
-- Name: movies movies_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);


--
-- Name: payment_partners payment_partners_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.payment_partners
    ADD CONSTRAINT payment_partners_pkey PRIMARY KEY (id);


--
-- Name: payments payments_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: qor_admin_settings qor_admin_settings_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.qor_admin_settings
    ADD CONSTRAINT qor_admin_settings_pkey PRIMARY KEY (id);


--
-- Name: theaters theaters_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.theaters
    ADD CONSTRAINT theaters_pkey PRIMARY KEY (id);


--
-- Name: tickets tickets_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.tickets
    ADD CONSTRAINT tickets_pkey PRIMARY KEY (id);


--
-- Name: idx_qor_admin_settings_deleted_at; Type: INDEX; Schema: public; Owner: user
--

CREATE INDEX idx_qor_admin_settings_deleted_at ON public.qor_admin_settings USING btree (deleted_at);


--
-- Name: movie_session_in_theaters movie_session_in_theaters_movie_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.movie_session_in_theaters
    ADD CONSTRAINT movie_session_in_theaters_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies(id);


--
-- Name: movie_session_in_theaters movie_session_in_theaters_theater_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.movie_session_in_theaters
    ADD CONSTRAINT movie_session_in_theaters_theater_id_fkey FOREIGN KEY (theater_id) REFERENCES public.theaters(id);


--
-- Name: payment_partners payment_partners_api_key_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.payment_partners
    ADD CONSTRAINT payment_partners_api_key_id_fkey FOREIGN KEY (api_key_id) REFERENCES public.api_keys(id);


--
-- Name: payments payments_payment_partner_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_payment_partner_id_fkey FOREIGN KEY (payment_partner_id) REFERENCES public.payment_partners(id);


--
-- Name: payments payments_ticket_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_ticket_id_fkey FOREIGN KEY (ticket_id) REFERENCES public.tickets(id);


--
-- Name: tickets tickets_customer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.tickets
    ADD CONSTRAINT tickets_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES public.consumers(id);


--
-- Name: tickets tickets_movie_session_in_theater_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.tickets
    ADD CONSTRAINT tickets_movie_session_in_theater_id_fkey FOREIGN KEY (movie_session_in_theater_id) REFERENCES public.movie_session_in_theaters(id);


--
-- PostgreSQL database dump complete
--

