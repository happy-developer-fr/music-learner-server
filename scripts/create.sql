DROP SCHEMA IF EXISTS music_learner CASCADE;
CREATE SCHEMA music_learner;
CREATE TABLE music_learner.song
(
    SONG_ID UUID PRIMARY KEY NOT NULL,
    NOTES jsonb
);

