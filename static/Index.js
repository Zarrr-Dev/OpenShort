const button = document.getElementById("copyBtn");
const text = document.getElementById("outputUrl")

button.addEventListener("click", () => writeClipboardText(outputUrl.value));

async function writeClipboardText(text) {
    try {
        await navigator.clipboard.writeText(text);
    } catch (error) {
        console.error(error.message);
    }
    
}