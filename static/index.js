const btn = document.getElementById('btn');
const input = document.getElementById('inpt');
let sedangProsses = false;

btn.addEventListener("click", async function(e)  {
    if (sedangProsses) {
    	return;
    }
    if (input.value.trim()) {
        const originalUrl = document.getElementById('inpt').value.trim();
        const data = new URLSearchParams();
        data.append("url", originalUrl);
        sedangProsses = true;
        btn.disabled = true;
        
        try {
        const response = await fetch("/api/shorten", {
            method: "POST",
            body: data
        })  
        
            if (!response.ok) {
            	const pesan = await response.text();
                window.alert(pesan)
                return;
            } 
            const code = await response.text();
            
            const riwayat = ambilDataRiwayat();
            riwayat.push({ code: code, url: originalUrl});
            simpanRiwayat(riwayat);
            
            const params = new URLSearchParams();
            params.set("code", code);
            params.set("url", originalUrl);
            window.location.href = `ready.html?${params.toString()}`;
        } catch(err) {
        	window.alert("Gagal terhubung ke server")
        }
        
        finally {
        	sedangProsses = false;
        	btn.disabled = false;
        }
        
    } else {
        window.alert("Paste Your URL");
    }
});

	function ambilDataRiwayat() {
		const dataTersimpan = localStorage.getItem('riwayat');
		if (dataTersimpan === null) {
			return[];
		} else {
			return JSON.parse(dataTersimpan);
		}
	}
	
	function simpanRiwayat(data) {
		localStorage.setItem('riwayat', JSON.stringify(data));
	}