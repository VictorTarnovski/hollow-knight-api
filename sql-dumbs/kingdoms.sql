CREATE TABLE kingdoms (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	"name" VARCHAR(20) NOT NULL,
	CONSTRAINT kingdoms_pk PRIMARY KEY (id)
);

INSERT INTO kingdoms (name) VALUES ('Hallownest');
