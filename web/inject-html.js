/**
 * Fetches an HTML content from a given path
 *
 * @param {string} filePath - Path to the HTML file
 * @returns {string} HTML content as string
 * @throws {Error} Throws an error if fetch fails
 */
async function fetchHTML(filePath) {
  const response = await fetch(filePath);
  if (!response.ok) {
    throw new Error(`Failed to fetch ${filePath}: ${response.statusText}`);
  }
  return await response.text();
}

/**
 * Safely sets HTML content to a given element
 *
 * @param {HTMLElement} elem - The HTML element to set the content for
 * @param {string} html - The HTML content
 */
function setElementHTML(elem, html) {
  // For now, setting innerHTML directly.
  // In a production setting, consider sanitizing the HTML content for security.
  elem.innerHTML = html;
}

/**
 * Processes <script> tags in a given HTML element
 *
 * @param {HTMLElement} elem - The HTML element containing <script> tags
 */
function processScripts(elem) {
  elem.querySelectorAll("script").forEach((script) => {
    // Create a new empty <script> tag
    const newScript = document.createElement("script");
    // Copy attributes of existing script tag to a new one
    Array.from(script.attributes).forEach((attr) =>
      newScript.setAttribute(attr.name, attr.value)
    );
    // Inject content of existing script tag to a new one
    newScript.appendChild(document.createTextNode(script.innerHTML));
    // Replace existing script tag with a new one
    script.parentNode.replaceChild(newScript, script);
  });
}

/**
 * Injects HTML content into a given HTML element
 *
 * @param {string} filePath - Path to the HTML file
 * @param {HTMLElement} elem - The HTML element to inject the content into
 */
async function injectHTML(filePath, elem) {
  try {
    const html = await fetchHTML(filePath);
    setElementHTML(elem, html);
    processScripts(elem);
  } catch (err) {
    console.error(err.message);
  }
}

/**
 * Processes all HTML elements with "include" attribute
 *
 * This function injects the content of the file specified in the "include" attribute
 * into the corresponding HTML element.
 */
function injectAll() {
  document.querySelectorAll("div[include]").forEach((elem) => {
    const filePath = elem.getAttribute("include");
    injectHTML(filePath, elem);
  });
}

// Execute the injection
injectAll();
