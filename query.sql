-- name: GetCat :one
SELECT * FROM cats
WHERE id = $1
LIMIT 1;

-- name: GetCats :many
SELECT * FROM cats;

-- name: CreateCat :exec
INSERT INTO cats (
    id, name, experience, breed, salary
) VALUES (
             $1, $2, $3, $4, $5
         )
    ON CONFLICT (id) DO UPDATE
    SET name = $2,
        experience = $3,
        breed = $4,
        salary = $5;

-- name: UpdateCat :exec
UPDATE cats
set name = $2,
    experience = $3,
    breed = $4,
    salary = $5
WHERE id = $1;

-- name: DeleteCat :exec
DELETE FROM cats
WHERE id = $1;

-- name: GetMission :one
SELECT
    missions.id, missions.targets, missions.assignee_id, missions.complete,
    cats.name AS assignee_name, cats.experience AS assignee_experience, cats.breed AS assignee_breed,
    cats.salary AS assignee_salary
FROM missions
LEFT JOIN cats
ON missions.assignee_id  = cats.id
WHERE missions.id = $1
LIMIT 1;

-- name: GetMissions :many
SELECT
    missions.id, missions.targets, missions.assignee_id, missions.complete,
    cats.name AS assignee_name, cats.experience AS assignee_experience, cats.breed AS assignee_breed,
    cats.salary AS assignee_salary
FROM missions
LEFT JOIN cats ON missions.assignee_id  = cats.id;

-- name: CreateMission :exec
INSERT INTO missions (
    id, targets, assignee_id, complete
) VALUES (
             $1, $2, $3, $4
         )
    ON CONFLICT (id) DO UPDATE
                            SET targets = $2,
                            assignee_id = $3,
                            complete = $4;


-- name: UpdateMission :exec
UPDATE missions
set targets = $2,
    assignee_id = $3,
    complete = $4
WHERE id = $1;

-- name: DeleteMission :exec
DELETE FROM missions
WHERE id = $1;