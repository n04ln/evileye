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
INSERT INTO users(screenname, password) VALUES('yt8492', 'mayamito');
INSERT INTO users(screenname, password) VALUES('flying_hato_bus', 'hatobus');
INSERT INTO users(screenname, password) VALUES('NoahOrberg', '12345678');

INSERT INTO tarekomi(status, threshold, targetuserid, url, description) VALUES(1, 10, 1, "https://web.archive.org/save/http://d.hatena.ne.jp/aplle-5/20121014", "hatobus's dark history");
INSERT INTO tarekomi(status, threshold, targetuserid, url, description) VALUES(1, 10, 1, "https://web.archive.org/save/https://twitter.com/yt8492/status/1114130271316504576", "yt8492's kozirase tweet");
INSERT INTO tarekomi(status, threshold, targetuserid, url, description) VALUES(0, 10, 1, "https://web.archive.org/save/https:/twitter.com/NoahOrberg/status/835029021612089344", "noah's kozirase tweet");

-- FOR VOTE TEST
INSERT INTO tarekomi(status, threshold, targetuserid, url, description) VALUES(0, 1, 1, "https://web.archive.org/save/https://twitter.com/yt8492/status/1119128753148325888", "yt8492's kozirase tweet");
