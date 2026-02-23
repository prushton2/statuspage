# Sample compose format

```yaml
services:
  frontend:
    container_name: statuspage-frontend
    build: frontend/

  backend:
    container_name: statuspage-backend
    build: backend/
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock # Needs to connect to the docker socket outside the container
    environment:
      IGNORE_CONTAINERS: |
        statuspage-caddy
  caddy:
    image: caddy
    container_name: statuspage-caddy
    restart: always
    ports:
      - "8080:80"
    volumes:
      - ./Caddyfile:/etc/caddy/Caddyfile
```