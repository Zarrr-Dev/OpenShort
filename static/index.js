const btn = document.getElementById('btn');
const input = document.getElementById('inpt');

btn.addEventListener("click", async function(e)  {
    if (input.value.trim()) {
        const originalUrl = document.getElementById('inpt').value.trim();
        const data = new URLSearchParams();
        data.append("url", originalUrl);

        const response = await fetch("/api/shorten", {
            method: "POST",
            body: data
        });

            if (!response.ok) {
                console.log("permintaan gagal");
                return;
            }
            const code = await response.text();
            console.log(code);

            const params = new URLSearchParams();
            params.set("code", code);
            params.set("url", originalUrl);
            window.location.href = `ready.html?${params.toString()}`;
    } else {
        e.preventDefault();
        window.alert("Paste Your URL");
    }
});