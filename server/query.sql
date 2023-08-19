-- name: ListLocations :many
SELECT id, st_astext(coords) as coords
FROM "Location"
WHERE ST_DWithin(coords, ST_MakePoint($1, $2)::geography, 100);

-- name: CreateLocation :one
INSERT INTO "Location" (coords)
VALUES (ST_Point($1, $2, 4326))
RETURNING *;

-- USER QUERIES
-- name: GetUser :one
SELECT * FROM "User" WHERE email = $1;

-- name: CreateUser :one
INSERT INTO "User" ("name", "email", "image", "updated_at")
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: CreateAccount :one
INSERT INTO "Account" ("userId", "type", "provider", "providerAccountId", "refresh_token", "access_token", "token_type", "scope")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;
