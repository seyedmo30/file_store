-- public.store_information definition

-- Drop table

-- DROP TABLE public.store_information;

CREATE TABLE public.store_information (
	"name" varchar(250) NULL,
	id serial4 NOT NULL,
	tag _varchar NULL,
	"type" varchar(10) NULL,
	hash varchar(100) NULL,
	filename varchar(250) NULL,
	CONSTRAINT store_information_hash_key UNIQUE (hash),
	CONSTRAINT store_information_pkey PRIMARY KEY (id)
);