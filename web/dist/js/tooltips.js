"use strict";

export function showTooltip(tooltip) {
    tooltip.querySelector(".tooltip-text").classList.remove("hidden");
}

export function hideTooltip(tooltip) {
    tooltip.querySelector(".tooltip-text").classList.add("hidden");
}
