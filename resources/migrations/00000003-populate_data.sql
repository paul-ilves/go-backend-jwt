-- +migrate Up

insert into "city" (name, country_code)
values ('Kiev', 'ua'),
       ('Lisbon', 'pt'),
       ('Berlin', 'de'),
       ('New York', 'us'),
       ('Moscow', 'ru'),
       ('Saint Petersburg', 'ru');

insert into "area" (name, city_id, country_code)
values ('Kiev Right Bank', (select city.id from "city" where city.name = 'Kiev' and city.country_code = 'ua'), 'ua'),
       ('Kiev Left Bank', (select city.id from "city" where city.name = 'Kiev' and city.country_code = 'ua'), 'ua'),
       ('Trukhanov island', (select city.id from "city" where city.name = 'Kiev' and city.country_code = 'ua'), 'ua'),
       ('West Berlin', (select city.id from "city" where city.name = 'Berlin' and city.country_code = 'ua'), 'ua'),
       ('East Berlin', (select city.id from "city" where city.name = 'Berlin' and city.country_code = 'ua'), 'ua');

insert into "organisation" (name, full_name, created_at)
values ('Cleantastic', 'Cleantastic OOO', '2016-06-22 19:10:25-07'),
       ('Clean-o-Magic', 'Clean-o-Magic', '2015-08-11 19:10:21-09'),
       ('Un Momento', 'Un Momento Inc.', '2001-05-11 19:10:11-09'),
       ('Get It Done!', 'Get It Done! Inc.', '2001-02-12 12:10:11-09'),
       ('Cleanerizers', 'Cleanerizers Inc.', '2020-12-12 22:00:11-09');

insert into laundry (name, description, address, geolocation, date_added, area_id, organisation_id)
VALUES ('Cleantastic 1', 'Cleantastic 1 descr.', 'Cleantastic 1 address', 'lat=12.977540&long=77.599510', '2016-06-22',
        (select area.id from area where area.name = 'Kiev Right Bank'),
        (select organisation.id from organisation where organisation.name = 'Cleantastic')),
       ('Cleantastic 2', 'Cleantastic 2 descr.', 'Cleantastic 2 address', 'lat=12.977540&long=77.599510', '2016-06-22',
        (select area.id from area where area.name = 'Kiev Right Bank'),
        (select organisation.id from organisation where organisation.name = 'Cleantastic')),
       ('Cleantastic 3', 'Cleantastic 3 descr.', 'Cleantastic 3 address', 'lat=12.977540&long=77.599510', '2016-06-22',
        (select area.id from area where area.name = 'Kiev Right Bank'),
        (select organisation.id from organisation where organisation.name = 'Un Momento')),
       ('Un Momento 1', 'Un Momento 1 descr.', 'Un Momento 1 address', 'lat=12.977540&long=77.599510', '2016-06-22',
        (select area.id from area where area.name = 'Kiev Left Bank'),
        (select organisation.id from organisation where organisation.name = 'Un Momento')),
       ('Un Momento 2', 'Un Momento 2 descr.', 'Un Momento 2 address', 'lat=12.977540&long=77.599510', '2016-06-22',
        (select area.id from area where area.name = 'Kiev Left Bank'),
        (select organisation.id from organisation where organisation.name = 'Un Momento')),
       ('Un Momento 3', 'Un Momento 3 descr.', 'Un Momento 3 address', 'lat=12.977540&long=77.599510', '2016-06-22',
        (select area.id from area where area.name = 'Kiev Left Bank'),
        (select organisation.id from organisation where organisation.name = 'Un Momento')),
       ('Get It Done! 1', 'Get It Done! 1 descr.', 'Get It Done! 1 address', 'lat=12.977540&long=77.599510',
        '2016-06-22',
        (select area.id from area where area.name = 'Kiev Left Bank'),
        (select organisation.id from organisation where organisation.name = 'Get It Done!')),
       ('Cleanerizers 1', 'Cleanerizers 1 descr.', 'Cleanerizers 1 address', 'lat=12.977540&long=77.599510',
        '2016-06-22',
        (select area.id from area where area.name = 'Kiev Right Bank'),
        (select organisation.id from organisation where organisation.name = 'Cleanerizers'))
;

insert into "user" (first_name, last_name, last_address, email, phone_number, password_hash, role, created_at,
                    organisation_id)
VALUES ('Pavel', 'Streletskiy', null, 'strelok@stalker.com', '555-12-57',
        '$2a$04$DUfwjpJ.gHc4gf3P9fzf/Ok/hH5W.11ASVt1mAolrLjCI6aDs/0q2', 'customer', '2016-06-22 19:10:25-07',
        null),
       ('Shram', 'The Mercenary', null, 'shram@stalker.com', '555-13-58',
        '$2a$04$DUfwjpJ.gHc4gf3P9fzf/Ok/hH5W.11ASVt1mAolrLjCI6aDs/0q2', 'customer', '2016-06-22 19:10:25-07',
        null),
       ('Alexander', 'Degtyaryov', null, 'degtyaryov@sbu.gov.ua', '555-11-11',
        '$2a$04$DUfwjpJ.gHc4gf3P9fzf/Ok/hH5W.11ASVt1mAolrLjCI6aDs/0q2', 'customer',
        '2016-06-22 19:10:25-07',
        null),
       ('Sidorovich', 'Sidorovich', null, 'sidorovich@merchant.ua', '555-14-57',
        '$2a$04$DUfwjpJ.gHc4gf3P9fzf/Ok/hH5W.11ASVt1mAolrLjCI6aDs/0q2', 'admin', '2016-06-22 19:10:25-07',
        null),
       ('Barman', 'Barman', null, 'barman@merchant.ua', '555-77-77',
        '$2a$04$DUfwjpJ.gHc4gf3P9fzf/Ok/hH5W.11ASVt1mAolrLjCI6aDs/0q2', 'admin', '2016-06-22 19:10:25-07',
        null);

-- +migrate Down

truncate table "address" cascade;
truncate table "order_item" cascade;
truncate table "service_category" cascade;
truncate table "area" cascade;
truncate table "organisation" cascade;
truncate table "user" cascade;
truncate table "service_item" cascade;
truncate table "order" cascade;
truncate table "city" cascade;
truncate table "laundry" cascade;
