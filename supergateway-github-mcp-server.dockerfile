from docker_mirror/github-github-mcp-server:latest-linux-amd64 as builder




from docker.cnb.cool/masx200/docker_mirror/mcp-github-server:2025-08-03-15-23-58 as base


copy --from=builder /server/github-mcp-server /server/github-mcp-server


entrypoint ["docker-entrypoint.sh"]


CMD [ "node" ,"/root/mcp-demo-streamable-http-bridge/bridge-streamable.js" ,"/server/github-mcp-server","stdio"]


env PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/server:/root/mcp-demo-streamable-http-bridge



env BRIDGE_API_PORT=3000