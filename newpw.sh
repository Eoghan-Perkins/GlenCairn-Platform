#!/bin/bash

# Replace with your desired new root password
NEW_PASSWORD="hotjazz911"

echo "Stopping MySQL service..."
sudo service mysql stop

echo "Starting MySQL in safe mode..."
sudo mysqld_safe --skip-grant-tables &

# Wait a bit to make sure MySQL has started in safe mode
sleep 5

echo "Resetting MySQL root password..."
mysql -u root <<EOF
ALTER USER 'root'@'localhost' IDENTIFIED BY '$NEW_PASSWORD';
FLUSH PRIVILEGES;
EOF

# Stop the safe mode MySQL instance
echo "Stopping MySQL safe mode..."
sudo killall mysqld_safe
sudo killall mysqld

# Wait a bit to make sure MySQL has stopped completely
sleep 5

echo "Starting MySQL service normally..."
sudo service mysql start

echo "MySQL root password has been reset to '$NEW_PASSWORD'"
echo "Remember to update any applications or scripts that use the MySQL root password."

# Cleanup
unset NEW_PASSWORD
