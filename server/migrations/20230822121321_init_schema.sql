-- +goose Up
-- +goose StatementBegin
-- CreateTable
CREATE TABLE "User" (
    "id" UUID NOT NULL DEFAULT gen_random_uuid(),
    "name" TEXT,
    "email" TEXT,
    "emailVerified" TIMESTAMPTZ(3),
    "image" TEXT,
    "created_at" TIMESTAMPTZ(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ(3) NOT NULL,

    CONSTRAINT "User_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Account" (
    "id" UUID NOT NULL DEFAULT gen_random_uuid(),
    "userId" UUID NOT NULL,
    "type" TEXT NOT NULL,
    "provider" TEXT NOT NULL,
    "providerAccountId" TEXT NOT NULL,
    "refresh_token" TEXT,
    "access_token" TEXT,
    "expires_at" INTEGER,
    "token_type" TEXT,
    "scope" TEXT,
    "id_token" TEXT,
    "session_state" TEXT,
    "oauth_token_secret" TEXT,
    "oauth_token" TEXT,

    CONSTRAINT "Account_pkey" PRIMARY KEY ("id")
);

-- AddReference
ALTER TABLE "Account" ADD CONSTRAINT "Account_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- CreateIndex
CREATE UNIQUE INDEX "Account_provider_providerAccountId_key" ON "Account"("provider", "providerAccountId");

-- CreateIndex
CREATE UNIQUE INDEX "User_email_key" ON "User"("email");

-- CreateExtension
CREATE EXTENSION IF NOT EXISTS "postgis";

-- AlterTable
ALTER TABLE "User" ADD COLUMN "coords" geometry(Point, 4326);

-- CreateIndex
CREATE INDEX "coords_idx" ON "User" USING GIST ("coords");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "Account";
DROP TABLE "User";
-- +goose StatementEnd