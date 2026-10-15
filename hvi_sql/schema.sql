CREATE TABLE IF NOT EXISTS role (
	name VARCHAR(255) PRIMARY KEY NOT NULL
);

CREATE TABLE IF NOT EXISTS person (
	id UUID PRIMARY KEY NOT NULL,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	is_email_verified BOOL NOT NULL,
	date_account_creation TIMESTAMP WITH TIME ZONE NOT NULL,
	date_last_connexion TIMESTAMP WITH TIME ZONE NOT NULL,
	role_name VARCHAR(255) NOT NULL,
	FOREIGN KEY (role_name) REFERENCES role (name)
);

CREATE TABLE IF NOT EXISTS category (
	id UUID PRIMARY KEY NOT NULL,
	name VARCHAR(255) NOT NULL,
	date_creation TIMESTAMP WITH TIME ZONE NOT NULL,
	date_last_modification TIMESTAMP WITH TIME ZONE NOT NULL,
	id_creator UUID NOT NULL,
	id_last_editor UUID NOT NULL,
	FOREIGN KEY (id_creator) REFERENCES person (id),
	FOREIGN KEY (id_last_editor) REFERENCES person (id)
);

CREATE TABLE IF NOT EXISTS section (
	id UUID PRIMARY KEY NOT NULL,
	name VARCHAR(255) NOT NULL,
	id_category UUID NOT NULL,
	is_visible_by_external BOOL NOT NULL,
	date_creation TIMESTAMP WITH TIME ZONE NOT NULL,
	date_last_modification TIMESTAMP WITH TIME ZONE NOT NULL,
	id_creator UUID NOT NULL,
	id_last_editor UUID NOT NULL,
	FOREIGN KEY (id_category) REFERENCES category (id),
	FOREIGN KEY (id_creator) REFERENCES person (id),
	FOREIGN KEY (id_last_editor) REFERENCES person (id)
);

CREATE TABLE IF NOT EXISTS subject (
	id UUID PRIMARY KEY NOT NULL,
	title VARCHAR(255) NOT NULL,
	id_section UUID NOT NULL,
	is_pinned BOOL NOT NULL,
	date_creation TIMESTAMP WITH TIME ZONE NOT NULL,
	date_last_modification TIMESTAMP WITH TIME ZONE NOT NULL,
	id_creator UUID NOT NULL,
	id_last_editor UUID NOT NULL,
	FOREIGN KEY (id_section) REFERENCES section (id),
	FOREIGN KEY (id_creator) REFERENCES person (id),
	FOREIGN KEY (id_last_editor) REFERENCES person (id)
);

CREATE TABLE IF NOT EXISTS message (
	id UUID PRIMARY KEY NOT NULL,
	content TEXT NOT NULL,
	id_subject UUID NOT NULL,
	is_first_message BOOL NOT NULL,
	date_creation TIMESTAMP WITH TIME ZONE NOT NULL,
	date_last_modification TIMESTAMP WITH TIME ZONE NOT NULL,
	id_creator UUID NOT NULL,
	id_last_editor UUID NOT NULL,
	FOREIGN KEY (id_subject) REFERENCES subject (id),
	FOREIGN KEY (id_creator) REFERENCES person (id),
	FOREIGN KEY (id_last_editor) REFERENCES person (id)
);


CREATE TABLE IF NOT EXISTS dead_url (
	url VARCHAR(2000) PRIMARY KEY NOT NULL
);
