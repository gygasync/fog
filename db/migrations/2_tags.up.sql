CREATE TABLE Tag ( 
	Id                   varbinary(16) NOT NULL  PRIMARY KEY  ,
	Name                 varchar(255) NOT NULL    
 );

CREATE TABLE TagToTag ( 
	Id                   integer NOT NULL PRIMARY KEY AUTOINCREMENT ,
	Source               varbinary(16) NOT NULL    ,
	Target               varbinary(16) NOT NULL    ,
	FOREIGN KEY ( Source ) REFERENCES Tag( Id )  ,
	FOREIGN KEY ( Target ) REFERENCES Tag( Id )  
 );

CREATE TABLE Reference ( 
	Id                   integer NOT NULL PRIMARY KEY AUTOINCREMENT ,
	Tag                  varbinary(16) NOT NULL    ,
	Item                 varbinary(16) NOT NULL    ,
	FOREIGN KEY ( Tag ) REFERENCES Tag( Id )  
 );

CREATE INDEX Idx_Reference ON Reference ( Item );
