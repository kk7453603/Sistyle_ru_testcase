BEGIN;

CREATE TABLE projects (
	id UUID PRIMARY KEY,
	name VARCHAR(255)
);

CREATE TABLE programs (
	id UUID PRIMARY KEY,
	name VARCHAR(255),
	name_en VARCHAR(255),
	is_public BOOLEAN,
	project_id UUID REFERENCES projects(id)
);

CREATE TABLE session_subjects (
    id SERIAL PRIMARY KEY,
    color VARCHAR(50),
    program_id UUID REFERENCES programs(id)
);

CREATE TABLE session_subject_text_infos ( 
    id SERIAL PRIMARY KEY,
    subject_id INT REFERENCES session_subjects(id),
    lang INT,
    name VARCHAR(255),
    name_plural VARCHAR(255)
);

CREATE TABLE session_types (
    id SERIAL PRIMARY KEY,
    color VARCHAR(50),
    program_id UUID REFERENCES programs(id)
);

CREATE TABLE session_type_text_infos (
    id SERIAL PRIMARY KEY,
    type_id INT REFERENCES session_types(id),
    lang INT,
    name VARCHAR(255),
    name_plural VARCHAR(255)
);

CREATE TABLE session_zones (
    id SERIAL PRIMARY KEY,
    program_id UUID REFERENCES programs(id)
);

CREATE TABLE session_zone_text_infos (
    id SERIAL PRIMARY KEY,
    zone_id INT REFERENCES session_zones(id),
    lang INT,
    name VARCHAR(255),
    description TEXT
);

CREATE TABLE session_speaker_groups (
    id SERIAL PRIMARY KEY,
    color VARCHAR(50),
    program_id UUID REFERENCES programs(id)
);

CREATE TABLE session_speaker_group_text_infos (
    id SERIAL PRIMARY KEY,
    group_id INT REFERENCES session_speaker_groups(id),
    lang INT,
    name VARCHAR(255),
    name_plural VARCHAR(255)
);

CREATE TABLE speakers (
    id SERIAL PRIMARY KEY,
    photo VARCHAR(255),
    not_show_in_speakers BOOLEAN,
    is_important BOOLEAN,
    important_sort INT,
    is_online BOOLEAN,
    additional_info TEXT,
    program_id UUID REFERENCES programs(id)
);

CREATE TABLE speaker_text_infos (
    id SERIAL PRIMARY KEY,
    speaker_id INT REFERENCES speakers(id),
    lang INT,
    first_name VARCHAR(255),
    second_name VARCHAR(255),
    last_name VARCHAR(255),
    prefix VARCHAR(50),
    position VARCHAR(255),
    biography TEXT );

CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    date TIMESTAMP,
    start_time TIME,
    end_time TIME,
    subject_id INT REFERENCES session_subjects(id),
    type_id INT REFERENCES session_types(id),
    zone_id INT REFERENCES session_zones(id),
    broadcast_stream_url VARCHAR(255),
    broadcast_stream_url_en VARCHAR(255), 
    broadcast_site_url VARCHAR(255), 
    broadcast_site_url_en VARCHAR(255), 
    closed_event BOOLEAN, 
    program_id UUID REFERENCES programs(id) 
);

CREATE TABLE session_speaker_attributes (
    id SERIAL PRIMARY KEY,
    speaker_id INT REFERENCES speakers(id),
    session_id INT REFERENCES sessions(id),
    online_mode INT,
    program_id UUID REFERENCES programs(id)
);


CREATE TABLE session_text_infos ( 
    id SERIAL PRIMARY KEY, 
    session_id INT REFERENCES sessions(id), 
    lang INT, 
    title VARCHAR(255), 
    description TEXT, 
    organizer VARCHAR(255), 
    organizer_contact VARCHAR(255), 
    description_file_name VARCHAR(255) 
);

CREATE TABLE session_speakers ( 
    session_id INT REFERENCES sessions(id), 
    speaker_id INT REFERENCES speakers(id),
    speaker_group_id INT REFERENCES session_speaker_groups(id),
    session_attribute_id INT REFERENCES session_speaker_attributes(id),
    PRIMARY KEY (session_id, speaker_id)
);

COMMIT;