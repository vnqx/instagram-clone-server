CREATE TABLE "User" (
    "id" SERIAL PRIMARY KEY,
    "email" TEXT NOT NULL UNIQUE,
    "passwordHash" TEXT NOT NULL
);

CREATE TABLE "Post" (
    "id" SERIAL PRIMARY KEY,
    "photos" TEXT[] NOT NULL,
    "description" TEXT NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "user_id" INT NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "User"(id) ON DELETE CASCADE
);