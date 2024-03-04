CREATE table collectable_types (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	"name" VARCHAR(20) NOT NULL,
	CONSTRAINT collectable_types_pk PRIMARY KEY (id)
);

INSERT INTO collectable_types ("name") VALUES ('Item');
INSERT INTO collectable_types ("name") VALUES ('Spell');
INSERT INTO collectable_types ("name") VALUES ('Nail Art');
