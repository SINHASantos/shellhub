/* eslint-disable */
// @ts-nocheck

/*

All environment variables should be read by this wrapper, which handles
the difference between development and production environments.
In production mode, the environment variables are stored in the global
`window.env` object, which is automatically generated by the container's
entrypoint in the `/env.js` file.
This is necessary because direct access to system environment variables
is not possible. On the other hand, in development mode, environment
variables can be accessed directly through `process.env` object.

*/


export const envVariables = {
  isEnterprise: (window.env || process.env).VUE_APP_SHELLHUB_ENTERPRISE === "true",
  isCloud: (window.env || process.env).VUE_APP_SHELLHUB_CLOUD === "true",
  premiumPaywall: (window.env || process.env).VUE_APP_SHELLHUB_PAYWALL === "true",
  stripePublishableKey: (window.env || process.env).VUE_APP_SHELLHUB_STRIPE_PUBLISHABLE_KEY,
  billingEnable: (window.env || process.env).VUE_APP_SHELLHUB_BILLING === "true",
  version: (window.env || process.env).VUE_APP_SHELLHUB_VERSION,
  announcementsEnable: (window.env || process.env).VUE_APP_SHELLHUB_ANNOUNCEMENTS === "true",
  stripeKey: (window.env || process.env).VUE_APP_SHELLHUB_STRIPE_PUBLISHABLE_KEY,
  sentryDsn: (window.env || process.env).VUE_APP_SHELLHUB_SENTRY_DSN_VERSION,
  googleAnalyticsID: (window.env || process.env).VUE_APP_SHELLHUB_GOOGLE_ANALYTICS_ID,
};
