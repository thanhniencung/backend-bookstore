-- +migrate Up
CREATE TABLE "orders"
(
    "user_id" TEXT NOT NULL,
    "order_id" TEXT NOT NULL,
    "status" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "card"
(
    "order_id" TEXT NOT NULL,
    "product_id" TEXT NOT NULL,
    "product_name" TEXT NOT NULL,
    "product_image" TEXT NOT NULL,
    "quantity" INTEGER DEFAULT 1 NOT NULL,
    "price" NUMERIC NOT NULL
);

CREATE TABLE "cate"
(
    "cate_id" text NOT NULL PRIMARY KEY,
    "cate_name" text NOT NULL,
    "created_at" timestamp with time zone,
    "deleted_at" time with time zone,
    "updated_at" time with time zone
);

CREATE TABLE "product"
(
    "product_id" text NOT NULL PRIMARY KEY,
    "product_name" text NOT NULL,
    "quantity" integer NOT NULL,
    "sold_items" integer NOT NULL,
    "price" numeric NOT NULL,
    "cate_id" text NOT NULL,
    "product_image" text NOT NULL,
    "deleted_at" timestamp with time zone,
    "created_at" timestamp with time zone default NOW(),
    "updated_at" timestamp with time zone default NOW(),
    "user_id" text NOT NULL
);

CREATE TABLE "users"
(
    "user_id" text NOT NULL,
    "phone" text NOT NULL,
    "password" text NOT NULL,
    "avatar" text NOT NULL,
    "display_name" text NOT NULL,
    "role" text NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (user_id, phone),
    CONSTRAINT users_phone_key UNIQUE (phone),
    CONSTRAINT users_userid_key UNIQUE (user_id)
);