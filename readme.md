# Sample compose format

```yaml
services:
  frontend:
    build: 
      context: frontend/
      args:
        - REACT_APP_BACKEND_URL=http://localhost:3001/ # Backend URL
    ports:
      - 3000:80
  backend:
    build: backend/
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock # Needs to connect to the docker socket outside the container
    ports:
      - 3001:3000
```