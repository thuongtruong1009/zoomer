CREATE TABLE "public"."rooms" (
    "id" text NOT NULL,
    "name" text,
    "description" text,
    "category" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "created_by" text,
    CONSTRAINT "rooms_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE TABLE "public"."users" (
    "id" text NOT NULL,
    "username" text NOT NULL,
    "password" text NOT NULL,
    "limit" bigint NOT NULL,
    CONSTRAINT "users_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "users_username_key" UNIQUE ("username")
) WITH (oids = false);

ALTER TABLE ONLY "public"."rooms" ADD CONSTRAINT "fk_rooms_user" FOREIGN KEY (created_by) REFERENCES users(id) NOT DEFERRABLE;

