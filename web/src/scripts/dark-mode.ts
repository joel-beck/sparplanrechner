function prefersSystemDarkmode(): boolean {
    return window.matchMedia("(prefers-color-scheme: dark)").matches;
}

function showLightIcon(themeToggle: HTMLElement, lightIcon: string): void {
    themeToggle.classList.add("bg-yellow-500", "-translate-x-3");
    themeToggle.classList.remove("bg-gray-700", "translate-x-3");

    setTimeout(() => {
        themeToggle.innerHTML = lightIcon;
    }, 100);
}

function setLightmode(): void {
    localStorage.setItem("isDarkmode", "false");
    document.documentElement.classList.remove("dark");
}

function showDarkIcon(themeToggle: HTMLElement, darkIcon: string): void {
    themeToggle.classList.add("bg-gray-700", "translate-x-3");
    themeToggle.classList.remove("bg-yellow-500", "-translate-x-3");

    setTimeout(() => {
        themeToggle.innerHTML = darkIcon;
    }, 100);
}

function setDarkmode(): void {
    localStorage.setItem("isDarkmode", "true");
    document.documentElement.classList.add("dark");
}

export function setInitialTheme(
    themeToggleIcon: HTMLElement,
    darkIcon: string,
    lightIcon: string,
): void {
    if (prefersSystemDarkmode()) {
        setDarkmode();
        showDarkIcon(themeToggleIcon, darkIcon);
        return;
    }

    setLightmode();
    showLightIcon(themeToggleIcon, lightIcon);
}

function isDarkmode(): boolean {
    return localStorage.getItem("isDarkmode") === "true";
}

export function toggleTheme(
    themeToggleIcon: HTMLElement,
    darkIcon: string,
    lightIcon: string,
): void {
    if (isDarkmode()) {
        setLightmode();
        showLightIcon(themeToggleIcon, lightIcon);
        return;
    }

    setDarkmode();
    showDarkIcon(themeToggleIcon, darkIcon);
}
