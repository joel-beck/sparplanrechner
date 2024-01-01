/**
 * Shows the tooltip by removing the 'hidden' class from the tooltip text.
 *
 * @param tooltip - The tooltip element.
 */
export function showTooltip(tooltip: HTMLElement): void {
    const tooltipText = tooltip.querySelector(".tooltip-text");

    if (!(tooltipText instanceof HTMLElement)) {
        throw new Error(".tooltip-text not found");
    }

    tooltipText.classList.remove("hidden");
}

/**
 * Hides the tooltip by adding the 'hidden' class to the tooltip text.
 *
 * @param tooltip - The tooltip element.
 */
export function hideTooltip(tooltip: HTMLElement): void {
    const tooltipText = tooltip.querySelector(".tooltip-text");

    if (!(tooltipText instanceof HTMLElement)) {
        throw new Error(".tooltip-text not found");
    }

    tooltipText.classList.add("hidden");
}
