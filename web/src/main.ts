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

		Alpine.store("formula1", {
			items: {
				Ferrari: "Charles Leclerc",
				Mercedes: "Lewis Hamilton",
				RedBull: "Max Verstappen",
			},
		});
	});

	Alpine.plugin(mask);
	Alpine.start();
}

function main () {
    initAlpine();
}

main();