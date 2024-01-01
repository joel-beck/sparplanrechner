// Credits for this file go to https://dev.to/andreygermanov/modular-html-19o6

/**
 * Fetches an HTML content from a given path
 *
 * @param filePath - Path to the HTML file
 * @returns HTML content as string
 * @throws Throws an error if fetch fails
 */
async function fetchHTML(filePath: string): Promise<string> {
    const response = await fetch(filePath);
    if (!response.ok) {
        throw new Error(`Failed to fetch ${filePath}: ${response.statusText}`);
    }
    return await response.text();
}

/**
 * Safely sets HTML content to a given element
 *
 * @param elem - The HTML element to set the content for
 * @param html - The HTML content
 */
function setElementHTML(elem: HTMLElement, html: string): void {
    // For now, setting innerHTML directly.
    // In a production setting, consider sanitizing the HTML content for security.
    elem.innerHTML = html;
}

/**
 * Processes <script> tags in a given HTML element
 *
 * @param elem - The HTML element containing <script> tags
 */
function processScripts(elem: HTMLElement): void {
    elem.querySelectorAll("script").forEach((script) => {
        const newScript = document.createElement("script");
        Array.from(script.attributes).forEach((attr) => {
            newScript.setAttribute(attr.name, attr.value);
        });
        newScript.appendChild(document.createTextNode(script.innerHTML));
        script.parentNode?.replaceChild(newScript, script);
    });
}

/**
 * Injects HTML content into a given HTML element
 *
 * @param filePath - Path to the HTML file
 * @param elem - The HTML element to inject the content into
 */
async function injectHTMLFile(
    filePath: string,
    elem: HTMLElement,
): Promise<void> {
    try {
        const html = await fetchHTML(filePath);
        setElementHTML(elem, html);

        processScripts(elem);
    } catch (err) {
        console.error((err as Error).message);
    }
}

/**
 * Processes all HTML elements with "include" attribute
 *
 * This function injects the content of the file specified in the "include" attribute
 * into the corresponding HTML element.
 */
export async function injectHTMLFiles(): Promise<void> {
    const elements = Array.from(
        document.querySelectorAll<HTMLElement>("div[include]"),
    );
    for (const elem of elements) {
        const filePath = elem.getAttribute("include");
        if (filePath) {
            await injectHTMLFile(filePath, elem);
        }
    }
}
