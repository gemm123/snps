--
-- PostgreSQL database dump
--

-- Dumped from database version 13.4
-- Dumped by pg_dump version 13.4

-- Started on 2024-01-11 16:43:08

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
-- TOC entry 3 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 3053 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 202 (class 1259 OID 27044)
-- Name: cart_products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cart_products (
    id uuid NOT NULL,
    id_cart uuid NOT NULL,
    id_product uuid NOT NULL,
    quantity integer NOT NULL,
    total integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.cart_products OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 27032)
-- Name: carts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.carts (
    id uuid NOT NULL,
    id_user uuid NOT NULL,
    total integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.carts OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 27066)
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id uuid NOT NULL,
    name character varying NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 27101)
-- Name: order_products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_products (
    id uuid NOT NULL,
    id_order uuid NOT NULL,
    id_product uuid NOT NULL,
    quantity integer NOT NULL,
    total integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.order_products OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 27086)
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    id uuid NOT NULL,
    id_user uuid NOT NULL,
    total integer NOT NULL,
    status character varying NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 27056)
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id uuid NOT NULL,
    id_category uuid NOT NULL,
    name character varying NOT NULL,
    description character varying NOT NULL,
    price integer NOT NULL,
    stok integer NOT NULL,
    image character varying NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.products OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 27022)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    email character varying NOT NULL,
    password character varying NOT NULL,
    name character varying NOT NULL,
    address character varying NOT NULL,
    phone character varying NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 3043 (class 0 OID 27044)
-- Dependencies: 202
-- Data for Name: cart_products; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.cart_products VALUES ('3db3198e-4c43-49ed-85e0-58bae81ba5f5', '9d7df7f5-0892-4a89-8cfe-51006c13f904', '9ed52a00-84cb-4f74-9f55-6206c453c3b9', 2, 40000, '2024-01-10 09:11:08.763058+07', '2024-01-10 09:11:08.763058+07');
INSERT INTO public.cart_products VALUES ('6c1cc1a4-c189-41f0-bbcd-6923817c1512', '9d7df7f5-0892-4a89-8cfe-51006c13f904', '8d075149-98d4-413e-829b-a950a0c67ed6', 2, 30000, '2024-01-10 10:01:47.929651+07', '2024-01-10 10:01:47.929651+07');


--
-- TOC entry 3042 (class 0 OID 27032)
-- Dependencies: 201
-- Data for Name: carts; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.carts VALUES ('9d7df7f5-0892-4a89-8cfe-51006c13f904', 'd6cbde8b-9b04-41d4-8cae-49cc1cddaff8', 70000, '2024-01-10 09:11:08.755578+07', '2024-01-10 10:01:47.936056+07');


--
-- TOC entry 3045 (class 0 OID 27066)
-- Dependencies: 204
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.categories VALUES ('59ed8b9f-968f-493d-a775-a6136987d76d', 'Baju', '2024-01-09 16:43:38.40683+07', '2024-01-09 16:43:38.40683+07');
INSERT INTO public.categories VALUES ('66157187-9518-4474-8cef-7b045b5e966e', 'Celana', '2024-01-09 16:43:38.411182+07', '2024-01-09 16:43:38.411182+07');
INSERT INTO public.categories VALUES ('14b1abc9-2612-4821-a0d0-162e54e6ca3c', 'Topi', '2024-01-09 16:43:38.412543+07', '2024-01-09 16:43:38.412543+07');


--
-- TOC entry 3047 (class 0 OID 27101)
-- Dependencies: 206
-- Data for Name: order_products; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.order_products VALUES ('f0bb4a73-959b-4141-be74-9f6891d8e48b', '16d4c03f-dc61-4abd-9797-27b722e7b0f6', '9e09cfe6-a395-44c7-afb9-c50a6bd250ca', 1, 40000, '2024-01-10 22:50:54.548985+07', '2024-01-10 22:50:54.548985+07');
INSERT INTO public.order_products VALUES ('370bb769-842c-4591-a83e-998284c7e8ba', '16d4c03f-dc61-4abd-9797-27b722e7b0f6', '892e1fca-9652-4c7e-9b58-759bd7ff1722', 2, 70000, '2024-01-10 22:50:54.560523+07', '2024-01-10 22:50:54.560523+07');
INSERT INTO public.order_products VALUES ('67412843-d40e-4239-8dd4-744628ae1032', '82c49621-b087-47d9-b69a-cbba57dd523e', '892e1fca-9652-4c7e-9b58-759bd7ff1722', 2, 70000, '2024-01-11 00:08:15.798775+07', '2024-01-11 00:08:15.798775+07');
INSERT INTO public.order_products VALUES ('a0a2e49b-d5ac-4ac9-b97e-0c8b9f6180a5', '9b1f9daa-09e5-47dd-b1be-3d3ee73e2898', '8d075149-98d4-413e-829b-a950a0c67ed6', 1, 15000, '2024-01-11 13:50:01.037527+07', '2024-01-11 13:50:01.037527+07');
INSERT INTO public.order_products VALUES ('76d36a23-7317-4a18-9b67-7c5fe078bf37', '9b1f9daa-09e5-47dd-b1be-3d3ee73e2898', 'efd16e40-99df-41a9-8cbc-18e29ed80a38', 1, 18000, '2024-01-11 13:50:01.043568+07', '2024-01-11 13:50:01.043568+07');


