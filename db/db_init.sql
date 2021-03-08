drop database if exists url_shortener;
drop role if exists url_shortener;

-- CREATE USER

create user url_shortener createdb createuser password 'some_password';

-- CREATE URL_SHORTENER DATABASE

create database url_shortener owner url_shortener;
connect url_shortener

DROP TABLE if exists urls;

CREATE TABLE urls
(
  num_id serial PRIMARY KEY,
  url varchar(255) NOT NULL
)

ALTER TABLE urls OWNER TO url_shortener;
