INSERT INTO role (name) VALUES ('administrator');
INSERT INTO role (name) VALUES ('user');
INSERT INTO role (name) VALUES ('moderator');

INSERT INTO person (
	id,
	name,
	email,
	password,
	is_email_verified,
	date_account_creation,
	date_last_connexion,
	role_name
) VALUES (
	'f02b5c6c-4e0f-4b3f-82ce-84ff4ea23aec',
	'Vincent',
	'vincent.gimbert@yahoo.fr',
	'a',
	false,
	'2026-10-09 13:50:38 CET',
	'2026-10-09 13:50:38 CET',
	'administrator'
);

INSERT INTO category (
	id,
	name,
	date_creation,
	date_last_modification,
	id_creator,
	id_last_editor
) VALUES (
	'5a4761ae-d3e7-47a6-b00d-3f7285f441d6',
	'Témoignages de Victimes et de Proches',
	'2026-10-09 13:54:21 CET',
	'2026-10-09 13:54:21 CET',
	'f02b5c6c-4e0f-4b3f-82ce-84ff4ea23aec',
	'f02b5c6c-4e0f-4b3f-82ce-84ff4ea23aec'
);

INSERT INTO section (
	id,
	name,
	is_visible_by_external,
	id_category,
	date_creation,
	date_last_modification,
	id_creator,
	id_last_editor
) VALUES (
	'09da059e-6e49-4ef2-a88a-def90715a791',
	'Témoignages de victimes visibles par tous',
	true,
	'5a4761ae-d3e7-47a6-b00d-3f7285f441d6',
	'2026-10-09 13:56:54 CET',
	'2026-10-09 13:56:54 CET',
	'f02b5c6c-4e0f-4b3f-82ce-84ff4ea23aec',
	'f02b5c6c-4e0f-4b3f-82ce-84ff4ea23aec'
);

INSERT INTO section (
	id,
	name,
	is_visible_by_external,
	id_category,
	date_creation,
	date_last_modification,
	id_creator,
	id_last_editor
) VALUES (
	'497dcba3-ecbf-4587-a2dd-5eb0665e6880',
	'Témoignages de victimes visibles par les membres uniquement',
	false,
	'5a4761ae-d3e7-47a6-b00d-3f7285f441d6',
	'2026-10-11 09:01:05 CET',
	'2026-10-11 09:01:05 CET',
	'f02b5c6c-4e0f-4b3f-82ce-84ff4ea23aec',
	'f02b5c6c-4e0f-4b3f-82ce-84ff4ea23aec'
);

INSERT INTO subject (
	id,
	title,
	is_pinned,
	id_section,
	date_creation,
	date_last_modification,
	id_creator,
	id_last_editor
) VALUES (
	'02a784ce-8a99-4b4f-b151-bec3b4830f2b',
	'Abus à 9 ans',
	false,
	'09da059e-6e49-4ef2-a88a-def90715a791',
	'2026-10-11 10:11:19 CET',
	'2026-10-11 10:11:19 CET',
	'f02b5c6c-4e0f-4b3f-82ce-84ff4ea23aec',
	'f02b5c6c-4e0f-4b3f-82ce-84ff4ea23aec'
);

INSERT INTO message (
	id,
	content,
	id_subject,
	is_first_message,
	date_creation,
	date_last_modification,
	id_creator,
	id_last_editor
) VALUES (
	'93605524-b798-4a57-8f82-218650b1c18a',
	'J''ai été abusé à 9 ans par l''amie d''une de mes amies à un anniversaire.',
	'02a784ce-8a99-4b4f-b151-bec3b4830f2b',
	true,
	'2026-10-11 10:11:19 CET',
	'2026-10-11 10:11:19 CET',
	'f02b5c6c-4e0f-4b3f-82ce-84ff4ea23aec',
	'f02b5c6c-4e0f-4b3f-82ce-84ff4ea23aec'
);
