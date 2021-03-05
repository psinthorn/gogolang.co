CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "last_name" varchar,
  "email" varchar,
  "avatar" varchar,
  "password" varchar,
  "status" varchar,
  "date_created" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "contents" (
  "id" bigserial PRIMARY KEY,
  "title" varchar,
  "sub_title" varchar,
  "content" varchar,
  "content_type" varchar,
  "category" int,
  "image" varchar,
  "tags" varchar,
  "author_id" varchar NOT NULL,
  "status" varchar,
  "date_created" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "title" varchar,
  "describtion" varchar,
  "image" varchar,
  "author_id" varchar NOT NULL,
  "status" varchar,
  "date_created" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "contents" ADD FOREIGN KEY ("category") REFERENCES "categories" ("id");

ALTER TABLE "contents" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id");

ALTER TABLE "categories" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id");