--
-- TOC entry 3046 (class 0 OID 27086)
-- Dependencies: 205
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.orders VALUES ('82c49621-b087-47d9-b69a-cbba57dd523e', '2f3feaab-f6c6-4cfc-8ab3-befc62b430c9', 70000, 'Unpaid', '2024-01-11 00:08:15.795466+07', '2024-01-11 00:08:15.795466+07');
INSERT INTO public.orders VALUES ('16d4c03f-dc61-4abd-9797-27b722e7b0f6', '2f3feaab-f6c6-4cfc-8ab3-befc62b430c9', 110000, 'Paid', '2024-01-10 22:50:54.535882+07', '2024-01-11 00:18:19.453695+07');
INSERT INTO public.orders VALUES ('9b1f9daa-09e5-47dd-b1be-3d3ee73e2898', '49cc698d-dd20-4bd0-a913-fb26df0890db', 33000, 'Paid', '2024-01-11 13:50:01.023912+07', '2024-01-11 13:50:39.935305+07');


--
-- TOC entry 3044 (class 0 OID 27056)
-- Dependencies: 203
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.products VALUES ('9ed52a00-84cb-4f74-9f55-6206c453c3b9', '59ed8b9f-968f-493d-a775-a6136987d76d', 'T-Shirt 1', 'Comfortable cotton T-shirt for everyday wear.', 20000, 50, 'image', '2024-01-09 16:47:32.432072+07', '2024-01-09 16:47:32.432072+07');
INSERT INTO public.products VALUES ('9dc323bf-d39a-46f9-b2c6-e90e49abb9c7', '59ed8b9f-968f-493d-a775-a6136987d76d', 'T-Shirt 2', 'Casual and stylish T-shirt for any occasion.', 22500, 50, 'image', '2024-01-09 16:47:32.436462+07', '2024-01-09 16:47:32.436462+07');
INSERT INTO public.products VALUES ('8da39d30-3b2b-4fc7-8bb8-5dcbac6a8e57', '59ed8b9f-968f-493d-a775-a6136987d76d', 'T-Shirt 3', 'Classic T-shirt with a modern twist.', 25000, 50, 'image', '2024-01-09 16:48:19.514488+07', '2024-01-09 16:48:19.514488+07');
INSERT INTO public.products VALUES ('52a2a978-2172-40ec-9c23-ac2a417d6438', '66157187-9518-4474-8cef-7b045b5e966e', 'Jeans 1', 'Durable denim jeans for a timeless look.', 30000, 50, 'image', '2024-01-09 16:49:04.965636+07', '2024-01-09 16:49:04.965636+07');
INSERT INTO public.products VALUES ('278f9d60-28f3-4707-9b3b-ade3d910dff0', '66157187-9518-4474-8cef-7b045b5e966e', 'Jeans 2', 'Slim-fit jeans for a sleek and stylish appearance.', 32500, 50, 'image', '2024-01-09 16:49:42.741319+07', '2024-01-09 16:49:42.741319+07');
INSERT INTO public.products VALUES ('d50613e6-8b2f-4f2a-bad7-384a490b06ff', '14b1abc9-2612-4821-a0d0-162e54e6ca3c', 'Snapback Hat 3', 'Stylish snapback hat for a casual vibe.', 20000, 50, 'image', '2024-01-09 16:54:41.176349+07', '2024-01-09 16:54:41.176349+07');
INSERT INTO public.products VALUES ('9e09cfe6-a395-44c7-afb9-c50a6bd250ca', '59ed8b9f-968f-493d-a775-a6136987d76d', 'Hoodie 1', 'Warm and cozy hoodie for chilly days.', 40000, 50, 'image', '2024-01-09 16:55:11.346668+07', '2024-01-09 16:55:11.346668+07');
INSERT INTO public.products VALUES ('d196411c-e207-461c-a173-123e2ac25622', '66157187-9518-4474-8cef-7b045b5e966e', 'Chino Pants 1', 'Classic chino pants for a polished look.', 35000, 50, 'image', '2024-01-09 16:55:51.012317+07', '2024-01-09 16:55:51.012317+07');
INSERT INTO public.products VALUES ('d47ad40b-7f09-4005-8ccf-2ab853275378', '66157187-9518-4474-8cef-7b045b5e966e', 'Chino Pants 2', 'Versatile chino pants suitable for any occasion.', 38500, 50, 'image', '2024-01-09 16:56:27.362764+07', '2024-01-09 16:56:27.362764+07');
INSERT INTO public.products VALUES ('8cf49736-fbf7-4d53-ba7c-804043234681', '66157187-9518-4474-8cef-7b045b5e966e', 'Chino Pants 3', 'Modern chino pants with a contemporary style.', 42000, 50, 'image', '2024-01-09 16:57:05.044786+07', '2024-01-09 16:57:05.044786+07');
INSERT INTO public.products VALUES ('843b80a1-4d0e-4ddf-8d57-b1e7ea15971a', '14b1abc9-2612-4821-a0d0-162e54e6ca3c', 'Beanie Hat 1', 'Cozy beanie hat for a laid-back look.', 12000, 50, 'image', '2024-01-09 16:57:40.450508+07', '2024-01-09 16:57:40.450508+07');
INSERT INTO public.products VALUES ('3b6e3ec9-b5ae-4999-be62-b227ab09cf78', '14b1abc9-2612-4821-a0d0-162e54e6ca3c', 'Beanie Hat 2', 'Knit beanie hat to keep you warm in style.', 14000, 50, 'image', '2024-01-09 17:25:31.112637+07', '2024-01-09 17:25:31.112637+07');
INSERT INTO public.products VALUES ('6d97105e-343a-46fe-9d54-53698e8aae92', '14b1abc9-2612-4821-a0d0-162e54e6ca3c', 'Beanie Hat 3', 'Fashionable beanie hat for a trendy outfit.', 16000, 50, 'image', '2024-01-09 17:26:09.327832+07', '2024-01-09 17:26:09.327832+07');
INSERT INTO public.products VALUES ('581bf2c5-c3e3-4888-919a-545b54af68db', '59ed8b9f-968f-493d-a775-a6136987d76d', 'Long Sleeve Shirt 1', 'Elegant long sleeve shirt for a sophisticated look.', 28000, 50, 'image', '2024-01-09 17:26:46.131939+07', '2024-01-09 17:26:46.131939+07');
INSERT INTO public.products VALUES ('e5b6d407-4c98-496f-9141-194feef16d0c', '66157187-9518-4474-8cef-7b045b5e966e', 'Cargo Pants 1', 'Functional cargo pants with multiple pockets.', 45000, 50, 'image', '2024-01-09 17:27:17.478216+07', '2024-01-09 17:27:17.478216+07');
INSERT INTO public.products VALUES ('59799d6c-bcc5-4e80-b6cb-ef908252abb2', '14b1abc9-2612-4821-a0d0-162e54e6ca3c', 'Trucker Hat 1', 'Classic trucker hat for a casual and cool style.', 17500, 50, 'image', '2024-01-09 17:27:49.455498+07', '2024-01-09 17:27:49.455498+07');
INSERT INTO public.products VALUES ('fc1c1e4b-89ff-42f2-81df-32dd480d697b', '59ed8b9f-968f-493d-a775-a6136987d76d', 'Polo Shirt 1', 'Versatile polo shirt suitable for various occasions.', 23000, 50, 'image', '2024-01-09 17:28:25.503164+07', '2024-01-09 17:28:25.503164+07');
INSERT INTO public.products VALUES ('892e1fca-9652-4c7e-9b58-759bd7ff1722', '66157187-9518-4474-8cef-7b045b5e966e', 'Jeans 3', 'Comfortable and versatile jeans for everyday wear.', 35000, 48, 'image', '2024-01-09 16:50:20.915025+07', '2024-01-11 00:08:15.802935+07');
INSERT INTO public.products VALUES ('8d075149-98d4-413e-829b-a950a0c67ed6', '14b1abc9-2612-4821-a0d0-162e54e6ca3c', 'Snapback Hat 1', 'Classic snapback hat for a sporty look.', 15000, 49, 'image', '2024-01-09 16:51:21.662625+07', '2024-01-11 13:50:01.042535+07');
INSERT INTO public.products VALUES ('efd16e40-99df-41a9-8cbc-18e29ed80a38', '14b1abc9-2612-4821-a0d0-162e54e6ca3c', 'Snapback Hat 2', 'Trendy snapback hat with a modern design.', 18000, 49, 'image', '2024-01-09 16:54:08.469357+07', '2024-01-11 13:50:01.045721+07');


