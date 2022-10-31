CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "user_name" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "expired_at" timestamptz NOT NULL
);

alter table "sessions" add foreign key ("user_name") references "accounts" ("user_name");