const copybtn = document.getElementById('copybtn');
const text = document.getElementById('outputUrl');
const urlParams = new URLSearchParams(window.location.search);
const Url = urlParams.get('url');
const baseURL = window.location.origin;
const path = '/s/';
const code = urlParams.get('code');
const hasil = baseURL + path + code;
text.value = hasil;

copybtn.addEventListener("click", () => writeClipboardText(text.value));

async function writeClipboardText(text) {
    try {
        await navigator.clipboard.writeText(text);
    } catch (error) {
        console.error(error.message);
    }
}


const url = document.getElementById('originUrl');
url.innerText = urlParams.get('url');
