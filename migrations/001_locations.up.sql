CREATE TABLE locations (
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT,
    mtime TEXT CHECK (mtime IS date(mtime))
) STRICT;
