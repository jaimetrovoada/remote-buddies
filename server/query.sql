-- name: ListNearbyUsers :many
SELECT name, st_astext(coords) as coords
FROM "User"
WHERE ST_DWithin(coords, ST_MakePoint($1, $2)::geography, $3);

-- name: UpdateUserLocation :exec
UPDATE "User"
SET coords = ST_Point($1, $2, 4326)
WHERE email = $3;

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
