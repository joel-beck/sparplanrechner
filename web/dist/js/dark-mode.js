"use strict";

function prefersSystemDarkmode() {
    return window.matchMedia("(prefers-color-scheme: dark)").matches;
}

function showLightIcon(themeToggle, lightIcon) {
    themeToggle.classList.add("bg-yellow-500", "-translate-x-3");
    themeToggle.classList.remove("bg-gray-700", "translate-x-3");

    setTimeout(() => {
        themeToggle.innerHTML = lightIcon;
    }, 100);
}

function setLightmode() {
    localStorage.setItem("isDarkmode", false);
    document.documentElement.classList.remove("dark");
}

function showDarkIcon(themeToggle, darkIcon) {
    themeToggle.classList.add("bg-gray-700", "translate-x-3");
    themeToggle.classList.remove("bg-yellow-500", "-translate-x-3");

    setTimeout(() => {
        themeToggle.innerHTML = darkIcon;
    }, 100);
}

function setDarkmode() {
    localStorage.setItem("isDarkmode", true);
    document.documentElement.classList.add("dark");
}

export function setInitialTheme(themeToggleIcon, darkIcon, lightIcon) {
    if (prefersSystemDarkmode()) {
        setDarkmode();
        showDarkIcon(themeToggleIcon, darkIcon);
        return;
    }

    setLightmode();
    showLightIcon(themeToggleIcon, lightIcon);
}

function isDarkmode() {
    return localStorage.getItem("isDarkmode") === "true";
}

export function toggleTheme(themeToggleIcon, darkIcon, lightIcon) {
    if (isDarkmode()) {
        setLightmode();
        showLightIcon(themeToggleIcon, lightIcon);
        return;
    }

    setDarkmode();
    showDarkIcon(themeToggleIcon, darkIcon);
}
