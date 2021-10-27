CREATE TABLE MetadataType ( 
	Id                   integer NOT NULL  PRIMARY KEY AUTOINCREMENT  ,
	Name                 varchar(255) NOT NULL    ,
	CONSTRAINT Unq_MetadataType UNIQUE ( Name ) 
 );

CREATE TABLE Metadata ( 
	Id                   varbinary(16) NOT NULL  PRIMARY KEY  ,
	MetaType             integer NOT NULL    ,
	Reference            varbinary(16) NOT NULL    ,
    Value                text     ,
	FOREIGN KEY ( MetaType ) REFERENCES MetadataType( Id )  
 );

CREATE INDEX Idx_Metadata ON Metadata ( Id );
