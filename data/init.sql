create table if not exists `id_generator` (
  `id` bigint(20)  not null auto_increment,
  `name` varchar(64) not null,
  `current` bigint(20) unsigned not null,
  `modified` timestamp not null,
  primary key (`id`),
  unique key `name` (`name`)
) engine=innodb default charset=utf8mb4;

INSERT INTO short_url.id_generator (id, name, current, modified) VALUES (1, 'SURL', 1, now());
