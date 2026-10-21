const btn = document.getElementById("button");

btn.addEventListener("click", function() {
    window.location.href = `ready.html?${params.toString()}`;
});