--
-- TOC entry 3041 (class 0 OID 27022)
-- Dependencies: 200
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES ('d6cbde8b-9b04-41d4-8cae-49cc1cddaff8', 'test@test.com', '$2a$10$FJlxYampHtC8mHcfuLQiGuzyZ15F1HK7HKIPmkj1Ds1aRRAJ81nRC', 'test@test.com', 'address', '62123123', '2024-01-09 15:16:31.135255+07', '2024-01-09 15:16:31.135255+07');
INSERT INTO public.users VALUES ('2f3feaab-f6c6-4cfc-8ab3-befc62b430c9', 'test2@test.com', '$2a$10$JJ7ND7926CLycysAG7iucudbULLrXcMWC0O75NbXlLPqOTK1jsN5i', 'test2@test.com', 'address2', '62123123', '2024-01-09 15:21:53.236875+07', '2024-01-09 15:21:53.236875+07');
INSERT INTO public.users VALUES ('49cc698d-dd20-4bd0-a913-fb26df0890db', 'gema@gema.com', '$2a$10$HCHa94Grnm.EnYPxNl1SCeLVpQ74W9q.L70fYZEckwBTTEQaiU5IS', 'gema@gema.com', 'surabaya', '62822123123', '2024-01-11 11:02:53.781991+07', '2024-01-11 11:02:53.781991+07');


