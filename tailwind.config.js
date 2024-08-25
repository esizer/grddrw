/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.{templ,go}", "./src/js/**/*.js"],
  theme: {
    container: {
      center: true,
    },
    fontFamily: {
      sans: ["Inter", "sans-serif"],
    },
    boxShadow: {
      DEFAULT: "0.25rem 0.25rem 0 0 rgba(0, 0, 0, 0.6)",
      none: "box-shadow: 0 0 rgba(0,0,0,0)",
      header: "0 0.25rem 0 0 rgba(0,0,0,0.6)",
      footer: "0 -0.25rem 0 0 rgba(0,0,0,0.6)",
    },
    extend: {
      gridTemplateColumns: {
        16: "repeat(16, minmax(0, 1fr))",
      },
    },
  },
  plugins: [],
};
