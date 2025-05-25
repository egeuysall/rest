# Install dependencies
FROM node:22.16.0-alpine AS deps
WORKDIR /app

# Install pnpm globally
RUN npm install -g pnpm

COPY package.json pnpm-lock.yaml* ./
RUN pnpm install --frozen-lockfile

# Build app
FROM node:22.16.0-alpine AS builder
WORKDIR /app

RUN npm install -g pnpm
COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN pnpm build

# Production image
FROM node:22.16.0-alpine AS runner
WORKDIR /app

ENV NODE_ENV=production
ENV PORT=3000

COPY --from=builder /app/public ./public
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json

EXPOSE 3000

CMD ["pnpm", "start"]