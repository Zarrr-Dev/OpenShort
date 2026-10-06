const btn = document.getElementById('btn');
const input = document.getElementById('inpt');
const originalUrl = document.getElementById('inpt').value.trim();

btn.addEventListener("click", function(e)  {
    if (input.value.trim()) {
        const encodeUrl = encodeURIComponent("originalUrl")
        window.location.href = `ready.html?url=${encodeUrl}`;
    } else {
        e.preventDefault();
        window.alert("Paste Your URL");
    }
});
