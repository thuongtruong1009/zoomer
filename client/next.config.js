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
    // images: {
    //   domains: ['images.unsplash.com'],
    // },

    webpack(config, { isServer }) {
      // Enable webpack caching for faster subsequent builds
      if (!isServer) {
        config.cache = {
          type: 'filesystem',
        };
      }

      return config;
    },
}
