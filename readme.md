# Sample compose format

```yaml
services:
  frontend:
    build: 
      context: frontend/
      args:
        - VITE_BACKEND_URL=http://localhost:3001/ # Backend URL
    ports:
      - 3000:80
  backend:
    build: backend/
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock # Needs to connect to the docker socket outside the container
    ports:
      - 3001:3000
  caddy:
    image: caddy
    container_name: searchengine-caddy
    restart: always
    ports:
      - "${PORT}:80"
    volumes:
      - ./config-prod/Caddyfile:/etc/caddy/Caddyfile
```