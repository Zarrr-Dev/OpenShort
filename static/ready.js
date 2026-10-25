const copybtn = document.getElementById('copybtn');
const text = document.getElementById('outputUrl');
const urlParams = new URLSearchParams(window.location.search);
const baseURL = window.location.origin;
const path = '/s/';
const code = urlParams.get('code');


copybtn.addEventListener("click", () => writeClipboardText(text.value));

async function writeClipboardText(text) {
    try {
        await navigator.clipboard.writeText(text);
    } catch (error) {
        console.error(error.message);
        window.alert("Failed to copy URL. Please copy it manually")
    }
}
if (code) {
	text.value = baseURL + path + code;
} else {
	text.value = "";
}

const url = document.getElementById('originUrl');
url.innerText = urlParams.get('url');
