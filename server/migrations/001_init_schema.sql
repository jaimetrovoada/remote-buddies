-- Write your migrate up statements here

-- createTable
CREATE TABLE "User" (
    "id" UUID NOT NULL DEFAULT gen_random_uuid(),
    "name" TEXT,
    "email" TEXT,
    "emailVerified" TIMESTAMPTZ(3),
    "image" TEXT,
    "coords" geometry(Point, 4326),
    "interests" varchar ARRAY[5],
    "created_at" TIMESTAMPTZ(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ(3) NOT NULL,
    "oauth_type" TEXT,
    "oauth_provider" TEXT,
    "oauth_providerAccountId" TEXT NOT NULL,
    "oauth_refresh_token" TEXT,
    "oauth_access_token" TEXT,
    "oauth_expires_at" INTEGER,
    "oauth_token_type" TEXT,
    "oauth_scope" TEXT,
    "oauth_id_token" TEXT,
    "oauth_session_state" TEXT,
    "oauth_token_secret" TEXT,
    "oauth_token" TEXT,
    
    CONSTRAINT "User_pkey" PRIMARY KEY ("id")
);

-- CreateUserIdIndex
CREATE UNIQUE INDEX "UserId_key" ON "User"("id");

-- CreateEmailIndex
CREATE UNIQUE INDEX "User_email_key" ON "User"("email");

-- CreateCoordsIndex
CREATE INDEX "Coords_idx" ON "User" USING GIST ("coords");

---- create above / drop below ----

DROP TABLE "User";

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
