import colors from 'vuetify/es5/util/colors'

export default {
  srcDir: 'src',
  // Disable server-side rendering: https://go.nuxtjs.dev/ssr-mode
  ssr: false,

  // Target: https://go.nuxtjs.dev/config-target
  target: 'static',

  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    titleTemplate: '%s - shs-web',
    title: 'shs-web',
    htmlAttrs: {
      lang: 'en',
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: ['~/plugins/axios', '~/plugins/axios-accessor', '~/plugins/firebase'],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: ['@nuxt/typescript-build', '@nuxtjs/composition-api/module', '@nuxtjs/vuetify'],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: ['@nuxtjs/axios'],

  router: {
    middleware: ['authenticated'],
  },

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {},

  // Vuetify module configuration: https://go.nuxtjs.dev/config-vuetify
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    theme: {
      dark: false,
      themes: {
        light: {
          primary: colors.lightBlue.accent2,
          primaryLight: colors.lightBlue.lighten3,
          secondary: colors.amber.darken3,
          accent: colors.grey.darken3,
          info: colors.teal.lighten1,
          warning: colors.amber.base,
          error: colors.deepOrange.accent4,
          success: colors.green.accent4,
          white: colors.shades.white,
          grey: colors.grey.lighten4,
        },
        dark: {
          primary: colors.lightBlue.accent2,
          primaryLight: colors.lightBlue.lighten3,
          secondary: colors.amber.darken3,
          accent: colors.grey.darken3,
          info: colors.teal.lighten1,
          warning: colors.amber.base,
          error: colors.deepOrange.accent4,
          success: colors.green.accent4,
          white: colors.shades.white,
          grey: colors.grey.lighten4,
        },
      },
    },
  },

  typescript: {
    typeCheck: {
      eslint: {
        files: './src/**/*.{ts,js,vue}',
      },
    },
  },

  env: {
    apiURL: process.env.API_URL,
    firebaseApiKey: process.env.FIREBASE_API_KEY,
    firebaseProjectId: process.env.FIREBASE_PROJECT_ID,
    firebaseMessagingSenderId: process.env.FIREBASE_MESSAGING_SENDER_ID,
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {},
}
