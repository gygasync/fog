CREATE TABLE Directory ( 
	Id                   varbinary(32) NOT NULL  PRIMARY KEY  ,
	Path                 text NOT NULL    ,
	Dateadded            datetime  DEFAULT CURRENT_TIMESTAMP   ,
	Lastchecked          datetime     
 );

CREATE INDEX Idx_Directory ON Directory ( Id );

CREATE TABLE File ( 
	Id                   varbinary(32) NOT NULL  PRIMARY KEY  ,
	Path                 text     ,
	Parentdirectory      varbinary(32) NOT NULL    ,
	Checksum             binary(32)     ,
	Lastchecked          datetime     ,
	FOREIGN KEY ( Parentdirectory ) REFERENCES Directory( Id )  
 );

CREATE INDEX Idx_File ON File ( Id );
