Create Table Users (
    UserId Integer Primary Key,
    Alias Nvarchar(50) Not Null Unique,
    FullName Nvarchar(100) Not Null,
    Active Boolean Default 0,
    Locked Boolean Default 0
);