CREATE TABLE Directory ( 
	Id                   varchar(36) NOT NULL  PRIMARY KEY  ,
	Path                 text NOT NULL    ,
	Dateadded            datetime  DEFAULT CURRENT_TIMESTAMP   ,
	Lastchecked          datetime     
 );

CREATE INDEX Idx_Directory ON Directory ( Id );

CREATE TABLE File ( 
	Id                   varchar(36) NOT NULL  PRIMARY KEY  ,
	Path                 text     ,
	Parentdirectory      varchar(36) NOT NULL    ,
	Checksum             binary(32)     ,
	Lastchecked          datetime     ,
	FOREIGN KEY ( Parentdirectory ) REFERENCES Directory( Id )  
 );

CREATE INDEX Idx_File ON File ( Id );

