-- +migrate Up
CREATE TABLE "city"
(
    "id"           bigserial,
    "name"         varchar(128),
    "country_code" varchar(3),
    "province"     varchar(128),
    PRIMARY KEY ("id")
);

CREATE TABLE "area"
(
    "id"           bigserial,
    "name"         varchar(128),
    "city_id"      integer,
    "country_code" varchar(3),
    PRIMARY KEY ("id"),
    CONSTRAINT "FK_area.city_id"
        FOREIGN KEY ("city_id")
            REFERENCES "city" ("id")
);

CREATE TABLE "organisation"
(
    "id"         bigserial,
    "name"       varchar(128),
    "full_name"  varchar(256),
    "created_at" timestamp,
    PRIMARY KEY ("id")
);

CREATE TABLE "user"
(
    "id"              bigserial,
    "first_name"      varchar(128),
    "last_name"       varchar(128),
    "last_address"    integer,
    "email"           varchar(128),
    "phone_number"    varchar(128),
    "password_hash"   varchar(512),
    "role"            varchar(64),
    "birthdate"       date,
    "status"          varchar(16),
    "created_at"      timestamp,
    "organisation_id" integer,
    PRIMARY KEY ("id"),
    CONSTRAINT "FK_user.organisation_id"
        FOREIGN KEY ("organisation_id")
            REFERENCES "organisation" ("id")
);

CREATE TABLE "address"
(
    "id"            bigserial,
    "user_id"       integer,
    "geo"           text,
    "street_name"   varchar(128),
    "number"        integer,
    "building_name" varchar(128),
    "zone_name"     varchar(128),
    "postal_index"  varchar(32),
    "addr_line1"    text,
    "addr_line2"    text,
    "postal_code"   varchar(16),
    "area_id"       integer,
    PRIMARY KEY ("id"),
    CONSTRAINT "FK_address.area_id"
        FOREIGN KEY ("area_id")
            REFERENCES "area" ("id"),
    CONSTRAINT "FK_address.user_id"
        FOREIGN KEY ("user_id")
            REFERENCES "user" ("id")
);

CREATE TABLE "laundry"
(
    "id"              bigserial,
    "name"            varchar(256),
    "description"     text,
    "address"         varchar(512),
    "geolocation"     varchar(512),
    "date_added"      date,
    "area_id"         integer,
    "organisation_id" integer,
    PRIMARY KEY ("id"),
    CONSTRAINT "FK_laundry.organisation_id"
        FOREIGN KEY ("organisation_id")
            REFERENCES "organisation" ("id"),
    CONSTRAINT "FK_laundry.area_id"
        FOREIGN KEY ("area_id")
            REFERENCES "area" ("id")
);

CREATE TABLE "order"
(
    "id"           bigserial,
    "public_id"    varchar(64),
    "name"         varchar(256),
    "created_at"   timestamp,
    "due_date"     date,
    "fulfilled_at" timestamp,
    "address_from" integer,
    "address_to"   integer,
    "status"       varchar(64),
    "customer_id"  integer,
    "deliverer_id" integer,
    "laundry_id"   integer,
    PRIMARY KEY ("id"),
    CONSTRAINT "FK_order.deliverer_id"
        FOREIGN KEY ("deliverer_id")
            REFERENCES "user" ("id"),
    CONSTRAINT "FK_order.customer_id"
        FOREIGN KEY ("customer_id")
            REFERENCES "user" ("id"),
    CONSTRAINT "FK_order.laundry_id"
        FOREIGN KEY ("laundry_id")
            REFERENCES "laundry" ("id")
);

CREATE TABLE "service_item"
(
    "id"                  bigserial,
    "service_category_id" integer,
    "name"                varchar(256),
    "self_reference"      integer,
    "laundry_id"          integer,
    "price"               numeric,
    "currency"            varchar(16),
    PRIMARY KEY ("id"),
    CONSTRAINT "FK_service_item.laundry_id"
        FOREIGN KEY ("laundry_id")
            REFERENCES "laundry" ("id")
);

CREATE TABLE "service_category"
(
    "id"              serial,
    "name"            varchar(100),
    "service_item_id" integer,
    PRIMARY KEY ("id"),
    CONSTRAINT "FK_service_category.service_item_id"
        FOREIGN KEY ("service_item_id")
            REFERENCES "service_item" ("id")
);

CREATE TABLE "order_item"
(
    "id"       bigserial,
    "name"     varchar(256),
    "price"    numeric,
    "currency" varchar(16),
    "order_id" integer,
    PRIMARY KEY ("id"),
    CONSTRAINT "FK_order_item.order_id"
        FOREIGN KEY ("order_id")
            REFERENCES "order" ("id")
);

CREATE TABLE "refresh_token"
(
    "id"         serial,
    "sign"       varchar(128),
    "device_id"  varchar(64),
    "issued_at"  timestamp,
    "expires_at" timestamp,
    "user_id"    integer,
    PRIMARY KEY ("id"),
    CONSTRAINT "FK_refresh_token.user_id"
        FOREIGN KEY ("user_id")
            REFERENCES "user" ("id")
);


-- +migrate Down
DROP TABLE IF EXISTS "address" cascade;
DROP TABLE IF EXISTS "order_item" cascade;
DROP TABLE IF EXISTS "service_category" cascade;
DROP TABLE IF EXISTS "area" cascade;
DROP TABLE IF EXISTS "organisation" cascade;
DROP TABLE IF EXISTS "user" cascade;
DROP TABLE IF EXISTS "service_item" cascade;
DROP TABLE IF EXISTS "order" cascade;
DROP TABLE IF EXISTS "city" cascade;
DROP TABLE IF EXISTS "laundry" cascade;
DROP TABLE IF EXISTS "refresh_token" cascade;
