CREATE TABLE `users` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `screenname` VARCHAR(32) NOT NULL,
    `password` BLOB NOT NULL
);

CREATE TABLE `blocks` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `prevhash` VARCHAR(256) NOT NULL,
    `data` VARCHAR(1024), -- TODO: tarekomiとうまくリンクさせる
    `hash` VARCHAR(256) NOT NULL
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
)
