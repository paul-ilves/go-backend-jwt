-- +migrate Up
alter table "city" alter column "name" set not null;
alter table "city" alter column "country_code" set not null;

alter table "area" alter column "name" set not null;
alter table "area" alter column "country_code" set not null;

alter table "organisation" alter column "name" set not null;

alter table "user" alter column "first_name" set not null;
alter table "user" alter column "email" set not null;
alter table "user" alter column "phone_number" set not null;
alter table "user" alter column "password_hash" set not null;
alter table "user" add constraint "email_must_be_unique" unique ("email");
alter table "user" add constraint "phone_number_must_be_unique" unique ("phone_number");

alter table "address" alter column "user_id" set not null;
alter table "address" alter column "street_name" set not null;
alter table "address" alter column "number" set not null;
alter table "address" alter column "area_id" set not null;

alter table "laundry" alter column "name" set not null;
alter table "laundry" alter column "area_id" set not null;
alter table "laundry" alter column "address" set not null;

alter table "order" alter column "status" set not null;
alter table "order" alter column "customer_id" set not null;
alter table "order" alter column "laundry_id" set not null;

alter table "order_item" alter column "name" set not null;

alter table "service_item" alter column "laundry_id" set not null;
alter table "service_item" alter column "price" set not null;
alter table "service_item" alter column "currency" set not null;

-- +migrate Down
alter table "city" alter column "name" drop not null;
alter table "city" alter column "country_code" drop not null;

alter table "area" alter column "name" drop not null;
alter table "area" alter column "country_code" drop not null;

alter table "organisation" alter column "name" drop not null;

alter table "user" alter column "first_name" drop not null;
alter table "user" alter column "email" drop not null;
alter table "user" alter column "phone_number" drop not null;
alter table "user" alter column "password_hash" drop not null;
alter table "user" drop constraint "email_must_be_unique";
alter table "user" drop constraint "phone_number_must_be_unique";

alter table "address" alter column "user_id" drop not null;
alter table "address" alter column "street_name" drop not null;
alter table "address" alter column "number" drop not null;
alter table "address" alter column "area_id" drop not null;

alter table "laundry" alter column "name" drop not null;
alter table "laundry" alter column "area_id" drop not null;
alter table "laundry" alter column "address" drop not null;

alter table "order" alter column "status" drop not null;
alter table "order" alter column "customer_id" drop not null;
alter table "order" alter column "laundry_id" drop not null;

alter table "order_item" alter column "name" drop not null;

alter table "service_item" alter column "laundry_id" drop not null;
alter table "service_item" alter column "price" drop not null;
alter table "service_item" alter column "currency" drop not null;