--
-- TOC entry 2895 (class 2606 OID 27050)
-- Name: cart_products cart_products_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart_products
    ADD CONSTRAINT cart_products_unique UNIQUE (id);


--
-- TOC entry 2893 (class 2606 OID 27038)
-- Name: carts carts_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT carts_unique UNIQUE (id);


--
-- TOC entry 2899 (class 2606 OID 27075)
-- Name: categories categories_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_unique UNIQUE (id);


--
-- TOC entry 2903 (class 2606 OID 27107)
-- Name: order_products order_products_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_products
    ADD CONSTRAINT order_products_unique UNIQUE (id);


--
-- TOC entry 2901 (class 2606 OID 27095)
-- Name: orders orders_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_unique UNIQUE (id);


--
-- TOC entry 2897 (class 2606 OID 27065)
-- Name: products products_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_unique UNIQUE (id);


--
-- TOC entry 2891 (class 2606 OID 27031)
-- Name: users users_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_unique UNIQUE (id);


--
-- TOC entry 2905 (class 2606 OID 27051)
-- Name: cart_products cart_products_carts_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart_products
    ADD CONSTRAINT cart_products_carts_fk FOREIGN KEY (id_cart) REFERENCES public.carts(id);


--
-- TOC entry 2906 (class 2606 OID 27081)
-- Name: cart_products cart_products_products_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart_products
    ADD CONSTRAINT cart_products_products_fk FOREIGN KEY (id_product) REFERENCES public.products(id);


--
-- TOC entry 2904 (class 2606 OID 27039)
-- Name: carts carts_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT carts_users_fk FOREIGN KEY (id_user) REFERENCES public.users(id);


--
-- TOC entry 2909 (class 2606 OID 27108)
-- Name: order_products order_products_orders_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_products
    ADD CONSTRAINT order_products_orders_fk FOREIGN KEY (id_order) REFERENCES public.orders(id);


--
-- TOC entry 2910 (class 2606 OID 27113)
-- Name: order_products order_products_products_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_products
    ADD CONSTRAINT order_products_products_fk FOREIGN KEY (id_product) REFERENCES public.products(id);


--
-- TOC entry 2908 (class 2606 OID 27096)
-- Name: orders orders_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_users_fk FOREIGN KEY (id_user) REFERENCES public.users(id);


--
-- TOC entry 2907 (class 2606 OID 27076)
-- Name: products products_categories_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_categories_fk FOREIGN KEY (id_category) REFERENCES public.categories(id);


-- Completed on 2024-01-11 16:43:09

--
-- PostgreSQL database dump complete
--

