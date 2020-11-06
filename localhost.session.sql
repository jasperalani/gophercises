-- CREATE OR REPLACE USER test IDENTIFIED BY 'test';

create DATABASE test;

INSERT INTO `forwarding-table` (short,`long`) VALUES ("aaaaaaa","www.google.com");

select * from `forwarding-table`;

SELECT short, `long` FROM `forwarding-table` WHERE short = "UJ37cEb";