module.exports = {
    reactStrictMode: true,
    async rewrites() {
      return [
        {
          source: "/api/auth/:provider",
          destination: "/api/auth/[provider]",
        },
      ];
    },
}
