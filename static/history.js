const daftar = document.getElementById('daftar');
const dataTersimpan = localStorage.getItem('riwayat');
if ( dataTersimpan=== null ) {
	daftar.textContent = "No history yet";
} else {
	const riwayat = JSON.parse(dataTersimpan);
	
	riwayat.reverse();
	const terbaru = riwayat.slice(0, 10);
	
	terbaru.forEach((item) => {
	const linkCode = document.createElement('a');
	const teksUrl = document.createElement('p');
	
	linkCode.textContent = item.code;
	linkCode.href = '/s/' + item.code;
	teksUrl.textContent = item.url;
	const newDiv = document.createElement('div')
	newDiv.classList.add('riwayat-item')
	linkCode.classList.add('kotak-link');
	
	newDiv.append(linkCode, teksUrl);
	daftar.append(newDiv);
	
})};