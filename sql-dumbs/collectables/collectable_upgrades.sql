/* Clojure Table */
CREATE TABLE collectable_upgrades (
	collectable_id UUID NOT NULL,
	upgraded_collectable_id UUID NOT NULL,
	length INT NOT NULL DEFAULT 0,  
	CONSTRAINT collectable_upgrades_pk PRIMARY KEY (collectable_id, upgraded_collectable_id),
	CONSTRAINT collectable_fk FOREIGN KEY (collectable_id) REFERENCES collectables(id),
	CONSTRAINT upgraded_collectable_fk FOREIGN KEY (upgraded_collectable_id) REFERENCES collectables(id)
);

/* Query all children of a node */
SELECT upgraded_collectables.* FROM collectable_upgrades 
INNER JOIN collectables upgraded_collectables ON upgraded_collectables.id = collectable_upgrades.upgraded_collectable_id
WHERE collectable_upgrades.collectable_id = (SELECT id FROM collectables WHERE NAME = $collectable_name);

/* Spell Upgrades */
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Vengeful Spirit'), (SELECT id FROM collectables WHERE name = 'Vengeful Spirit'), 0);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Vengeful Spirit'), (SELECT id FROM collectables WHERE name = 'Shade Soul'), 1);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Desolate Dive'), (SELECT id FROM collectables WHERE name = 'Desolate Dive'), 0);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Desolate Dive'), (SELECT id FROM collectables WHERE name = 'Descending Dark'), 1);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Howling Wraiths'), (SELECT id FROM collectables WHERE name = 'Howling Wraiths'), 0);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Howling Wraiths'), (SELECT id FROM collectables WHERE name = 'Abyss Shriek'), 1);

/* Nail Upgrades */
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Old Nail'), (SELECT id FROM collectables WHERE name = 'Old Nail'), 0);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Old Nail'), (SELECT id FROM collectables WHERE name = 'Sharpened Nail'), 1);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Old Nail'), (SELECT id FROM collectables WHERE name = 'Channelled Nail'), 2);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Old Nail'), (SELECT id FROM collectables WHERE name = 'Coiled Nail'), 3);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Old Nail'), (SELECT id FROM collectables WHERE name = 'Pure Nail'), 4);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Sharpened Nail'), (SELECT id FROM collectables WHERE name = 'Sharpened Nail'), 0);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Sharpened Nail'), (SELECT id FROM collectables WHERE name = 'Channelled Nail'), 1);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Sharpened Nail'), (SELECT id FROM collectables WHERE name = 'Coiled Nail'), 2);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Sharpened Nail'), (SELECT id FROM collectables WHERE name = 'Pure Nail'), 3);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Channelled Nail'), (SELECT id FROM collectables WHERE name = 'Channelled Nail'), 0);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Channelled Nail'), (SELECT id FROM collectables WHERE name = 'Coiled Nail'), 1);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Channelled Nail'), (SELECT id FROM collectables WHERE name = 'Pure Nail'), 2);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Coiled Nail'), (SELECT id FROM collectables WHERE name = 'Coiled Nail'), 0);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Coiled Nail'), (SELECT id FROM collectables WHERE name = 'Pure Nail'), 1);
INSERT INTO collectable_upgrades (collectable_id, upgraded_collectable_id, length) VALUES ((SELECT id FROM collectables WHERE name = 'Pure Nail'), (SELECT id FROM collectables WHERE name = 'Pure Nail'), 0);

