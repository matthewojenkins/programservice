delete
from program;
delete
from workout;
delete
from progression;
delete
from workout_sets;
delete
from set;

DO
$$
    declare
        v_program_id int;
        v_workout_id int;
    begin
    insert into program(program_name, description) VALUES ('MJ Bodybuilding','Home bodybuilding Push/Pull/Legs') returning program_id into v_program_id;

    insert into workout(program_id, cycle, cycle_name, phase, phase_name)
    select v_program_id, 1, 'Starting Round', null, null
    returning workout_id into v_workout_id;

    insert into set(workout_id, exercise, weight, weight_unit, reps, hold_time, set_order, start_date, end_date)
    VALUES (v_workout_id, 'Back Squat', 20, 'KG', 10, 0, 1, '10-12-2020 08:00:00', '10-12-2020 08:01:00');
    insert into set(workout_id, exercise, weight, weight_unit, reps, hold_time, set_order, start_date, end_date)
    VALUES (v_workout_id, 'Back Squat', 50, 'KG', 10, 0, 2, '10-12-2020 08:00:00', '10-12-2020 08:01:00');



    commit;
    end;
$$ language plpgsql;


select * from workout;