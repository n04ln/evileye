CREATE TABLE `users` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `screenname` VARCHAR(32) NOT NULL,
    -- `password` BLOB NOT NULL
    `password` VARCHAR(32) NOT NULL
);

CREATE TABLE `blocks` (
    `id` integer primary key autoincrement,
    `prevhash` varchar(256) not null,
    `data` varchar(1024), -- todo: tarekomiとうまくリンクさせる
    `create_time` integer, -- unix time
    `hash` varchar(256) -- NOTE: これは確認用
);

CREATE TABLE `tarekomi` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `status` INTEGER NOT NULL, -- NOTE: 0 .. PENDING, 1 .. APPROVED, 2 .. REJECTED
    `threshold` INTEGER NOT NULL,
    `targetuserid` INTEGER NOT NULL,
    `url` VARCHAR(256) NOT NULL,
    `description` VARCHAR(256) NOT NULL
);

CREATE TABLE `votes` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `userid` INTEGER NOT NULL,
    `tarekomiid` INTEGER NOT NULL,
    `description` VARCHAR(256) NOT NULL
);

CREATE TABLE `stars` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `userid` INTEGER NOT NULL,
    `tarekomiid` INTEGER NOT NULL
);

INSERT INTO `blocks`(prevhash, create_time, data, hash) VALUES('initial', 0, '{}', 'ac1b5c0961a7269b6a053ee64276ed0e20a7f48aefb9f67519539d23aaf10149');

INSERT INTO users(screenname, password) VALUES('shinka', 'morisama');
