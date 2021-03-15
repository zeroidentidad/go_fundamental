CREATE USER 'webusr'@'%';

GRANT SELECT, INSERT ON snippetbox.* TO 'webusr'@'%';

-- Important: Swap 'pass' with a diff password.
ALTER USER 'webusr'@'%' IDENTIFIED BY 'passx123';