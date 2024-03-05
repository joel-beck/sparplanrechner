import mask from "@alpinejs/mask";
import Alpine from "alpinejs";

function initAlpine() {
    document.addEventListener("alpine:init", () => {
        Alpine.store("darkMode", {
            on: window.matchMedia("(prefers-color-scheme: dark)").matches,
            toggle() {
                this.on = !this.on;
            },
        });
    });

    Alpine.plugin(mask); // eslint-disable-line
    Alpine.start();
}

function main() {
    initAlpine();
}

main();
