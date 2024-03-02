import mask from "@alpinejs/mask";
import collapse from "@alpinejs/collapse";
import intersect from "@alpinejs/intersect";
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
    Alpine.plugin(collapse); // eslint-disable-line
    Alpine.plugin(intersect); // eslint-disable-line
    Alpine.start();
}

function main() {
    initAlpine();
}

main();
