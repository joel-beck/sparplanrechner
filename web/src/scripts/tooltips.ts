/**
 * Shows the tooltip by removing the 'hidden' class from the tooltip text.
 *
 * @param tooltip - The tooltip element.
 */
export function showTooltip(tooltip: HTMLElement): void {
    const tooltipText = tooltip.querySelector(".tooltip-text") as HTMLElement;
    tooltipText.classList.remove("hidden");
}

/**
 * Hides the tooltip by adding the 'hidden' class to the tooltip text.
 *
 * @param tooltip - The tooltip element.
 */
export function hideTooltip(tooltip: HTMLElement): void {
    const tooltipText = tooltip.querySelector(".tooltip-text") as HTMLElement;
    tooltipText.classList.add("hidden");
}
