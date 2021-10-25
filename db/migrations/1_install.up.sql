CREATE TABLE Directory ( 
	Id                   varbinary(16) NOT NULL  PRIMARY KEY  ,
	Path                 text NOT NULL    ,
	Dateadded            datetime  DEFAULT CURRENT_DATE   ,
	Lastchecked          datetime     ,
	ParentDirectory      varbinary(16)     ,
	FOREIGN KEY ( ParentDirectory ) REFERENCES Directory( Id )  
 );

CREATE INDEX Idx_Directory ON Directory ( Id );

CREATE TABLE File ( 
	Id                   varbinary(16) NOT NULL  PRIMARY KEY  ,
	Path                 bit NOT NULL    ,
	ParentDirectory      varbinary(16)     ,
	Checksum             text     ,
	Lastchecked          datetime     ,
	MimeType             varchar(100)     ,
	FOREIGN KEY ( ParentDirectory ) REFERENCES Directory( Id )  
 );

CREATE INDEX Idx_File ON File ( Id );
