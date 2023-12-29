/**
 * Toggles the visibility of a table container and updates the icon accordingly.
 * @param tableContainer - The table container element to collapse or expand.
 * @param icon - The icon element to update based on the visibility state.
 */
export function collapseTable(
    tableContainer: HTMLElement,
    icon: HTMLElement,
): void {
    tableContainer.classList.toggle("hidden");

    if (icon.classList.contains("fa-chevron-down")) {
        icon.classList.remove("fa-chevron-down");
        icon.classList.add("fa-chevron-up");
    } else {
        icon.classList.remove("fa-chevron-up");
        icon.classList.add("fa-chevron-down");
    }
}
