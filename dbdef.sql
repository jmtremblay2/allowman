create table households (
  hid          integer,

  PRIMARY KEY (hid)
);

create table budgets (
  bid integer,

  hid integer,

  name text,
  amounti real,

  PRIMARY KEY (bid),
  FOREIGN KEY (hid) REFERENCES households(hid)
);

create table accounts (
  accid integer,

  hid integer,

  name text,

  PRIMARY KEY (accid),
  FOREIGN KEY (hid) REFERENCES households(hid)
);

create table statements (
  stmtid integer,

  accid integer,

  stmtdate date,
  stmtstart date,
  stmtend date,

  PRIMARY KEY (stmtid),
  FOREIGN KEY (accid) REFERENCES accounts(accid)
);


create table tags (
  tid integer,

  descr text,

  PRIMARY KEY (tid)
);

create table transactions (
  trid integer,

  stmtid integer,

  descr text,
  ammount real,

  PRIMARY KEY (trid),
  FOREIGN KEY (stmtid) REFERENCES statements(stmtid)
);

create table trtags (
  trid integer,
  tid integer,

  primarytag boolean,

  PRIMARY KEY (trid, tid),
  FOREIGN KEY (trid) REFERENCES transactions(trid),
  FOREIGN KEY (tid) REFERENCES tags(tid)
);


