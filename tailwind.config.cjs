/** @type {import('tailwindcss').Config} */
module.exports = {
    darkMode: "class",
    content: ["./web/**/*.{html,js}"],
    theme: {
        extend: {
            colors: {
                "color-text-light": "#313131",
                "color-text-dark": "#D3D3D3",
                "color-bg-light": "#F8FAFC",
                // "color-bg-dark": "#193549",
                "color-bg-dark": "#112534",
                // "color-accent-light": "#F56565",
                "color-accent-light": "#E05C5D",
                // "color-accent-dark": "#F49608",
                "color-accent-dark": "#DB8910",
                "color-hover-light": "#CD5456",
                "color-hover-dark": "#C07941",
                "color-form-light": "#EFEFEF",
                "color-form-dark": "#E5E4E2",
            },
            fontFamily: {
                body: ["Nunito"],
            },
        },
    },
    plugins: [
        require("@tailwindcss/typography"),
        require("@tailwindcss/forms"),
        require("@tailwindcss/aspect-ratio"),
        require("@tailwindcss/container-queries"),
    ],
};
