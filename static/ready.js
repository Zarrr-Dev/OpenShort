const copybtn = document.getElementById('copybtn');
const text = document.getElementById('outputUrl');
const urlParams = new URLSearchParams(window.location.search);
Url = urlParams.get('url');


copybtn.addEventListener("click", () => writeClipboardText(text.value));

async function writeClipboardText(text) {
    try {
        await navigator.clipboard.writeText(text);
    } catch (error) {
        console.error(error.message);
    }
}