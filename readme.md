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
    # IGNORE_CONTAINERS are ignored by the backend
    # TOP_NETWORKS are names of networks that are always on top of the list
    # BOTTOM_NETWORKS are names of networks that are always on the bottom of the list
    #   Both of the above fields will preserve the order that they are listed in the env
    environment:
      IGNORE_CONTAINERS: |
        statuspage-caddy
      TOP_NETWORKS: |
        minecraft_default
        plex_default
      BOTTOM_NETWORKS: |
        backup_default
        statuspage_default
        server_default
        host
  caddy:
    image: caddy
    container_name: statuspage-caddy
    restart: always
    ports:
      - "8080:80"
    volumes:
      - ./Caddyfile:/etc/caddy/Caddyfile
```