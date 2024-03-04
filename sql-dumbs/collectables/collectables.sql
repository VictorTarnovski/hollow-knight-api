CREATE TABLE collectables (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	"name" varchar(20) NOT NULL,
	description varchar(250) NOT NULL,
	collectable_type_id uuid NOT NULL,
	CONSTRAINT collectables_pk PRIMARY KEY (id),
	constraint collectable_type_fk foreign key(collectable_type_id) references collectable_types(id) 
);

/* Items */
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Geo', 'The currency of Hallownest, made from fossilised shells of various forms.
Can be traded for goods or used as toll in various old mechanisms', (select id from collectable_types where name = 'Item'));
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Old Nail', 'A traditional weapon of Hallownest. Its blade is blunt with age and wear', (select id from collectable_types where name = 'Item'));
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Sharpened Nail', 'A traditional weapon of Hallownest, restored to lethal form', (select id from collectable_types where name = 'Item'));
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Channelled Nail', 'A cleft weapon of Hallownest. The blade is exquisitly balanced', (select id from collectable_types where name = 'Item'));
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Coiled Nail', 'A powerful weapon of Hallownest, refined beyond all others', (select id from collectable_types where name = 'Item'));
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Pure Nail', 'The ultimate weapon of Hallownest. Crafted to perfection, this ancient nail reveals its true form', (select id from collectable_types where name = 'Item'));
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Pale Ore', 'Rare, pale metal that emanates an icy chill', (select id from collectable_types where name = 'Item'));
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Mask shard', 'A shard of an ancient mask, worn to protect oneself from harm', (select id from collectable_types where name = 'Item'));


/* Spells */
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Focus', 'Focus collected SOUL to repair your shell and heal damage.
Strike enemies to gather SOUL', (select id from collectable_types where name = 'Spell'));
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Vengeful Spirit', 'Conjure a spirit that will fly forward and burn foes in its path.
The spirit requires SOUL to be conjured. Strike enemies to gather SOUL', (select id from collectable_types where name = 'Spell'));
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Desolate Dive', 'Strike the ground with a concentrated force of SOUL. This force can destroy foes or break through fragile structures.
The force requires SOUL to be conjured. Strike enemies to gather SOUL', (select id from collectable_types where name = 'Spell'));
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Howling Wraiths', 'Blast foes with screaming SOUL.
The Wraiths requires SOUL to be conjured. Strike enemies to gather SOUL', (select id from collectable_types where name = 'Spell'));
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Shade Soul', 'Conjure a shadow that will fly forward and burn foes in its path.
The shadow requires SOUL to be conjured. Strike enemies to gather SOUL', (select id from collectable_types where name = 'Spell'));
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Descending Dark', 'Strike the ground with a concentrated force of SOUL and Shadow. This force can destroy foes or break through fragile structures.
The force requires SOUL to be conjured. Strike enemies to gather SOUL', (select id from collectable_types where name = 'Spell'));
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Abyss Shriek', 'Blast foes with screaming SOUL and Shadow.
The Wraiths requires sic SOUL to be conjured. Strike enemies to gather SOUL', (select id from collectable_types where name = 'Spell'));

/* Nail Arts */
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Cyclone Slash', 'The signature Nail Art of Nailmaster Mato. A spinning attack that rapidly strikes foes on all sides', (select id from collectable_types where name = 'Nail Art'));
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Dash Slash', 'The signature Nail Art of Nailmaster Oro. Strike ahead quickly after dashing forward', (select id from collectable_types where name = 'Nail Art'));
INSERT INTO collectables (name, description, collectable_type_id) VALUES ('Great Slash', 'The signature Nail Art of Nailmaster Sheo. Unleashes a huge slash directly in front of you which deals extra damage to foes', (select id from collectable_types where name = 'Nail Art'));