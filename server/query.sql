-- name: ListLocations :many
SELECT id, st_astext(coords) as coords
FROM "Location"
WHERE ST_DWithin(coords, ST_MakePoint($1, $2)::geography, 100);

-- name: CreateLocation :one
INSERT INTO "Location" (coords)
VALUES (ST_Point($1, $2, 4326))
RETURNING *;

