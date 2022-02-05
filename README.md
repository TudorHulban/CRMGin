# GinCRM

## Init testing
Testing should start with:<br/>
a. clean up RDBMS
```
setup.CleanRDBMS()
```

b. create tables
```
require.Nil(t, cgorm.MigrateDBSchema())
```

c. insert data
Pointers should be used to pass to ORM with data to insert.

## Persisting structures
Definition should reside in /pkg/persistence package. 

## Test User creation
```
curl -v POST -F "usercode=john" -F "password=1234" http://localhost:8080/auth/createuser
curl -v POST -F "usercode=john" -F "password=1234" http://localhost:8080/auth/login
```

