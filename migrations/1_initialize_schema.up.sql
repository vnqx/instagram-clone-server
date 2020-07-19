CREATE TABLE "Post" (
    "id" SERIAL PRIMARY KEY,
    "photos" TEXT[] NOT NULL,
    "description" TEXT NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now()
);