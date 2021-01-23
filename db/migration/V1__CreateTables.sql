create table if not exists program
(
    program_id   serial,
    program_name varchar(50),
    description  varchar(4000)
);

create table if not exists set
(
    set_id      serial,
    workout_id  int not null,
    exercise    varchar(50) not null,
    start_date  timestamp null,
    end_date    timestamp not null,
    weight      int null,
    weight_unit varchar(10) null,
    reps        int default 1 not null,
    hold_time   int null,
    set_order   int
);

create table if not exists workout
(
    workout_id serial,
    program_id int,
    workout_date    timestamp default now(),
    cycle      int,
    cycle_name varchar(20),
    phase      int,
    phase_name varchar(20),
    comment    varchar(1000)
);

create table if not exists workout_sets
(
    workout_id int,
    set_id     int
);

create table if not exists progression
(
    program_id                  int,
    exercise                    varchar(50),
    rep_increase                int,
    rep_increase_start          int,
    rep_increase_max            int,
    weight_increase             int,
    weight_increase_percentage  int,
    phase_count_before_increase int
);

-- Create user if not exists
DO
$do$
    BEGIN
        IF NOT EXISTS(
                SELECT
                FROM pg_catalog.pg_roles
                WHERE rolname = 'webserviceuser') THEN
            CREATE ROLE webserviceuser LOGIN PASSWORD 'webserviceuserpwd';
        END IF;
    END
$do$;

-- Permissions

grant select, update, insert, delete on program to webserviceuser;
grant select, update, insert, delete on workout to webserviceuser;
grant select, update, insert, delete on set to webserviceuser;
grant select, update, insert, delete on workout_sets to webserviceuser;
grant select, update, insert, delete on progression to webserviceuser;

grant select, update on program_program_id_seq to webserviceuser;
grant select, update on set_set_id_seq to webserviceuser;
grant select, update on workout_workout_id_seq to webserviceuser;

