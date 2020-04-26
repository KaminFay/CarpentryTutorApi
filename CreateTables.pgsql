CREATE TABLE IF NOT EXISTS users (
    user_id serial PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    username TEXT,
    role_id INT,
    passwd TEXT
);

CREATE TABLE IF NOT EXISTS classes(
    class_id serial PRIMARY KEY,
    class_name TEXT,
    class_description TEXT,
    class_time TIMESTAMP
);

CREATE TABLE IF NOT EXISTS roles(
    role_id serial PRIMARY KEY,
    role_name TEXT,
    role_description TEXT
);



CREATE TABLE IF NOT EXISTS user_role_lookup(
    user_id integer NOT NULL,
    role_id integer NOT NULL,
    PRIMARY KEY (user_id, role_id),
    CONSTRAINT user_role_role_id_fkey FOREIGN KEY (role_id)
        REFERENCES roles(role_id) MATCH SIMPLE
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    CONSTRAINT user_role_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES users(user_id) MATCH SIMPLE
        ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS student_class_lookup (
    user_id integer NOT NULL,
    class_id integer NOT NULL,
    PRIMARY KEY (user_id, class_id),
    CONSTRAINT student_class_student_id_fkey FOREIGN KEY (user_id)
        REFERENCES users(user_id) MATCH SIMPLE
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    CONSTRAINT student_class_class_id_fkey FOREIGN KEY (class_id)
        REFERENCES classes(class_id) MATCH SIMPLE
        ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS teacher_class_lookup (
    user_id integer NOT NULL,
    class_id integer NOT NULL,
    PRIMARY KEY (user_id, class_id),
    CONSTRAINT teacher_class_teacher_id_fkey FOREIGN KEY (user_id)
        REFERENCES users(user_id) MATCH SIMPLE
        ON UPDATE NO ACTION ON DELETE NO ACTION,
    CONSTRAINT teacher_class_class_id_fkey FOREIGN KEY (class_id)
        REFERENCES classes(class_id) MATCH SIMPLE
        ON UPDATE NO ACTION ON DELETE NO ACTION
);