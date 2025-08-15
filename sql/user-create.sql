CREATE USER 'web'@'%' IDENTIFIED BY 'pass';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'%';
-- Important: Make sure to swap 'pass' with a password of your own choosing.
FLUSH PRIVILEGES;
