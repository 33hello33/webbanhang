CREATE TABLE "accounts" (
  "user_name" varchar NOT NULL PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "hash_password" varchar NOT NULL,
  "email" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "unit" varchar NOT NULL DEFAULT 'cái',
  "price_import" bigint NOT NULL DEFAULT 0,
  "amount" bigint NOT NULL DEFAULT 0,
  "price" bigint NOT NULL DEFAULT 0,
  "warehouse" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "id_supplier" bigserial
);

CREATE TABLE "suppliers" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "address" varchar,
  "notes" varchar
);

CREATE TABLE "customers" (
  "id" bigserial,
  "phone" varchar NOT NULL PRIMARY KEY,
  "name" varchar NOT NULL,
  "address" varchar
);

CREATE TABLE "invoices" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "customers_phone"  varchar NOT NULL,
  "total_money" bigint NOT NULL,
  "had_paid" bigint NOT NULL,
  "is_deleted" boolean DEFAULT false NOT NULL
);

CREATE TABLE "invoice_detail" (
  "id" bigserial PRIMARY KEY,
  "invoice_id" bigserial NOT NULL,
  "product_id" bigserial NOT NULL,
  "price_at_sell" bigint NOT NULL,
  "amount" float NOT NULL,
  "total_price" bigint NOT NULL,
  "discount" float DEFAULT 0 NOT NULL,
  "last_price" bigint NOT NULL
);

CREATE INDEX ON "accounts" ("user_name");

CREATE INDEX ON "products" ("name");

CREATE INDEX ON "suppliers" ("name");

CREATE INDEX ON "customers" ("phone");

ALTER TABLE "products" ADD FOREIGN KEY ("id_supplier") REFERENCES "suppliers" ("id");

ALTER TABLE "invoices" ADD FOREIGN KEY ("customers_phone") REFERENCES "customers" ("phone");

ALTER TABLE "invoice_detail" ADD FOREIGN KEY ("invoice_id") REFERENCES "invoices" ("id");

ALTER TABLE "invoice_detail" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
