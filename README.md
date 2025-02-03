## react-fsd-template

Production deploy: `docker compose -f docker-compose.yml up --build`

#### Check: 
- Backend: http://localhost:8080/swagger/index.html
- Frontend: http://localhost:3000/
### Ensure you have ignored .env files in your own project
### Checkout .gitignore and /frontend/.gitignore files

## logs:
`docker compose -f docker-compose.yml logs --tail=50` - log
`docker compose -f docker-compose.yml build --no-cache frontend` - reload no cache

RUN echo "Building frontend with API URL: $VITE_API_URL"

### Frontend todo:
- make base Redux config
- fulfill @shared/


## DOCKERFILE
```Dockerfile
# Stage 1: Build Stage
FROM node:20-alpine AS builder
WORKDIR /app

RUN corepack enable && corepack prepare pnpm@latest --activate
COPY package.json pnpm-lock.yaml ./
# RUN corepack enable && pnpm install --frozen-lockfile
RUN pnpm install --frozen-lockfile

# Copy the rest source code and build the project
COPY . .

# Pass build-time variable for API URL and build the project
ARG VITE_API_URL
ENV VITE_API_URL=$VITE_API_URL
RUN echo "Building frontend with API URL: $VITE_API_URL"

RUN pnpm build

# Copy static assets to the build output
# RUN cp -r public/* dist/

# Stage 2: Production Stage using Caddy
FROM caddy:latest
WORKDIR /usr/share/caddy

# Copy built files from builder stage
COPY --from=builder /app/dist .
RUN mkdir -p /usr/share/caddy/dist
# Copy the Caddy configuration file from the new location
COPY config/caddy/Caddyfile /etc/caddy/Caddyfile

EXPOSE 3000
CMD ["caddy", "run", "--config", "/etc/caddy/Caddyfile"]
```

## DOCKERFILE 2
```Dockerfile
# Stage 1: Build Stage
FROM node:alpine AS builder
WORKDIR /app

COPY package.json pnpm-lock.yaml ./
RUN corepack enable && corepack prepare pnpm@latest --activate
# COPY package.json pnpm-lock.yaml ./
# RUN corepack enable && pnpm install --frozen-lockfile
RUN pnpm install --frozen-lockfile

# Copy the rest source code and build the project
COPY . .

# Pass build-time variable for API URL and build the project
ARG VITE_API_URL
ENV VITE_API_URL=$VITE_API_URL
RUN echo "Building frontend with API URL: $VITE_API_URL"

# RUN pnpm build
RUN pnpm build && ls -lah dist

# Copy static assets to the build output
# RUN cp -r public/* dist/

# Stage 2: Production Stage using Caddy
FROM caddy:latest
WORKDIR /usr/share/caddy

# Copy built files from builder stage
COPY --from=builder /app/dist /usr/share/caddy/dist
RUN mkdir -p /usr/share/caddy/dist
# Copy the Caddy configuration file from the new location
COPY config/caddy/Caddyfile /etc/caddy/Caddyfile

EXPOSE 3000
CMD ["caddy", "run", "--config", "/etc/caddy/Caddyfile"]
```