CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "pass" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "pass_changed" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00z',
  "created_at" timestamptz NOT NULL DEFAULT NOW()
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");