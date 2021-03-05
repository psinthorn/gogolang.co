CREATE TABLE `users` (
  `id` bigserial PRIMARY KEY,
  `last_name` varchar(255),
  `email` varchar(255),
  `avatar` varchar(255),
  `password` varchar(255),
  `status` varchar(255),
  `date_created` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `contents` (
  `id` bigserial PRIMARY KEY,
  `title` varchar(255),
  `sub_title` varchar(255),
  `content` varchar(255),
  `content_type` varchar(255),
  `category` int,
  `image` varchar(255),
  `tags` varchar(255),
  `author_id` varchar(255) NOT NULL,
  `status` varchar(255),
  `date_created` timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE `categories` (
  `id` bigserial PRIMARY KEY,
  `title` varchar(255),
  `describtion` varchar(255),
  `image` varchar(255),
  `author_id` varchar(255) NOT NULL,
  `status` varchar(255),
  `date_created` timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE `contents` ADD FOREIGN KEY (`category`) REFERENCES `categories` (`id`);

ALTER TABLE `contents` ADD FOREIGN KEY (`author_id`) REFERENCES `users` (`id`);

ALTER TABLE `categories` ADD FOREIGN KEY (`author_id`) REFERENCES `users` (`id`);
