CREATE TABLE "user" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL
);

CREATE TABLE "task" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar,
  "status" int NOT NULL,
  "priority" int NOT NULL,
  "due_date" bigint NOT NULL
);

CREATE TABLE "user_task" (
  "user_id" integer NOT NULL,
  "task_id" integer NOT NULL
);

-- CREATE TABLE "status" (
--   "id" int PRIMARY KEY,
--   "status" varchar
-- );

-- ALTER TABLE "task" ADD FOREIGN KEY ("status") REFERENCES "status" ("id");

-- ALTER TABLE "user_task" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

-- ALTER TABLE "user_task" ADD FOREIGN KEY ("task_id") REFERENCES "task" ("id");
