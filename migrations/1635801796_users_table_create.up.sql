Create Table Users (
    UserId Integer Primary Key,
    Alias Varchar(50) Not Null Unique,
    FullName Varchar(100) Not Null,
    Active Boolean Default False,
    Locked Boolean Default False
);
