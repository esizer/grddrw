/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.templ"],
  theme: {
    fontFamily: {
      sans: ["Inter", "sans-serif"],
      shadow: {
        DEFAULT: "2rem 2rem 0 0 rgba(0, 0, 0, 0.7)",
      },
    },
    extend: {
      gridTemplateColumns: {
        16: "repeat(16, minmax(0, 1fr))",
      },
    },
  },
  plugins: [],
};
