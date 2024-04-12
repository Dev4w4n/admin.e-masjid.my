# Database setup

1. Create postgres image
```
docker build -t saas-postgres-image .
```

2. Run the container
```
docker run -d -p 5432:5432 --name saas-postgres-container saas-postgres-image
```
