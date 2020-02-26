CREATE OR REPLACE FUNCTION next_id(OUT result bigint, seq text) AS $$
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
    $$ LANGUAGE PLPGSQL;


CREATE OR REPLACE FUNCTION convertTVkdau (x text) RETURNS text AS
$$
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
$$ LANGUAGE plpgsql;
--Generate new func insta5

--Create table admin_users
--Drop table admin_users
drop table if exists admin_users;
drop sequence if exists admin_users_id_seq;
CREATE SEQUENCE admin_users_id_seq;

CREATE TABLE admin_users (
    id bigint NOT NULL DEFAULT next_id('admin_users_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    email text,
    first_name text,
    last_name text,
    password text,
    username text,
    role text,
    last_login timestamp,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

--Create table movies
--Drop table movies
drop table if exists movies;
drop sequence if exists movies_id_seq;
CREATE SEQUENCE movies_id_seq;

CREATE TABLE movies (
    id bigint NOT NULL DEFAULT next_id('movies_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    name text,
    description text,
    image text,
    trailer text,
    started_at timestamp,
    duration integer,
    rating float,
    type integer
);

--Create table theaters
--Drop table theaters
drop table if exists theaters;
drop sequence if exists theaters_id_seq;
CREATE SEQUENCE theaters_id_seq;

CREATE TABLE theaters (
    id bigint NOT NULL DEFAULT next_id('theaters_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    name text,
    description text,
    state text,
    city text,
    district text,
    ward text,
    street text
);

--Create table movie_in_theaters
--Drop table movie_in_theaters
drop table if exists movie_session_in_theaters;
drop sequence if exists movie_session_in_theaters_id_seq;
CREATE SEQUENCE movie_session_in_theaters_id_seq;

CREATE TABLE movie_session_in_theaters (
    id bigint NOT NULL DEFAULT next_id('movie_session_in_theaters_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    start_time timestamp,
    end_time timestamp,
    movie_id bigint,
    theater_id bigint,
    foreign key (movie_id) references movies(id),
    foreign key (theater_id) references theaters(id)
);

--Create table consumers
--Drop table consumers
drop table if exists consumers;
drop sequence if exists consumers_id_seq;
CREATE SEQUENCE consumers_id_seq;

CREATE TABLE consumers (
    id bigint NOT NULL DEFAULT next_id('consumers_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    email text,
    name text,
    phone text,
    address text
);

--Create table tickets
--Drop table tickets
drop table if exists tickets;
drop sequence if exists tickets_id_seq;
CREATE SEQUENCE tickets_id_seq;

CREATE TABLE tickets (
    id bigint NOT NULL DEFAULT next_id('tickets_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    value integer,
    number_seat text,
    number_theater integer,
    movie_session_in_theater_id bigint,
    type integer,
    customer_id bigint,
    foreign key (customer_id) references customers(id)
    foreign key (movie_session_in_theater_id) references movie_session_in_theaters(id)
);

--Create table payments
--Drop table payments
drop table if exists payments;
drop sequence if exists payments_id_seq;
CREATE SEQUENCE payments_id_seq;

CREATE TABLE payments (
    id bigint NOT NULL DEFAULT next_id('payments_id_seq') primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    amount integer,
    ticket_id bigint,
    foreign key (ticket_id) references tickets(id)
);