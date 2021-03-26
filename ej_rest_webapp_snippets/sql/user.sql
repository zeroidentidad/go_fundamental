CREATE USER 'webusr'@'%';

GRANT SELECT, INSERT ON pastein.* TO 'webusr'@'%';

-- Important: Swap 'pass' with a diff password.
ALTER USER 'webusr'@'%' IDENTIFIED BY 'passx123';