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
INSERT INTO users(screenname, password) VALUES('guest', 'guestuser');

INSERT INTO tarekomi(status, threshold, targetuserid, url, description) VALUES(1, 10, 1, "https://web.archive.org/save/http://d.hatena.ne.jp/aplle-5/20121014", "hatobus's dark history");
INSERT INTO tarekomi(status, threshold, targetuserid, url, description) VALUES(1, 10, 1, "https://web.archive.org/save/https://hatobus.hatenablog.jp/entry/2018/07/05/104114", "hatobus's dark history blog");
INSERT INTO tarekomi(status, threshold, targetuserid, url, description) VALUES(0, 10, 1, "https://web.archive.org/web/20190430051911/http://noahorberg.hatenablog.com/entry/2018/12/16/204957?_ga=2.231062615.953365373.1556587926-311450958.1547619044", "noah's dark history blog");

-- FOR VOTE TEST
INSERT INTO tarekomi(status, threshold, targetuserid, url, description) VALUES(0, 1, 1, "https://web.archive.org/save/http://noahorberg.hatenablog.com/entry/2015/03/02/095231", "noah's dark history blog");
