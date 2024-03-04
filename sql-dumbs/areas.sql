CREATE TABLE areas (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	"name" VARCHAR(20) NOT NULL,
	"quote" VARCHAR(400) NOT NULL,
	kingdom_id uuid NOT NULL,
	CONSTRAINT areas_pk PRIMARY KEY (id),
	CONSTRAINT kingdom_fk FOREIGN KEY (kingdom_id) REFERENCES kingdoms(id) ON DELETE CASCADE ON UPDATE CASCADE
);

INSERT INTO areas (name, quote, kingdom_id) VALUES ('Dirtmouth', 'Our town''s fallen quiet you see. The other residents, they''ve all disappeared. Headed down that well, one by one, into the caverns below.
– Elderbug', (select id from kingdoms where name = 'Hallownest'));
INSERT INTO areas (name, quote, kingdom_id) VALUES ('Forgotten Crossroads', 'Still winding your way through these beautiful highways? Just imagine how they must have looked during the kingdom''s height, thick with traffic and bustling with life!
– Cornifer', (select id from kingdoms where name = 'Hallownest'));
INSERT INTO areas (name, quote, kingdom_id) VALUES ('Greenpath', 'The greater mind once dreamed of leaf and cast these caverns so.
In every bush and every vine the mind of Unn reveals itself to us.
– The Mosskin', (select id from kingdoms where name = 'Hallownest'));
INSERT INTO areas (name, quote, kingdom_id) VALUES ('Fungal Wastes', '
Needle is the term used to describe the bladed weapon wielded by Hornet. The needle is attached with a thread of silk.My understanding of Hallownest can be a little vague, but below those leafy caverns is a fungal grove, once home to peaceful creatures not quite bug and not quite plant.
– Elderbug', (select id from kingdoms where name = 'Hallownest'));
INSERT INTO areas (name, quote, kingdom_id) VALUES ('Fog Canyon', 'The canyon below us, the one thick with fog and crackling with strange energy... a Hunter can lose their senses down there. Be careful... strange and unnatural things lurk there.
– The Hunter', (select id from kingdoms where name = 'Hallownest'));
INSERT INTO areas (name, quote, kingdom_id) VALUES ('City of Tears', 'The city looks to be built into an enormous cavern, and the rain pours down from cracks in the stone above. There must be a lot of water up there somewhere.
– Quirrel', (select id from kingdoms where name = 'Hallownest'));
INSERT INTO areas (name, quote, kingdom_id) VALUES ('Royal Waterways', 'Aren''t these waterways thrilling? A labyrinth of pipes and tunnels... I couldn''t have asked for a better place to employ my talents. It''s all so orderly, so considered, nothing like the crude irregularity of those caverns...
– Cornifer', (select id from kingdoms where name = 'Hallownest'));
INSERT INTO areas (name, quote, kingdom_id) VALUES ('Crystal Peak', 'There is some strange power hidden in the crystals that grow up there in the peaks. They gleam and glow in the darkness, a bright point of searing heat in each one. They sing too, if you listen. Very softly...
– The Hunter', (select id from kingdoms where name = 'Hallownest'));
INSERT INTO areas (name, quote, kingdom_id) VALUES ('Resting Grounds', 'The Resting Grounds... Passengers would come here to conduct rituals for those who had passed on...Not any more though. Perhaps the dead conduct their own rituals now?
– The Last Stag', (select id from kingdoms where name = 'Hallownest'));
INSERT INTO areas (name, quote, kingdom_id) VALUES ('Howling Cliffs', 'Enjoying the bracing air? We''re quite close to Hallownest''s borders and those desolate plains that surround it...
– Cornifer', (select id from kingdoms where name = 'Hallownest'));
INSERT INTO areas (name, quote, kingdom_id) VALUES ('Deepnest', 'This whole area swarms with deadly critters, biting, burning, scratching types. I''d thought to test my strength against them. Now that bravado has left me and all I feel is tired and sore.
– Cloth', (select id from kingdoms where name = 'Hallownest'));
INSERT INTO areas (name, quote, kingdom_id) VALUES ('Ancient Basin', 'Did you know the caverns continue even below the capital? Few have ventured that deep so the details are scant. Those who made it back told of impossibly old structures and roads formed as though the rock itself possessed a will.
– Elderbug', (select id from kingdoms where name = 'Hallownest'));
INSERT INTO areas (name, quote, kingdom_id) VALUES ('Kingdom''s Edge', 'This ashen place is grave of Wyrm. Once told, it came to die. But what is death for that ancient being? More transformation methinks.
This failed kingdom is product of the being spawned from that event.
–  Bardoon', (select id from kingdoms where name = 'Hallownest'));
INSERT INTO areas (name, quote, kingdom_id) VALUES ('The Hive', 'Though this hive exists within Hallownest, we play no part in its attempt at perpetuation.
– Hive Queen Vespa', (select id from kingdoms where name = 'Hallownest'));
INSERT INTO areas (name, quote, kingdom_id) VALUES ('Queen''s Gardens', 'Have you heard of Hallownest''s Queen? Apparently these gardens were once her retreat... Now some vicious types are crawling all over the place and the plants have grown wild...
– Cornifer', (select id from kingdoms where name = 'Hallownest'));
INSERT INTO areas (name, quote, kingdom_id) VALUES ('White Palace', 'Have you visited the remnants of his palace? It''s down below this city, in the bedrock of the kingdom. Must''ve been an impressive sight in its time. Now there''s nothing left.
It''s a strange thing though. There are no signs of conflict around the area. It''s as though the whole place just vanished.
– Relic Seeker Lemm', (select id from kingdoms where name = 'Hallownest'));
INSERT INTO areas (name, quote, kingdom_id) VALUES ('The Abyss', 'Far down below us, beneath the kingdom, the air grows stiller and a sense of emptiness pervades. Can life flourish down there?
– The Hunter', (select id from kingdoms where name = 'Hallownest'));
INSERT INTO areas (name, quote, kingdom_id) VALUES ('Godhome', 'By what right dost thou trespass here, in this home of the Gods? Shrivel away and begone! Begone!
– Godseeker', (select id from kingdoms where name = 'Hallownest'));